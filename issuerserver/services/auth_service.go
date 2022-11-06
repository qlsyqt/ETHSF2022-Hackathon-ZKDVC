package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	auth "github.com/iden3/go-iden3-auth"
	"github.com/iden3/go-rapidsnark/verifier"
	"github.com/iden3/iden3comm/protocol"
	"io/ioutil"
	"issuerserver/config"
	common2 "issuerserver/models/common"
	"issuerserver/pkg/dvc"
	"issuerserver/pkg/logging"
	"issuerserver/store"
	"issuerserver/utils"
	"strconv"
	"time"
)

func Auth(offerId int64) (*protocol.AuthorizationRequestMessage, error) {
	cfg := config.GetConfig()

	offer, err := store.SelectOfferById(offerId)
	if err != nil {
		return nil, err
	}

	issuer := offer.IssuerDid
	challenge := []byte(strconv.FormatInt(offerId, 10)) //todo: should add more info here

	authRule := make(map[string]interface{})
	authRule["challenge"] = hexutil.Encode(challenge)
	authScope := protocol.ZeroKnowledgeProofRequest{
		ID:        0,
		CircuitID: "authKnn3",
		Rules:     authRule,
	}

	request := auth.CreateAuthorizationRequestWithMessage("authKnn3", "", issuer, cfg.Authentication.Callback)
	request.ID = strconv.FormatInt(offerId, 10)
	request.Body.Scope = append(request.Body.Scope, authScope)

	return &request, nil
}

func CallbackAuth(callbackResponse *protocol.AuthorizationResponseMessage) (string, string, error) {
	//Zkp verify
	proof := callbackResponse.Body.Scope[0].ZKProof
	signals := callbackResponse.Body.Scope[0].PubSignals
	vKey, _ := ioutil.ReadFile("zk/auth.vkey") //TODO：写入config
	err := verifier.VerifyGroth16(proof, vKey)
	if err != nil {
		fmt.Println(err)
	}
	logging.Info.Println("zk proof verify success!")
	userId := signals[0]
	if userId != callbackResponse.From {
		errors.New("Invalid holder did")
	}
	//Signature verify
	walletSignature := callbackResponse.Body.Scope[1].PubSignals[0]
	walletAddress := common.HexToAddress(callbackResponse.Body.Scope[1].PubSignals[1])
	digest, err := utils.HashChallengeAndWallet([]byte(callbackResponse.ID), walletAddress)
	if err != nil {
		return "", "", err
	}
	sigBytes, err := hexutil.Decode(walletSignature)
	if err != nil {
		return "", "", err
	}
	pubKey, err := crypto.SigToPub(digest.Bytes(), sigBytes)
	if err != nil {
		return "", "", err
	}
	recoveredAddress := utils.PublicKeyBytesToAddress(crypto.FromECDSAPub(pubKey))
	if walletAddress != recoveredAddress {
		return "", "", errors.New("invalid signature")
	}
	logging.Info.Println("wallet signature verify success!")
	//fetch dvc data
	offerId, err := strconv.ParseInt(callbackResponse.ID, 10, 64)
	if err != nil {
		return "", "", err
	}
	offer, err := store.SelectOfferById(offerId)
	if err != nil {
		return "", "", err
	}
	template, err := store.SelectTemplateById(offer.TemplateId)
	if err != nil {
		return "", "", err
	}
	logging.Info.Println("start fetching claim value")
	logging.Info.Println(template.DataCategory, " ", template.SubCategory, " ", walletAddress.Hex())
	claimValue, err := dvc.Query(context.Background(), template.DataCategory, template.SubCategory, walletAddress.Hex())
	if err != nil {
		return "", "", err
	}
	logging.Info.Println("claimValue ", claimValue)
	//fetch pre template
	preClaims := make([]common2.PreClaim, 0)
	err = json.Unmarshal([]byte(offer.PreClaims), &preClaims)
	if err != nil {
		return "", "", err
	}
	logging.Info.Println("preclaims:")
	logging.Info.Println(preClaims)
	//var chosenPreClaim common2.PreClaim
	var lower *common2.Boundary
	var upper *common2.Boundary
	found := false
	for _, p := range preClaims {
		lower, err = utils.RawToBoundary(p.LowerBound, true)
		if err != nil {
			return "", "", err
		}
		upper, err = utils.RawToBoundary(p.UpperBound, false)
		if err != nil {
			return "", "", err
		}
		if utils.InRange(claimValue, lower, upper) {
			found = true
			break
		}
	}
	if !found {
		fmt.Println("claim not match:", claimValue)
		return "", "", errors.New("preclaim Not Match")
	}
	claimRequest := common2.ClaimRequest{
		IssuerDid:         offer.IssuerDid,
		HolderDid:         callbackResponse.From,
		ExpiredAt:         time.Now().Add(time.Duration(999999)), //TODO
		DataCategory:      template.DataCategory,
		SubCategory:       template.SubCategory,
		LowerBoundInclude: lower.Include,
		LowerBoundValue:   lower.Value,
		UpperBoundInclude: upper.Include,
		UpperBoundValue:   upper.Value,
		HolderAddress:     walletAddress.String(),
	}
	logging.Info.Println("start issue claim")
	claim, err := IssueClaim(&claimRequest)
	if err != nil {
		return "", "", err
	}
	logging.Info.Println("return claim")
	binary, err := claim.MarshalBinary()
	if err != nil {
		return "", "", err
	}
	return hexutil.Encode(binary), offer.IssuerDid, nil
}

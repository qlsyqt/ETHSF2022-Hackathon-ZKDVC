package services

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	circuits "github.com/iden3/go-circuits"
	core "github.com/iden3/go-iden3-core"
	"github.com/iden3/go-merkletree-sql"
	mtsql "github.com/iden3/go-merkletree-sql/db/sql"
	"github.com/iden3/go-rapidsnark/prover"
	"github.com/iden3/go-rapidsnark/types"
	"github.com/iden3/iden3comm/protocol"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"time"
	"wallet-end/config"
	"wallet-end/models/database"
	response2 "wallet-end/models/response"
	"wallet-end/pkgs/db"
	"wallet-end/pkgs/witness"
	"wallet-end/store"
	"wallet-end/utils"
)

func Authentication(request *protocol.AuthorizationRequestMessage) error {
	//Santity check
	if myDid == nil {
		_, err := FetchIdentity()
		if err != nil {
			return err
		}
	}
	cfg := config.GetConfig()
	walletAddress := common.HexToAddress(utils.AddressFromPrivateKey(cfg.PolygonWallet.MainPrivateKey))
	//获取challenge数据
	challenge := request.Body.Scope[0].Rules["challenge"].(string)
	challengeBytes, err := hexutil.Decode(challenge)
	if err != nil {
		return err
	}
	challengeDigest, err := utils.HashChallengeAndWallet(challengeBytes, walletAddress)
	if err != nil {
		return err
	}
	fmt.Println("challenge digest :", challengeDigest.String())
	//Load identity from  db
	didData, err := store.SelectDidById(myDid.String())
	if err != nil {
		return err
	}
	//load trees
	ctx := context.Background()
	cltTree, rotTree, torTree, err := loadTrees(ctx, didData)
	log.Println("Tree loaded")
	fmt.Println("claims root:")
	fmt.Println(cltTree.Root().Hex())
	cltRoot, rotRoot, torRoot := cltTree.Root(), rotTree.Root(), torTree.Root()
	identityState, _ := merkletree.HashElems(cltRoot.BigInt(), rotRoot.BigInt(), torRoot.BigInt())
	//load auth claims
	authClaim, authClaimExistsProof, authClaimNonRevProof, err := loadAuthClaimWithProof(didData.AuthClaimHi, cltTree, rotTree)
	if err != nil {
		return err
	}
	//Build zk proof input
	zkInput, err := createZkInput(
		cfg.PolygonId.PrivateKey, challengeDigest,
		identityState, cltRoot, rotRoot, torRoot,
		authClaim, authClaimExistsProof, authClaimNonRevProof)
	if err != nil {
		return err
	}
	fmt.Println("zk input:")
	inputBytes, err := zkInput.InputsMarshal()
	if err != nil {
		return err
	}
	fmt.Println(string(inputBytes))
	//Create proof
	zkProof, err := createProof(zkInput, cfg)
	if err != nil {
		return err
	}
	fmt.Println("zk proof compulete")
	//Sign data using polygon wallet
	walletSignature := utils.SignDigest(challengeDigest.Bytes(), cfg.PolygonWallet.MainPrivateKey)
	//Create & send response
	response := createResponse(request, zkProof, walletSignature, walletAddress.Hex())
	claim, issuer, err := sendResponse(response, request.Body.CallbackURL)
	if err != nil {
		return err
	}
	//Save claim to db and update my state
	_, err = DoAddClaim(claim, issuer)
	if err != nil {
		return err
	}
	return nil
}

func loadAuthClaimWithProof(authClaimHi []byte, cltTree *merkletree.MerkleTree, rotTree *merkletree.MerkleTree) (*core.Claim, *merkletree.Proof, *merkletree.Proof, error) {
	authClaimData, err := store.SelectClaimByHIndex(new(big.Int).SetBytes(authClaimHi))
	if err != nil {
		return nil, nil, nil, err
	}
	authClaim := &core.Claim{}
	err = authClaim.UnmarshalBinary(authClaimData.ClaimBinary)
	if err != nil {
		return nil, nil, nil, err
	}
	authClaimExistProof, _, _ := cltTree.GenerateProof(context.Background(), new(big.Int).SetBytes(authClaimData.ClaimHi), cltTree.Root())
	authClaimNonRevProof, _, _ := rotTree.GenerateProof(context.Background(), new(big.Int).SetUint64(authClaim.GetRevocationNonce()), rotTree.Root())

	return authClaim, authClaimExistProof, authClaimNonRevProof, nil
}

func createZkInput(polygonIdKey string, challengeDigest *big.Int, didStatus *merkletree.Hash, cltRoot, rotRoot, torRoot *merkletree.Hash,
	authClaim *core.Claim, authClaimExistsProof, authClaimNonRevProof *merkletree.Proof) (*circuits.AuthInputs, error) {
	bjjPrivateKey, _, err := utils.HexPrivateKeyToBjjKeypair(polygonIdKey)
	if err != nil {
		return nil, err
	}
	polygonIdSignature := bjjPrivateKey.SignPoseidon(challengeDigest)

	authKnn3Input := circuits.AuthInputs{
		BaseConfig: circuits.BaseConfig{},
		ID:         myDid,
		AuthClaim: circuits.Claim{
			TreeState: circuits.TreeState{
				State:          didStatus,
				ClaimsRoot:     cltRoot,
				RevocationRoot: rotRoot,
				RootOfRoots:    torRoot,
			},
			Claim: authClaim,
			Proof: authClaimExistsProof,
			NonRevProof: &circuits.ClaimNonRevStatus{
				Proof: authClaimNonRevProof,
			},
		},
		Signature: polygonIdSignature,
		Challenge: challengeDigest,
	}

	return &authKnn3Input, nil
}

func createProof(input *circuits.AuthInputs, cfg *config.Config) (*types.ZKProof, error) {
	fmt.Println("start create proof")
	authProofConfig := cfg.Zkp["auth"]
	wasm, err := ioutil.ReadFile(authProofConfig.WasmPath)
	if err != nil {
		return nil, err
	}

	calculator, err := witness.NewCircom2WitnessCalculator(wasm, false)
	if err != nil {
		return nil, err
	}

	inputBytes, err := input.InputsMarshal()
	if err != nil {
		return nil, err
	}

	parsedInput, err := witness.ParseInputs(inputBytes)
	if err != nil {
		return nil, err
	}

	wtns, err := calculator.CalculateWTNSBin(parsedInput, false)
	if err != nil {
		return nil, err
	}

	zKey, err := ioutil.ReadFile(authProofConfig.ZkeyPath)
	if err != nil {
		panic(err)
	}

	//wtns, _ := ioutil.ReadFile("zk/witness.wtns")
	proof, err := prover.Groth16Prover(zKey, wtns)
	if err != nil {
		return nil, err
	}

	return proof, nil
}

func createResponse(request *protocol.AuthorizationRequestMessage, zkProof *types.ZKProof, walletSignature []byte, walletAddress string) *protocol.AuthorizationResponseMessage {
	respScopes := make([]protocol.ZeroKnowledgeProofResponse, 0)
	respScopes = append(respScopes,
		protocol.ZeroKnowledgeProofResponse{
			ID:        0,
			CircuitID: "auth",
			ZKProof:   *zkProof},

		protocol.ZeroKnowledgeProofResponse{
			ID:        1,
			CircuitID: "knn3",
			ZKProof: types.ZKProof{
				Proof:      nil,
				PubSignals: []string{hexutil.Encode(walletSignature), walletAddress},
			},
		},
	)

	response := protocol.AuthorizationResponseMessage{
		ID:       request.ID,
		Typ:      request.Typ,
		Type:     request.Type,
		ThreadID: "",
		Body: protocol.AuthorizationMessageResponseBody{
			Scope: respScopes,
		},
		From: myDid.String(),
		To:   request.From,
	}
	return &response

}

func sendResponse(payload *protocol.AuthorizationResponseMessage, callback string) (*core.Claim, string, error) {
	log.Println("send payload")
	respBytes, _ := json.Marshal(payload)
	log.Println(string(respBytes))
	bodyBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, "", err
	}
	fmt.Println("callback :", callback)
	req, err := http.NewRequest("POST", callback, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return nil, "", err
	}
	client := http.Client{Timeout: time.Duration(120) * time.Second}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return nil, "", err
	}
	if resp.StatusCode != 200 {
		return nil, "", errors.New("Call back returns " + resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, "", err
	}
	typedResponse := response2.Response{}
	err = json.Unmarshal(body, &typedResponse)
	if err != nil {
		return nil, "", err
	}
	if typedResponse.Code != response2.SUCCESS {
		return nil, "", errors.New("Call back says: " + typedResponse.Message)
	}
	m := typedResponse.Result.(map[string]interface{})
	claimStr := m["claim"].(string)
	issuer := m["issuer"].(string)
	log.Printf("claim readed %s, issuer %s", claimStr, issuer)

	ret := &core.Claim{}
	decoded, err := hexutil.Decode(claimStr)
	if err != nil {
		return nil, "", err
	}
	err = ret.UnmarshalBinary(decoded)
	if err != nil {
		return nil, "", err
	}
	return ret, issuer, nil
}

func loadTrees(ctx context.Context, identity *database.DidData) (*merkletree.MerkleTree, *merkletree.MerkleTree, *merkletree.MerkleTree, error) {
	log.Printf("start loading trees for %s", identity.Did)
	cltTree, err := loadTree(ctx, identity.CltId)
	if err != nil {
		return nil, nil, nil, err
	}
	rotTree, err := loadTree(ctx, identity.RotId)
	if err != nil {
		return nil, nil, nil, err
	}
	torTree, err := loadTree(ctx, identity.TorId)
	if err != nil {
		return nil, nil, nil, err
	}
	return cltTree, rotTree, torTree, nil
}

func loadTree(ctx context.Context, mtId uint64) (*merkletree.MerkleTree, error) {
	tree, err := merkletree.NewMerkleTree(ctx, mtsql.NewSqlStorage(db.GetPgDb(), mtId), 32)
	if err != nil {
		log.Printf("Failed to reload tree :%d", mtId)
		return nil, err
	}
	log.Println("load tree with root: ", tree.Root().Hex())
	return tree, nil
}

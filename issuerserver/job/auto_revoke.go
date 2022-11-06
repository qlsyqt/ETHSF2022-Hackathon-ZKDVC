package job

import (
	"context"
	core "github.com/iden3/go-iden3-core"
	"github.com/juju/errors"
	"github.com/robfig/cron"
	"issuerserver/config"
	"issuerserver/models/database"
	"issuerserver/pkg/codec"
	"issuerserver/pkg/dvc"
	"issuerserver/pkg/logging"
	"issuerserver/pkg/orm"
	"issuerserver/services"
	"issuerserver/store"
	"issuerserver/utils"
	"math/big"
	"strings"
)

var lastPkid int64 = int64(-1)

// Check dvc perioradically and revoke dcp claim if conditions not holding any more
func StartAutoRevoke() {
	orm.Init()
	//捞取所有claim
	//向dvc查询用户的地址信息
	jobCfg := config.GetConfig().Jobs["autoRevoke"]
	c := cron.New()
	cronExpress := jobCfg.Cron
	c.AddFunc(cronExpress, autoRevoke)
	c.Start()
	//select {} //This will stuck here. But can't stuck here because we have other things to do later
}

func autoRevoke() {
	//Load claims
	for true {
		claims, err := store.SelectClaimsByBatch(lastPkid, 200)
		if err != nil {
			logging.Error.Println("Failed to retrieve claims data", err)
			return
		}
		//Exit point
		if len(claims) == 0 {
			logging.Info.Println("No more revokable claims")
			return
		}
		logging.Info.Printf("%d fetched", len(claims))
		//claims is already sorted by offer id
		for _, c := range claims {
			if !c.Revoked && strings.Compare(c.ClaimType, "dvc") == 0 {
				if expired, err := dvcCheck(c); err != nil && expired {
					if err = revoke(c); err != nil {
						logging.Error.Println("revoke failed", err)
						return
					} else {
						logging.Info.Printf("%d revoked success", new(big.Int).SetBytes(c.ClaimHi))
					}
				}
			}
			lastPkid = c.Id
		}
	}
}

func dvcCheck(claimDbObj *database.ClaimData) (bool, error) {
	ctx := context.Background()
	//To claim core data
	claim := &core.Claim{}
	if err := claim.UnmarshalBinary(claimDbObj.ClaimBinary); err != nil {
		return false, errors.Trace(err)
	}
	return false, nil

	//Query DVC
	dcp := codec.DecodeDcpClaim(claim)
	claimValue, err := dvc.Query(ctx, dcp.DataCategory, dcp.SubCategory, dcp.HolderAddress)
	if err != nil {
		return false, nil
	}
	//Judge result
	return utils.InRange(claimValue, &dcp.LowerBoundary, &dcp.UpperBoundary), nil
}

func revoke(c *database.ClaimData) error {
	return services.RevokeClaim(new(big.Int).SetBytes(c.ClaimHi))
}

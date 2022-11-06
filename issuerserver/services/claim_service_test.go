package services

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"issuerserver/models/common"
	"issuerserver/models/database"
	"issuerserver/store"
	"issuerserver/utils"
	"os"
	"testing"
	"time"
)

func TestIssueClaim(t *testing.T) {
	os.Chdir("..")
	cleanTablesForDbg()
	issuer, err := CreateIdentity("aaa")
	if err != nil {
		panic(err)
	}
	fmt.Println("issuer created")
	holder := utils.RandomDid()
	_, holderWallet := utils.RandomECDSAKeyPair()

	fmt.Printf("generating holder %s,  %s", holder.String(), holderWallet.Hex())
	claimRequest := common.ClaimRequest{
		IssuerDid:         issuer,
		HolderDid:         holder.String(),
		ExpiredAt:         time.Now().Add(time.Duration(9999999)),
		DataCategory:      "snapshot",
		SubCategory:       "gitcoin.eth",
		LowerBoundInclude: true,
		LowerBoundValue:   10,
		UpperBoundInclude: false,
		UpperBoundValue:   20,
		HolderAddress:     holderWallet.Hex(),
	}
	_, err = IssueClaim(&claimRequest)
	if err != nil {
		panic(err)
	}
	claimRequest = common.ClaimRequest{
		IssuerDid:         issuer,
		HolderDid:         holder.String(),
		ExpiredAt:         time.Now().Add(time.Duration(9999999)),
		DataCategory:      "pagerank",
		SubCategory:       "gitcoin.eth",
		LowerBoundInclude: true,
		LowerBoundValue:   10,
		UpperBoundInclude: false,
		UpperBoundValue:   20,
		HolderAddress:     holderWallet.Hex(),
	}
	_, err = IssueClaim(&claimRequest)
	if err != nil {
		panic(err)
	}
	fmt.Println("issue claim complete")
}

func TestRevoke(t *testing.T) {
	os.Chdir("..")
	issuer, err := CreateIdentity("bvbb")
	if err != nil {
		panic(err)
	}
	fmt.Println("issuer created")
	holder := utils.RandomDid()
	_, holderWallet := utils.RandomECDSAKeyPair()

	fmt.Printf("generating holder %s,  %s", holder.String(), holderWallet.Hex())
	claimRequest := common.ClaimRequest{
		IssuerDid:         issuer,
		HolderDid:         holder.String(),
		ExpiredAt:         time.Now().Add(time.Duration(9999999)),
		DataCategory:      "snapshot",
		SubCategory:       "gitcoin.eth",
		LowerBoundInclude: true,
		LowerBoundValue:   10,
		UpperBoundInclude: false,
		UpperBoundValue:   20,
		HolderAddress:     holderWallet.Hex(),
	}
	claim, err := IssueClaim(&claimRequest)

	if err != nil {
		panic(err)
	}

	hi, _, _ := claim.HiHv()
	err = RevokeClaim(hi)
	if err != nil {
		panic(err)
	}
	fmt.Println("finish")
}

func TestPrepareData(t *testing.T) {
	//!!!!!请确认调用时连接的是本地数据库
	os.Chdir("..")
	//clean table
	//cleanTablesForDbg()
	//create an identity
	//CreateIdentity("account@knn3")

	did, err := GetDidByName("account@knn3")
	if err != nil {
		panic(err)
	}

	//build template
	template := &database.Template{
		Id:              0,
		TemplateName:    "aaaa",
		DataCategory:    "1",
		SubCategory:     "0x495f947276749ce646f68ac8c248420045cb7b5e",
		IsExpirable:     false,
		IsAutoRevokable: false,
		Classfications:  "",
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
	}
	err = store.InsertTemplate(template)
	if err != nil {
		panic(err)
	}
	//build offer
	preclaims := make([]*common.PreClaim, 2)
	preclaims[0] = &common.PreClaim{
		Name:           "1111",
		DataCategory:   "ddd",
		LowerBound:     []int{1, 1, 0, 0},
		UpperBound:     []int{1, 0, 10, 0},
		CreateDate:     time.Time{},
		ExpirationDate: "",
	}
	preclaims[1] = &common.PreClaim{
		Name:           "1111",
		DataCategory:   "ddd",
		LowerBound:     []int{1, 1, 10, 0},
		UpperBound:     []int{1, 0, 20, 0},
		CreateDate:     time.Time{},
		ExpirationDate: "",
	}
	ps, err := json.Marshal(&preclaims)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(ps))
	//TODO：offer缺isserDid字段
	offer := &database.Offer{
		Id:         0,
		Name:       "aaa",
		TemplateId: template.Id,
		PreClaims:  string(ps),
		IssuerDid:  did,
		Link:       "",
		CreatedAt:  time.Time{},
		UpdatedAt:  time.Time{},
	}
	err = store.InsertOffer(offer)
	if err != nil {
		panic(err)
	}

	//启动链

}

func cleanTablesForDbg() {
	mysql1 := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		"root", "WMmpmHawul", "tidb.uunmqwe9i4u.clusters.tidb-cloud.com", 4000, "hackthon")
	mysql2 := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		"root", "WMmpmHawul", "tidb.uunmqwe9i4u.clusters.tidb-cloud.com", 4000, "hackthonw")

	//mysql1 := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
	//	"root", "12345678", "127.0.0.1", 3306, "hackthon")
	//mysql2 := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
	//	"root", "12345678", "127.0.0.1", 3306, "hackthonw")

	mysqlDb1, _ := sql.Open("mysql", mysql1)
	mysqlDb2, _ := sql.Open("mysql", mysql2)
	//
	_, err := mysqlDb1.Exec("drop table if exists t_did ;")
	if err != nil {
		panic(err)
	}
	_, err = mysqlDb1.Exec("drop table if exists t_claim;")
	if err != nil {
		panic(err)
	}
	_, err = mysqlDb1.Exec("drop table if exists offer;")
	if err != nil {
		panic(err)
	}
	_, err = mysqlDb1.Exec("drop table if exists template;")
	if err != nil {
		panic(err)
	}

	_, err = mysqlDb2.Exec("drop table if exists t_did ;")
	if err != nil {
		panic(err)
	}
	_, err = mysqlDb2.Exec("drop table if exists t_claim;")
	if err != nil {
		panic(err)
	}
	_, err = mysqlDb2.Exec("drop table if exists offer;")
	if err != nil {
		panic(err)
	}
	_, err = mysqlDb2.Exec("drop table if exists template;")
	if err != nil {
		panic(err)
	}

	pg1 := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"127.0.0.1", 5432, "root", "12345678", "hackthon")
	pg2 := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"127.0.0.1", 5432, "root", "12345678", "hackthonw")
	pgdb1, _ := sql.Open("postgres", pg1)
	pgdb2, _ := sql.Open("postgres", pg2)

	_, err = pgdb1.Exec("drop table if exists mt_nodes;")
	if err != nil {
		panic(err)
	}
	_, err = pgdb1.Exec("drop table if exists mt_roots;")
	if err != nil {
		panic(err)
	}
	_, err = pgdb2.Exec("drop table if exists mt_nodes;")
	if err != nil {
		panic(err)
	}
	_, err = pgdb2.Exec("drop table if exists mt_roots;")
	if err != nil {
		panic(err)
	}
	fmt.Println("本地数据库清理完毕")
}

func TestPrepareTidb(t *testing.T) {
	cleanTablesForRemote()
}

func cleanTablesForRemote() {
	tidbStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		"root", "WMmpmHawul", "tidb.uunmqwe9i4u.clusters.tidb-cloud.com", 4000, "polygonid")
	mysql2 := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		"root", "12345678", "127.0.0.1", 3306, "hackthonw")
	tidb, _ := sql.Open("mysql", tidbStr)
	mysqlDb2, _ := sql.Open("mysql", mysql2)
	//

	_, err := tidb.Exec("drop table if exists t_claim;")
	if err != nil {
		panic(err)
	}

	_, err = mysqlDb2.Exec("drop table if exists t_did ;")
	if err != nil {
		panic(err)
	}
	_, err = mysqlDb2.Exec("drop table if exists t_claim;")
	if err != nil {
		panic(err)
	}
	pg1 := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"127.0.0.1", 5432, "root", "12345678", "hackthon")
	pg2 := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"127.0.0.1", 5432, "root", "12345678", "hackthonw")
	pgdb1, _ := sql.Open("postgres", pg1)
	pgdb2, _ := sql.Open("postgres", pg2)

	_, err = pgdb1.Exec("drop table if exists mt_nodes;")
	if err != nil {
		panic(err)
	}
	_, err = pgdb1.Exec("drop table if exists mt_roots;")
	if err != nil {
		panic(err)
	}
	_, err = pgdb2.Exec("drop table if exists mt_nodes;")
	if err != nil {
		panic(err)
	}
	_, err = pgdb2.Exec("drop table if exists mt_roots;")
	if err != nil {
		panic(err)
	}
	fmt.Println("本地数据库清理完毕")
}

//
//func TestBuildClaim(t *testing.T) {
//	os.Chdir("..")
//	issuerId := utils.RandomDid()
//	holderId := utils.RandomDid()
//
//	claimInfo := &common.ClaimRequest{
//		IssuerDid:         issuerId.String(),
//		HolderDid:         holderId.String(),
//		ExpiredAt:         time.Now().Add(time.Duration(9999999)),
//		DataCategory:      "Snapshot",
//		SubCategory:       "aabbcc.eth",
//		LowerBoundInclude: true,
//		LowerBoundValue:   0,
//		UpperBoundInclude: false,
//		UpperBoundValue:   5,
//	}
//
//	claim, err := generateClaimInfo(claimInfo)
//	assert.Equal(t, err, nil)
//	assert.NotEqual(t, claim, nil)
//}
//
//func TestSaveClaim(t *testing.T) {
//	os.Chdir("..")
//
//	did, err := CreateIdentity()
//	if err != nil {
//		panic(err)
//	}
//
//	issuerId := did
//	holderId := utils.RandomDid()
//
//	claimInfo := &common.ClaimRequest{
//		IssuerDid:         issuerId,
//		HolderDid:         holderId.String(),
//		ExpiredAt:         time.Now().Add(time.Duration(9999999)),
//		DataCategory:      "snapshot",
//		SubCategory:       "aabbcc.eth",
//		LowerBoundInclude: true,
//		LowerBoundValue:   0,
//		UpperBoundInclude: false,
//		UpperBoundValue:   5,
//	}
//
//	err = IssueClaim(claimInfo)
//	if err != nil {
//		panic(err)
//	}
//}

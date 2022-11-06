package services

import (
	"os"
	"testing"
)

func TestCreateIdentity(t *testing.T) {
	os.Chdir("..")
	cleanTablesForDbg()
	//did, err := CreateIdentity("account@knn3")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(did)
}

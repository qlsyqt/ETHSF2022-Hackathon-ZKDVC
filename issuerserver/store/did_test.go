package store

//func TestInsertAndSelectDid(t *testing.T) {
//	os.Chdir("../")
//	did := database.DidData{
//		Did:            "aaaa",
//		AuthPrivateKey: []byte{1, 1, 1},
//		AuthClaimHi:    []byte{2, 2, 2},
//		IsGenesis:      true,
//		CltId:          1,
//		RotId:          2,
//		TorId:          3,
//		CreatedAt:      time.Now(),
//		UpdatedAt:      time.Now(),
//	}
//
//	err := InsertDid(&did)
//	fmt.Println(err)
//
//	r, err := SelectDidById(did.Did)
//	fmt.Println(err)
//	fmt.Println(r)
//}

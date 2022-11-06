package dvc

import (
	"context"
	"fmt"
	"os"
	"testing"
)

func TestQuerySnapshot(t *testing.T) {
	os.Chdir("../..")
	//https://credentials.knn3.xyz/nft/0x64e16D972Dac15d0700764f64C9011432d59A79C
	cnt, err := Query(context.Background(), "2", "0xcrystal.eth", "0x64e16D972Dac15d0700764f64C9011432d59A79C")
	if err != nil {
		panic(err)
	}
	fmt.Println(cnt)
}

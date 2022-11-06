package dvc

import (
	"context"
	"fmt"
	"os"
	"testing"
)

func TestQueryNft(t *testing.T) {
	os.Chdir("../..")
	//https://credentials.knn3.xyz/nft/0x64e16D972Dac15d0700764f64C9011432d59A79C
	cnt, err := Query(context.Background(), "1", "0x495f947276749Ce646f68AC8c248420045cb7b5e", "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	if err != nil {
		panic(err)
	}
	fmt.Println(cnt)
}

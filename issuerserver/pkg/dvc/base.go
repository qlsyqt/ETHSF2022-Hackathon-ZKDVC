package dvc

import (
	"context"
	"errors"
	browser "github.com/EDDYCJY/fake-useragent"
	"io"
	"issuerserver/config"
	"issuerserver/utils"
	"net/http"
	"time"
)

//type DvcQuery interface {
//	Query(ctx context.Context, subCategory string, wallet string) (int, error)
//}

var queryStubs map[string]func(ctx context.Context, subCategory, wallet string) (int, error)
var encodeStubs map[string]func(subCategory string) []byte
var decodeStubs map[string]func(subCategory []byte) string

func Query(ctx context.Context, dataCategory, subCategory, wallet string) (int, error) {
	retry := 10

	if f, ok := queryStubs[dataCategory]; ok {
		for retry > 0 {
			val, err := f(ctx, subCategory, wallet)
			if err == nil {
				return val, err
			}
			retry--
		}

	}

	return 0, errors.New("data category not found:" + dataCategory)
}

func EncodeSubcategory(dataCategory, subCategory string) []byte {
	if f, ok := encodeStubs[dataCategory]; ok {
		return f(subCategory)
	}
	return []byte{}
}

func DecodeSubcategory(dataCategory string, subCategory []byte) string {
	if f, ok := decodeStubs[dataCategory]; ok {
		return f(subCategory)
	}
	return ""
}

func sendRequestAndGetResponse(ctx context.Context, dataCategory string, subCategory string, wallet string) ([]byte, error) {
	//Url
	cfg := config.GetConfig()
	url, err := utils.Concat(cfg.Dvc.Host, dataCategory, wallet)
	if err != nil {
		return nil, err
	}
	//Create request
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", browser.Random())
	client := http.Client{
		Timeout: time.Duration(cfg.Dvc.ReadTimeout) * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	//knn3 would return 404 if record not found, which is a valid scenario
	if resp.StatusCode == 404 {
		return nil, errors.New("404 not found")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("Invalid status " + resp.Status)
	}
	return body, nil
}

func init() {
	queryStubs = make(map[string]func(ctx context.Context, subCategory, wallet string) (int, error))
	queryStubs["0"] = ensQuery
	queryStubs["1"] = nftQuery
	queryStubs["2"] = snapshotQuery

	encodeStubs = make(map[string]func(subCategory string) []byte)
	encodeStubs["0"] = ensEncodeSubcategory
	encodeStubs["1"] = nftEncodeSubcategory
	encodeStubs["2"] = snapshotEncodeSubcategory

	decodeStubs = make(map[string]func(subCategory []byte) string)
	decodeStubs["0"] = ensDecodeSubcategory
	decodeStubs["1"] = nftDecodeSubcategory
	decodeStubs["2"] = snapshotDecodeSubcategory

}

package geminiapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Api struct {
	SandBox bool
	Key     string
	Secret  string
	BaseURL string
}

func (api *Api) GetGeminiEnv() EnvResponse {
	var envRes EnvResponse
	if api.SandBox {
		envRes.Message = "IN SANDBOX"
	}

	envRes.Message = "NOT IN SANDBOX"

	return envRes
}

func (api *Api) GetSymbols() ([]string, error) {
	symbolsURIPath := "v1/symbols"

	url := api.BaseURL + symbolsURIPath

	fmt.Printf("url is %v", url)

	// setup the request
	req, err := http.NewRequest(http.MethodGet, url, bytes.NewBuffer([]byte{}))

	if err != nil {
		fmt.Printf("ERROR IN CLIENT DO is %v\n", err)
		return nil, err
	}

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	// catch errors in the response
	if res.StatusCode >= 400 && res.StatusCode <= 499 {
		return nil, fmt.Errorf("%v\n%s", res.StatusCode, "something went wrong")
	}

	defer res.Body.Close()

	// read response body
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Printf("ERROR IN READ ALL is %v\n", err)

		return nil, err
	}

	// var symbolsDAO []string
	var symbolsDAO []string

	// turn body data into struc
	if err := json.Unmarshal(body, &symbolsDAO); err != nil {
		fmt.Printf("ERROR IN UNMARSHALL is %v\n", err)

		return nil, err
	}

	return symbolsDAO, nil

}

func (api *Api) GetCoinTicker(coinSymbol string) (TickerV2DAO, error) {
	tickerURIPath := "v2/ticker/"

	url := api.BaseURL + tickerURIPath + coinSymbol

	// setup the request
	req, err := http.NewRequest(http.MethodGet, url, bytes.NewBuffer([]byte{}))

	if err != nil {

		return TickerV2DAO{}, err
	}

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return TickerV2DAO{}, err
	}

	// catch errors in the response
	if res.StatusCode >= 400 && res.StatusCode <= 499 {
		return TickerV2DAO{}, fmt.Errorf("%v\n%s", res.StatusCode, "something went wrong")
	}

	defer res.Body.Close()

	// read response body
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return TickerV2DAO{}, err
	}

	var tickerDAO TickerV2DAO

	// turn body data into struc
	if err := json.Unmarshal(body, &tickerDAO); err != nil {
		return tickerDAO, err
	}

	return tickerDAO, nil

}

func New(isSandBox bool, key string, secret string) *Api {

	baseURL := "https://api.gemini.com/"

	if isSandBox {
		baseURL = "https://api.sandbox.gemini.com/"
	}

	return &Api{
		SandBox: isSandBox,
		Key:     key,
		Secret:  secret,
		BaseURL: baseURL,
	}
}

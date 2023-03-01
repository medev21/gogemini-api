package geminiapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Api struct {
	url    string
	key    string
	secret string
}

type TickerV2DAO struct {
	Symbol  string   `json:"symbol"`
	Open    float64  `json:"open,string"`
	High    float64  `json:"high,string"`
	Low     float64  `json:"low,string"`
	Close   float64  `json:"close,string"`
	Changes []string `json:"changes"`
	Bid     float64  `json:"bid,string"`
	Ask     float64  `json:"ask,string"`
}

func (api *Api) GetCoinTicker(coinSymbol string) (TickerV2DAO, error) {
	tickerURI := "ticker/"
	baseURL := "https://api.gemini.com/v2/"

	url := baseURL + tickerURI + coinSymbol

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
		return TickerV2DAO{}, err
	}

	return tickerDAO, nil

}

func New(url string, key string, secret string) *Api {
	return &Api{
		url:    url,
		key:    key,
		secret: secret,
	}
}

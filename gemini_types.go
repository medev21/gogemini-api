package geminiapi

type EnvResponse struct {
	Message string `json:"message"`
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

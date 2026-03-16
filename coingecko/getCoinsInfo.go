package coingecko

import (
	"encoding/json"
	"net/http"
)

type CoinResult struct {
	MarketCap             float64 `json:"market_cap"`
	MarketCapRank         int64   `json:"market_cap_rank"`
	TotalVolume           float64 `json:"total_volume"`
	High24                float64 `json:"high_24h"`
	Low24                 float64 `json:"low_24h"`
	PriceChangePercentage float64 `json:"price_change_percentage_24h"`
	PriceChange           float64 `json:"price_change_24h"`
	Coins                 float64 `json:"circulating_supply"`
	Image                 string  `json:"image"`
}

func GetCoinInfo(pair string) CoinResult {
	var Result []CoinResult
	url := "https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&ids=" + pair

	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&Result)

	if err != nil {
		panic(err)
	}
	return Result[0]
}

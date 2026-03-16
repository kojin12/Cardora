package handlers

import (
	"encoding/json"
	"log"
	"main/coingecko"
	"main/config"
	"net/http"
)

type CoinResponse struct {
	CoinInfo coingecko.CoinResult `json:"coin_info"`
}

func GetCoinInfoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	url := r.URL.Query()

	coin := url.Get("coin")
	log.Printf("Received parameters: coin=%s", coin)

	if coin == "" {
		http.Error(w, "coin is required", http.StatusBadRequest)
		return
	}
	var Result CoinResponse
	coinSymbol, ok := config.CoinConfig[coin]
	if !ok {
		http.Error(w, "unknown pair", http.StatusBadRequest)
		return
	}

	Result.CoinInfo = coingecko.GetCoinInfo(coinSymbol.CoinGecko)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Result)

}

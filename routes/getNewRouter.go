package routes

import (
	"main/handlers"
	"net/http"
)

func GetNewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/data/info", handlers.GetDataHandlers)
	mux.HandleFunc("/data/coin_info", handlers.GetCoinInfoHandler)
	return mux
}

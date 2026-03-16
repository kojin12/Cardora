package main

import (
	"log"
	routes "main/routes"
	"net/http"
)

func main() {
	router := routes.GetNewRouter()
	
	log.Println("server started :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

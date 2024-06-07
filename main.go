package main

import (
	"Biblioteca/routers"
	"Biblioteca/utils"
	"log"
	"net/http"
)

func main() {
	utils.InitDB()
	router := routers.InitRouter()
	log.Fatal(http.ListenAndServe(":8000", router))
}

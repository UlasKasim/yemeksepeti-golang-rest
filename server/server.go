package server

import (
	"log"
	"net/http"
	key_value "yemeksepeti-golang-rest/module/key_value"
)

func StartServer() {
	//Initialize key-value http server
	key_value.Initialize()

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

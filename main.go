package main

import (
	"log"
	"sync"
	"yemeksepeti-golang-rest/server"
)

func main() {
	//Initialize a wait group for server
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		log.Println("Starting server")

		//Simply start server
		server.StartServer()

	}()

	wg.Wait()

	log.Print("Server stopped")
}

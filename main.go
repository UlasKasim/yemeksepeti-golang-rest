package main

import (
	"context"
	"log"
	"sync"
	"yemeksepeti-golang-rest/server"
)

func main() {
	context.Background()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		log.Println("Starting server")

		server.StartServer()

		wg.Done()
	}()

	wg.Wait()

	log.Print("Server stopped")
}

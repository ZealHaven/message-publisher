package main

import (
	"log"
	"net/http"

	"github.com/zeal-haven/message-publisher/router"
)

func main() {
	r := router.Init()

	// Start at port 3001
	log.Printf("Server starting...")
	log.Fatal(http.ListenAndServe(":3001", r))
}

package main

import (
	"flag"
	"fmt"
	"log"
	"url-shortner/api"
	"url-shortner/api/routes"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
        fmt.Println("Error loading .env file")
    }
	router := routes.NewRouter()
	listenAddr := flag.String("listen-addr", ":3000", "server listen address")
	flag.Parse()

	server := api.NewServer(*listenAddr, router)
	log.Fatal(server.Start())
}

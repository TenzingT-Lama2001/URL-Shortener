package main

import (
	"flag"
	"log"
	"url-shortner/api"
	"url-shortner/api/routes"
)

func main() {

	router := routes.NewRouter()
	listenAddr := flag.String("listen-addr", ":3000", "server listen address")
	flag.Parse()

	server := api.NewServer(*listenAddr, router)
	log.Fatal(server.Start())
}

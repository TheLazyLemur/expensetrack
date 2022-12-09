package main

import (
	"flag"
	"log"
)

func main() {
	port := flag.String("port", ":8080", "port to listen on")
	flag.Parse()

	userStore, err := NewPostgresUserStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := userStore.Migrate(); err != nil {
		log.Fatal(err)
	}

	server := NewAPIServer(*port, userStore)
	log.Fatal(server.Run())
}

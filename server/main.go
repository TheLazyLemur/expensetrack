package main

import (
	"flag"
	"log"
	"sync"
)

var (
	reportMutex sync.Mutex
)

func main() {
	port := flag.String("port", ":8080", "port to listen on")
	dbString := flag.String("db", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable", "database connection string")
	flag.Parse()

	postgresStore, err := NewPostgresStore(dbString)
	if err != nil {
		log.Fatal(err)
	}

	if err := postgresStore.Migrate(); err != nil {
		log.Fatal(err)
	}

	bgServer := NewBackgroundServer(postgresStore)
	bgServer.StartbackgroundTasks()

	server := NewAPIServer(*port, postgresStore)
	log.Fatal(server.Run())
}

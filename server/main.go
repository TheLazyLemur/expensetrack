package main

import (
	"flag"
	"log"
	"sync"
	"time"
)

var (
	reportMutex sync.Mutex
)

func main() {
	StartbackgroundTasks()

	port := flag.String("port", ":8080", "port to listen on")
	flag.Parse()

	postgresStore, err := NewPostgresUserStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := postgresStore.Migrate(); err != nil {
		log.Fatal(err)
	}

	server := NewAPIServer(*port, postgresStore)
	log.Fatal(server.Run())
}

func StartbackgroundTasks() {
	reportTicker := time.NewTicker(5 * time.Second)

	go func() {
		for range reportTicker.C {
			GenerateReport()
		}
	}()
}


package main

import (
	"fmt"
	"log"
)

func GenerateReport() {
	userStore, err := NewPostgresUserStore()
	if err != nil {
		log.Fatal(err)
	}

	if reportMutex.TryLock() {
		log.Println("Report generation started...")

		users, err := userStore.GetUsers()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		for _, user := range users {
			fmt.Println(user)
		}

		log.Println("Report generation completed...")

		reportMutex.Unlock()
	} else {
		return
	}
}

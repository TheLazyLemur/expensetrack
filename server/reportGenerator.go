package main

import (
	"fmt"
	"log"
)

func (s *BackgroundServer) GenerateReport() {
    if reportMutex.TryLock() {
        defer reportMutex.Unlock()
        log.Println("Report generation started...")

        users, err := s.Storer.GetUsers()
        if err != nil {
            fmt.Println(err.Error())
            return
        }

        userIdToExpenses := make(map[int64][]Expense)

        for _, user := range users {
            expenses, err := s.Storer.GetExpensesByUser(user.ID)
            if err != nil {
                fmt.Println(err.Error())
                return
            }

            userIdToExpenses[user.ID] = expenses
        }

        log.Println("Report generation completed...")
    } else {
        log.Println("Lock is curretnly held")
        return
    }
}

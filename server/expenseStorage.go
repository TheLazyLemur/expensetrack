package main

import (
	"io"
	"mime/multipart"
	"os"
)

func (s *PostgresStore) CreateExpense(userId int64, amount int64, description string) error {
    query := `
    INSERT INTO expenses (user_id, amount, description) VALUES ($1, $2, $3);
    `
    _, err := s.db.Exec(query, userId, amount, description)
    return err
}

func (s *PostgresStore) GetExpense(expenseId int64) (*Expense, error) {
    query := `
    SELECT * FROM expenses WHERE id = $1;
    `
    row := s.db.QueryRow(query, expenseId)
    expense := Expense{}

    err := row.Scan(&expense.ID, &expense.UserId, &expense.Amount, &expense.Description, &expense.CreatedAt)
    if err != nil {
        return nil, err
    }

    return &expense, nil
}


func (s *PostgresStore) GetExpensesByUser(userId int64) ([]Expense, error) {
    query := `
    SELECT id, user_id, amount, description, created_at FROM expenses WHERE user_id = $1;
    `
    rows, err := s.db.Query(query, userId)
    if err != nil {
        return nil, err
    }

    defer rows.Close()
    expenses := []Expense{}

    for rows.Next() {
        expense := Expense{}

        err = rows.Scan(&expense.ID, &expense.UserId, &expense.Amount, &expense.Description, &expense.CreatedAt)
        if err != nil {
            return nil, err
        }
        expenses = append(expenses, expense)
    }
    return expenses, err
}

func (s *PostgresStore) UpdateExpense(expenseId int64, amount int64, description string) error {
    query := `
    UPDATE expenses SET amount = $1, description = $2 WHERE id = $3;
    `
    _, err := s.db.Exec(query, amount, description, expenseId)
    return err
}

func (s *PostgresStore) DeleteExpense(expenseId int64) error {
    query := `
    DELETE FROM expenses WHERE id = $1;
    `
    _, err := s.db.Exec(query, expenseId)
    return err
}

func (s *PostgresStore) StoreRecipt(expenseId int64, file multipart.File) error {
    id := GenerateUuid()

    f, err := os.Create("./recipts/" + id)
    if err != nil {
        return err 	
    }

    defer f.Close()

    if _, err = io.Copy(f, file); err != nil {
        return err
    }

    return nil
}

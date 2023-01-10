package main

import (
	"database/sql"
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

func (s *PostgresStore) StoreReceipt(expenseId int64, file multipart.File) error {
    fc := make(chan error)
    dbc := make(chan error)

    tx, err := s.db.Begin()

    if err != nil {
        return err 	
    }

    id := GenerateUuid()

    go func(id string) {
        err := SaveFile(id, file)
        fc <- err
    }(id)

    go func(expenseId int64, id string, tx *sql.Tx){
        err := InsertReceipt(expenseId, id, tx)
        dbc <- err
    }(expenseId, id, tx)


    for i := 0; i < 2; i++ {
        select {
            case err1 := <-fc:
            if err1 != nil {
                tx.Rollback()
                os.Remove("./receipts/" + id)
                return err1 
            }
            case err2 := <-dbc:
            if err2 != nil {
                tx.Rollback()
                os.Remove("./receipts/" + id)
                return err2
            }
        }
    }

    return tx.Commit()
}

func InsertReceipt(expenseId int64, id string, tx *sql.Tx) error {
    query := `
    INSERT INTO receipt (expense_id, file_name) VALUES ($1, $2);
    `

    _, err := tx.Exec(query, expenseId, id)
    return err
}

func SaveFile(id string, file multipart.File) error {
    f, err := os.Create("./receipts/" + id)
    if err != nil {
        return err 	
    }

    if _, err = io.Copy(f, file); err != nil {
        return err
    }

    defer f.Close()

    return nil
}

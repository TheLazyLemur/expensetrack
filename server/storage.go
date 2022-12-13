package main

import (
    "database/sql"
    _ "github.com/lib/pq"
)

type UserStorer interface {
    CreateUser(name string, email string, country string) error
    GetUser(id int64) (*User, error)
    UpdateUser(user *User) error
    DeleteUser(id int64) error
    GetUsers() ([]*User, error)
}

type ExpenseStorer interface {
    CreateExpense(userId int64, amount int64, description string) error
    GetExpensesByUser(userId int64) ([]Expense, error)
    UpdateExpense(expenseId int64, amount int64, description string) error
    DeleteExpense(expenseId int64) error
}

type Storer interface {
    UserStorer
    ExpenseStorer
}

type PostgresStore struct {
    db *sql.DB
}

func NewPostgresStore(dbString *string) (*PostgresStore, error) {
    db, err := sql.Open("postgres", *dbString)
    if err != nil {
        return nil, err
    }

    if err := db.Ping(); err != nil {
        return nil, err
    }

    return &PostgresStore{
        db: db,
    }, nil
}

func (s *PostgresStore) Migrate() error {
    query := `
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        name varchar(255) NOT NULL,
        email varchar(255) NOT NULL,
        country varchar(255) NOT NULL
    );

    CREATE TABLE IF NOT EXISTS expenses (
        id SERIAL PRIMARY KEY,
        user_id int NOT NULL,
        amount int NOT NULL,
        description varchar(255) NOT NULL,
        FOREIGN KEY (user_id) REFERENCES users (id)
    );

    ALTER TABLE expenses ADD COLUMN IF NOT EXISTS created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP;
    ALTER TABLE users ADD COLUMN IF NOT EXISTS created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP;
    `
    _, err := s.db.Exec(query)
    return err
}

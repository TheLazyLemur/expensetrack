package main

type CreateUserRequest struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Country string `json:"country"`
}

type UpdateUserRequest struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Country string `json:"country"`
}

type User struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Country string `json:"country"`
}

type CreateExpenseRequest struct {
	UserId      int64  `json:"user_id"`
	Amount      int64  `json:"amount"`
	Description string `json:"description"`
}

type UpdateExpenseRequest struct {
	Amount      int64  `json:"amount"`
	Description string `json:"description"`
}

type Expense struct {
	ID          int64  `json:"id"`
	UserId      int64  `json:"user_id"`
	Amount      int64  `json:"amount"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
}

func NewUser(id int64, name, email, country string) *User {
	return &User{
		ID:      id,
		Name:    name,
		Email:   email,
		Country: country,
	}
}

func NewExpense(id int64, userId int64, amount int64, description string, createdAt string) *Expense {
	return &Expense{
		ID:          id,
		UserId:      userId,
		Amount:      amount,
		Description: description,
		CreatedAt:   createdAt,
	}
}

package main

func (s *PostgresUserStore) CreateExpense(userId int64, amount int64, description string) error {
	query := `
        INSERT INTO expenses (user_id, amount, description) VALUES ($1, $2, $3);
    `
	_, err := s.db.Exec(query, userId, amount, description)
	return err
}

func (s *PostgresUserStore) GetExpense(expenseId int64) error {
	query := `
        SELECT id, user_id, amount, description, created_at FROM expenses WHERE id = $1;
    `
	_, err := s.db.Exec(query, expenseId)
	return err
}

func (s *PostgresUserStore) UpdateExpense(expenseId int64, amount int64, description string) error {
	query := `
        UPDATE expenses SET amount = $1, description = $2 WHERE id = $3;
    `
	_, err := s.db.Exec(query, amount, description, expenseId)
	return err
}

func (s *PostgresUserStore) DeleteExpense(expenseId int64) error {
	query := `
        DELETE FROM expenses WHERE id = $1;
    `
	_, err := s.db.Exec(query, expenseId)
	return err
}

package main

func (s *PostgresStore) CreateUser(name string, email string, country string) error {
	query := `
		INSERT INTO users (name, email, country) VALUES ($1, $2, $3);
	`
	_, err := s.db.Exec(query, name, email, country)
	return err
}

func (s *PostgresStore) GetUser(id int64) (*User, error) {
	query := `SELECT id, name, email, country FROM users WHERE id = $1;`

	row := s.db.QueryRow(query, id)
	user := &User{}
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Country)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *PostgresStore) GetUsers() ([]*User, error) {
	query := `SELECT id, name, email, country  FROM users;`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	users := []*User{}
	for rows.Next() {
		user := &User{}
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Country)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (s *PostgresStore) UpdateUser(user *User) error {
	query := `UPDATE users SET name = $1, email = $2, country = $3 WHERE id = $4;`

	_, err := s.db.Exec(query, user.Name, user.Email, user.Country, user.ID)
	return err
}

func (s *PostgresStore) DeleteUser(id int64) error {
	query := `DELETE FROM users WHERE id = $1;`
	_, err := s.db.Exec(query, id)
	return err
}

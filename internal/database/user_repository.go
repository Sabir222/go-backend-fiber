package database

import (
	"context"
	"fmt"
	"time"
)

// insert user into the database
func (s *service) AddUser(username, email, password string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	addUserQuery := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3)`
	_, err := s.db.ExecContext(ctx, addUserQuery, username, email, password)
	if err != nil {
		return fmt.Errorf("failed to insert user into database: %v", err)
	}
	return nil
}

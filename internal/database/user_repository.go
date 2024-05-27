package database

import (
	"context"
	"fmt"
	"time"
)

const queryTimeout = 5 * time.Second

func (s *service) AddUser(userName, userEmail, userPassword string) error {
	ctx, cancel := context.WithTimeout(context.Background(), queryTimeout)
	defer cancel()

	userQuery := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3)`
	_, err := s.db.ExecContext(ctx, userQuery, userName, userEmail, userPassword)
	if err != nil {
		return fmt.Errorf("failed to insert user %s into database: %v", userEmail, err)
	}
	return nil
}

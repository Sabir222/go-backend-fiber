// Functions to query insert update delete user from database
package repositories

import (
	"context"
	"database/sql"
	"github.com/sabir222/go-backend-fiber/internal/models"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(ctx context.Context, user *models.User) error {
	query := `INSERT INTO users (full_name, password_hash, email) 
              VALUES ($1, $2, $3) RETURNING id, created_at, updated_at, role`
	row := r.db.QueryRowContext(ctx, query, user.FullName, user.PasswordHash, user.Email)
	if err := row.Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt, &user.Role); err != nil {
		return err
	}
	return nil
}

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

/*

package repository

import (
    "context"
    "database/sql"
    "log"
    "time"

    "github.com/sabir222/go-backend-fiber/internal/models"
)

type UserRepository interface {
    CreateUser(ctx context.Context, user *models.User) error
    GetUserByEmail(ctx context.Context, email string) (*models.User, error)
    GetUserByID(ctx context.Context, id string) (*models.User, error)
    UpdateUser(ctx context.Context, user *models.User) error
    DeleteUser(ctx context.Context, id string) error
}

type userRepository struct {
    db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
    return &userRepository{db: db}
}

func (r *userRepository) CreateUser(ctx context.Context, user *models.User) error {
    query := `INSERT INTO users (id, full_name, password_hash, email, created_at, updated_at, role)
              VALUES ($1, $2, $3, $4, $5, $6, $7)`
    _, err := r.db.ExecContext(ctx, query, user.ID, user.FullName, user.PasswordHash, user.Email, user.CreatedAt, user.UpdatedAt, user.Role)
    if err != nil {
        return err
    }
    return nil
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
    query := `SELECT id, full_name, password_hash, email, created_at, updated_at, role FROM users WHERE email = $1`
    row := r.db.QueryRowContext(ctx, query, email)

    var user models.User
    if err := row.Scan(&user.ID, &user.FullName, &user.PasswordHash, &user.Email, &user.CreatedAt, &user.UpdatedAt, &user.Role); err != nil {
        if err == sql.ErrNoRows {
            return nil, nil
        }
        return nil, err
    }
    return &user, nil
}

func (r *userRepository) GetUserByID(ctx context.Context, id string) (*models.User, error) {
    query := `SELECT id, full_name, password_hash, email, created_at, updated_at, role FROM users WHERE id = $1`
    row := r.db.QueryRowContext(ctx, query, id)

    var user models.User
    if err := row.Scan(&user.ID, &user.FullName, &user.PasswordHash, &user.Email, &user.CreatedAt, &user.UpdatedAt, &user.Role); err != nil {
        if err == sql.ErrNoRows {
            return nil, nil
        }
        return nil, err
    }
    return &user, nil
}

func (r *userRepository) UpdateUser(ctx context.Context, user *models.User) error {
    query := `UPDATE users SET full_name = $1, password_hash = $2, email = $3, updated_at = $4, role = $5 WHERE id = $6`
    _, err := r.db.ExecContext(ctx, query, user.FullName, user.PasswordHash, user.Email, user.UpdatedAt, user.Role, user.ID)
    if err != nil {
        return err
    }
    return nil
}

func (r *userRepository) DeleteUser(ctx context.Context, id string) error {
    query := `DELETE FROM users WHERE id = $1`
    _, err := r.db.ExecContext(ctx, query, id)
    if err != nil {
        return err
    }
    return nil
}

*/

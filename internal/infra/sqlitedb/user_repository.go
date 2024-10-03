package sqlitedb

import (
	"context"
	"database/sql"
	"time"

	"github.com/mecitsemerci/go-todo-app/internal/domain/auth"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) GetByID(ctx context.Context, userID auth.UserID) (*auth.User, error) {
	query := `SELECT id, email, password_hash, role, created_at, updated_at, last_login FROM users WHERE id=?`

	var user auth.User
	err := u.db.QueryRow(query, userID.String()).
		Scan(
			&user.ID,
			&user.Email,
			&user.Password,
			&user.Role,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.LastLogin,
		)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (u *UserRepository) FindByEmail(ctx context.Context, email auth.Email) (*auth.User, error) {
	query := `SELECT id, email, password_hash, role, created_at, updated_at, last_login FROM users WHERE email=?`

	var user auth.User
	err := u.db.QueryRow(query, email.String()).
		Scan(
			&user.ID,
			&user.Email,
			&user.Password,
			&user.Role,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.LastLogin,
		)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (u *UserRepository) Insert(ctx context.Context, user *auth.User) error {
	query := `INSERT INTO users (id, email, password_hash, role, created_at, updated_at, last_login) VALUES (?,?,?,?,?,?,?)`

	_, err := u.db.Exec(query,
		user.ID,
		user.Email,
		user.Password,
		user.Role,
		user.CreatedAt,
		user.UpdatedAt,
		user.LastLogin,
	)

	return err
}

func (u *UserRepository) Update(ctx context.Context, user *auth.User) error {
	query := `UPDATE users SET email=?, password_hash=?, role=?, updated_at=?, last_login=? WHERE id=?`

	_, err := u.db.Exec(query,
		user.Email,
		user.Password,
		user.Role,
		user.UpdatedAt,
		user.LastLogin,
		user.ID,
	)

	return err
}

func (u *UserRepository) Delete(ctx context.Context, userID auth.UserID) error {
	query := `DELETE FROM users WHERE id=?`

	_, err := u.db.Exec(query, userID.String())

	return err
}

func (u *UserRepository) UpdateLastLogin(ctx context.Context, userID auth.UserID) error {
	query := `UPDATE users SET last_login=? WHERE id=?`

	_, err := u.db.Exec(query,
		time.Now().UTC(),
		userID.String(),
	)

	return err
}

func (u *UserRepository) IsEmailTaken(ctx context.Context, email auth.Email) (bool, error) {
	query := `SELECT COUNT(*) FROM users WHERE email=?`

	var count int
	err := u.db.QueryRow(query, email.String()).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (u *UserRepository) UpdatePassword(ctx context.Context, userID auth.UserID, password string) error {
	query := `UPDATE users SET password_hash=? WHERE id=?`

	passwordHash, err := auth.NewPasswordHash(password)
	if err != nil {
		return err
	}

	_, err = u.db.Exec(query,
		passwordHash,
		userID.String(),
	)

	return err
}

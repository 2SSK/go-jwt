package repository

import (
	"context"

	"github.com/2SSK/jwt/internal/model/user"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type userRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(ctx context.Context, u *user.User) (*user.User, error) {
	query := `
		INSERT INTO users (first_name, last_name, password, email, phone, user_type)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, created_at, updated_at`

	err := r.db.QueryRow(ctx, query,
		u.FirstName, u.LastName, u.Password, u.Email, u.Phone, u.UserType,
	).Scan(&u.ID, &u.CreatedAt, &u.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return u, nil
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*user.User, error) {
	query := `
		SELECT id, first_name, last_name, password, email, phone, user_type, created_at, updated_at
		FROM users
		WHERE email = $1`

	u := &user.User{}
	err := r.db.QueryRow(ctx, query, email).Scan(
		&u.ID, &u.FirstName, &u.LastName, &u.Password, &u.Email, &u.Phone, &u.UserType, &u.CreatedAt, &u.UpdatedAt,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return u, nil
}

func (r *userRepository) GetUserByID(ctx context.Context, id uuid.UUID) (*user.User, error) {
	query := `
		SELECT id, first_name, last_name, password, email, phone, user_type, created_at, updated_at
		FROM users
		WHERE id = $1`

	u := &user.User{}
	err := r.db.QueryRow(ctx, query, id).Scan(
		&u.ID, &u.FirstName, &u.LastName, &u.Password, &u.Email, &u.Phone, &u.UserType, &u.CreatedAt, &u.UpdatedAt,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return u, nil
}

func (r *userRepository) GetUsers(ctx context.Context, limit, offset int) ([]*user.User, error) {
	query := `
		SELECT id, first_name, last_name, password, email, phone, user_type, created_at, updated_at
		FROM users
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2`

	rows, err := r.db.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*user.User
	for rows.Next() {
		u := &user.User{}
		err := rows.Scan(
			&u.ID, &u.FirstName, &u.LastName, &u.Password, &u.Email, &u.Phone, &u.UserType, &u.CreatedAt, &u.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}

func (r *userRepository) UpdateUser(ctx context.Context, u *user.User) error {
	query := `
		UPDATE users
		SET first_name = $1, last_name = $2, email = $3, phone = $4, user_type = $5, updated_at = NOW()
		WHERE id = $6`

	_, err := r.db.Exec(ctx, query,
		u.FirstName, u.LastName, u.Email, u.Phone, u.UserType, u.ID,
	)

	return err
}

func (r *userRepository) DeleteUser(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM users WHERE id = $1`

	_, err := r.db.Exec(ctx, query, id)

	return err
}

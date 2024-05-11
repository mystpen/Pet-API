package user

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	_ "github.com/lib/pq"

	"github.com/mystpen/Pet-API/internal/dto"
	"github.com/mystpen/Pet-API/internal/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) CreatUser(ctx context.Context, request *dto.RegistrationRequest, hashPassword []byte) error {
	query := `INSERT INTO users (id, username, email, password_hash)
	VALUES ($1, $2, $3, $4)
	RETURNING id`

	user := model.User{
		ID:       uuid.New(),
		UserName: request.UserName,
		Email:    request.Email,
		Password: hashPassword,
	}

	err := ur.db.QueryRowContext(ctx, query, user.ID, user.UserName, user.Email, user.Password).Scan(&user.ID)
	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"`:
			return model.ErrDuplicateEmail
		default:
			return err
		}
	}

	return nil
}

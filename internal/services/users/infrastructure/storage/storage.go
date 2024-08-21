package storage

import (
	"context"
	"database/sql"
	"time"

	"coffeeshop-api/internal/services/users/model"
	"coffeeshop-api/pkg/errors"
)

type storage struct {
	db *sql.DB
}

func New(db *sql.DB) *storage {
	return &storage{db: db}
}

func (s *storage) CreateUser(ctx context.Context, req *model.CreateUserReqStorage) (*model.CreateUserResStorage, error) {
	// Check if email is busy
	var emailBusy bool

	if err := s.db.QueryRowContext(ctx, `
		SELECT true
		FROM api.users
		WHERE email = $1
	`,
		req.Email,
	).Scan(&emailBusy); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, errors.Wrap(err, "check email")
	}

	if emailBusy {
		return nil, errors.Wrap(errors.EmailIsBusy, "busy email")
	}

	// Add user
	var userID uint64

	if err := s.db.QueryRowContext(ctx, `
		INSERT INTO api.users(
			username,
			email,
			phone,
			password_hash
		)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`,
		req.Username,
		req.Email,
		req.Phone,
		req.PasswordHash,
	).Scan(&userID); err != nil {
		return nil, errors.Wrap(err, "create user")
	}

	return &model.CreateUserResStorage{UserID: userID}, nil
}

func (s *storage) CheckCredentials(ctx context.Context, req *model.CheckCredentialsReqStorage) (*model.CheckCredentialsResStorage, error) {
	var userID uint64

	if err := s.db.QueryRowContext(ctx, `
		SELECT id
		FROM api.users
		WHERE
			email = $1 AND
			password_hash = $2 AND
			deleted_at IS NULL
	`,
		req.Email,
		req.PasswordHash,
	).Scan(&userID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.Wrap(errors.InvalidCredentials, "invalid credentials")
		}

		return nil, errors.Wrap(err, "check email")
	}

	return &model.CheckCredentialsResStorage{UserID: userID}, nil
}

func (s *storage) GetMe(ctx context.Context, req *model.GetMeReqStorage) (*model.GetMeResStorage, error) {
	var user model.User

	if err := s.db.QueryRowContext(ctx, `
		SELECT
			id,
			username,
			email,
			phone
		FROM api.users
		WHERE id = $1
	`,
		req.UserID,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Phone,
	); err != nil {
		return nil, errors.Wrap(err, "get user")
	}

	return &model.GetMeResStorage{User: &user}, nil
}

func (s *storage) ChangePassword(ctx context.Context, req *model.ChangePasswordReqStorage) error {
	var changed bool

	if err := s.db.QueryRowContext(ctx, `
		UPDATE api.users
		SET password_hash = $1
		WHERE
			id = $2 AND
			password_hash = $3
		RETURNING true
	`,
		req.NewPasswordHash,
		req.UserID,
		req.OldPasswordHash,
	).Scan(&changed); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.Wrap(errors.InvalidCredentials, "invalid password")
		}

		return errors.Wrap(err, "update password")
	}

	return nil
}

// DeleteMe provides user soft delete.
func (s *storage) DeleteMe(ctx context.Context, req *model.DeleteMeReqStorage) (*model.DeleteMeResStorage, error) {
	var (
		user      model.User
		deletedAt *time.Time
	)

	if err := s.db.QueryRowContext(ctx, `
		SELECT
			id,
			username,
			email,
			phone,
			deleted_at
		FROM api.users
		WHERE id = $1
	`,
		req.UserID,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Phone,
		&deletedAt,
	); err != nil {
		return nil, errors.Wrap(err, "select user")
	}

	if deletedAt != nil {
		return nil, errors.Wrapf(errors.UserAlreadyDeleted, "already deleted at %v", *deletedAt)
	}

	if _, err := s.db.ExecContext(ctx, `
		UPDATE api.users
		SET deleted_at = $1
		WHERE id = $2
	`,
		time.Now(),
		req.UserID,
	); err != nil {
		return nil, errors.Wrap(err, "delete user")
	}

	return &model.DeleteMeResStorage{User: &user}, nil
}

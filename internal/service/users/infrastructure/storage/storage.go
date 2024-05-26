package storage

import (
	"database/sql"
	"time"

	"coffeeshop-api/internal/service/users/model"
	"coffeeshop-api/pkg/errors"
)

type storage struct {
	db *sql.DB
}

func New(db *sql.DB) *storage {
	return &storage{db: db}
}

func (s *storage) Create(req *model.StorageCreateReq) (*model.StorageCreateRes, error) {
	var res model.StorageCreateRes

	// Check if email is busy
	var emailBusy bool

	if err := s.db.QueryRow(`
		SELECT true
		FROM api.users
		WHERE email = $1
	`, req.Email,
	).Scan(&emailBusy); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, errors.Wrap(err, "check email")
	}

	if emailBusy {
		return nil, errors.Wrap(errors.EmailIsBusy, "busy email")
	}

	// Add user
	if err := s.db.QueryRow(`
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
	).Scan(&res.UserID); err != nil {
		return nil, errors.Wrap(err, "create user")
	}

	return &res, nil
}

func (s *storage) CheckCredentials(req *model.StorageCheckCredentialsReq) (*model.StorageCheckCredentialsRes, error) {
	var res model.StorageCheckCredentialsRes

	if err := s.db.QueryRow(`
		SELECT id
		FROM api.users
		WHERE
			email = $1 AND
			password_hash = $2 AND
			deleted = FALSE
	`,
		req.Email,
		req.PasswordHash,
	).Scan(&res.UserID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.Wrap(errors.InvalidCredentials, "invalid credentials")
		}

		return nil, errors.Wrap(err, "check email")
	}

	return &res, nil
}

func (s *storage) IsDeleted(req *model.StorageIsDeletedReq) (*model.StorageIsDeletedRes, error) {
	var deletedAt *time.Time

	if err := s.db.QueryRow(`
		SELECT deleted_at
		FROM api.users
		WHERE id = $1
	`, req.UserID).Scan(&deletedAt); err != nil {
		return nil, errors.Wrap(err, "get user")
	}

	return &model.StorageIsDeletedRes{Deleted: deletedAt != nil}, nil
}

func (s *storage) GetMe(req *model.StorageGetMeReq) (*model.StorageGetMeRes, error) {
	var user model.User

	if err := s.db.QueryRow(`
		SELECT
			username,
			email,
			phone
		FROM api.users
		WHERE id = $1
	`, req.UserID).Scan(
		&user.Username,
		&user.Email,
		&user.Phone,
	); err != nil {
		return nil, errors.Wrap(err, "get user")
	}

	user.ID = req.UserID

	return &model.StorageGetMeRes{User: &user}, nil
}

func (s *storage) ChangePassword(req *model.StorageChangePasswordReq) error {
	if _, err := s.db.Exec(`
		UPDATE api.users
		SET password_hash = $1
		WHERE id = $2
	`,
		req.NewPasswordHash,
		req.UserID,
	); err != nil {
		return errors.Wrap(err, "update password")
	}

	return nil
}

// Delete provides soft delete.
func (s *storage) DeleteMe(req *model.StorageDeleteMeReq) (*model.StorageDeleteMeRes, error) {
	var user model.User

	if err := s.db.QueryRow(`
		UPDATE api.users
		SET deleted_at = $1
		WHERE id = $2
		RETURNING
			username,
			email,
			phone
	`,
		time.Now(),
		req.UserID,
	).Scan(
		user.Username,
		user.Email,
		user.Phone,
	); err != nil {
		return nil, errors.Wrap(err, "delete user")
	}

	user.ID = req.UserID

	return &model.StorageDeleteMeRes{User: &user}, nil
}

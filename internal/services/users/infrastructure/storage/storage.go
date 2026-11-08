package storage

import (
	"context"
	"database/sql"
	"time"

	"coffeeshop-api/internal/services/users/model"
	"coffeeshop-api/pkg/errors"
	"coffeeshop-api/pkg/tracer"

	sq "github.com/Masterminds/squirrel"
)

var builder = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

type storage struct {
	db *sql.DB
}

func New(db *sql.DB) *storage {
	return &storage{db: db}
}

func (s *storage) CreateUser(ctx context.Context, req *model.CreateUserReqStorage) (*model.CreateUserResStorage, error) {
	ctx, span := tracer.Start(ctx, "CreateUser")
	defer span.End()

	// Check if email is busy
	var id uint64

	query, args := builder.
		Select("id").
		From("api.users").
		Where(sq.Eq{"email": req.Email}).
		MustSql()

	if err := s.db.QueryRowContext(ctx, query, args...).Scan(&id); err != nil &&
		!errors.Is(err, sql.ErrNoRows) {
		return nil, errors.Wrap(err, "check email")
	}

	if id != 0 {
		return nil, errors.Wrap(errors.EmailIsBusy, "busy email")
	}

	// Add user
	var userID uint64

	query, args = builder.
		Insert("api.users").
		Columns(
			"username",
			"email",
			"phone",
			"password_hash",
		).
		Values(
			req.Username,
			req.Email,
			req.Phone,
			[]byte(req.PasswordHash),
		).
		Suffix("RETURNING id").
		MustSql()

	if err := s.db.QueryRowContext(ctx, query, args...).Scan(&userID); err != nil {
		return nil, errors.Wrap(err, "create user")
	}

	return &model.CreateUserResStorage{UserID: userID}, nil
}

func (s *storage) CheckCredentials(ctx context.Context, req *model.CheckCredentialsReqStorage) (*model.CheckCredentialsResStorage, error) {
	ctx, span := tracer.Start(ctx, "CheckCredentials")
	defer span.End()

	var userID uint64

	query, args := builder.
		Select(
			"id",
		).
		From("api.users").
		Where(sq.And{
			sq.Eq{
				"email":         req.Email,
				"password_hash": []byte(req.PasswordHash),
			},
			sq.Expr("deleted_at IS NULL"),
		}).
		MustSql()

	if err := s.db.QueryRowContext(ctx, query, args...).Scan(&userID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.Wrap(errors.InvalidCredentials, "invalid credentials")
		}

		return nil, errors.Wrap(err, "check email")
	}

	return &model.CheckCredentialsResStorage{UserID: userID}, nil
}

func (s *storage) GetMe(ctx context.Context, req *model.GetMeReqStorage) (*model.GetMeResStorage, error) {
	ctx, span := tracer.Start(ctx, "GetMe")
	defer span.End()

	var user model.User

	query, args := builder.
		Select(
			"id",
			"username",
			"email",
			"phone",
		).
		From("api.users").
		Where(sq.Eq{"id": req.UserID}).
		MustSql()

	if err := s.db.QueryRowContext(ctx, query, args...).Scan(
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
	ctx, span := tracer.Start(ctx, "ChangePassword")
	defer span.End()

	var changed bool

	query, args := builder.
		Update("api.users").
		Set("password_hash", req.NewPasswordHash).
		Where(sq.Eq{
			"id":            req.UserID,
			"password_hash": req.OldPasswordHash,
		}).
		Suffix("RETURNING true").
		MustSql()

	if err := s.db.QueryRowContext(ctx, query, args...).Scan(&changed); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.Wrap(errors.InvalidCredentials, "invalid password")
		}

		return errors.Wrap(err, "update password")
	}

	return nil
}

// DeleteMe provides user soft delete.
func (s *storage) DeleteMe(ctx context.Context, req *model.DeleteMeReqStorage) (*model.DeleteMeResStorage, error) {
	ctx, span := tracer.Start(ctx, "DeleteMe")
	defer span.End()

	var (
		user      model.User
		deletedAt *time.Time
	)

	query, args := builder.
		Select(
			"id",
			"username",
			"email",
			"phone",
			"deleted_at",
		).
		From("api.users").
		Where(sq.Eq{"id": req.UserID}).
		MustSql()

	if err := s.db.QueryRowContext(ctx, query, args...).Scan(
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

	query, args = builder.
		Update("api.users").
		Set("deleted_at", time.Now()).
		Where(sq.Eq{"id": req.UserID}).
		MustSql()

	if _, err := s.db.ExecContext(ctx, query, args...); err != nil {
		return nil, errors.Wrap(err, "delete user")
	}

	return &model.DeleteMeResStorage{User: &user}, nil
}

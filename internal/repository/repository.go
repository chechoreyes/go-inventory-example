package repository

import (
	"context"

	"github.com/chechoreyes/max-inventory/internal/entity"
	"github.com/jmoiron/sqlx"
)

// Repositoy is the interface that wraps the basic CRUD operations, is that comunicate with database
// Mockery is for unit tests

//go:generate mockery --name=Repository --output=repository --inpackage
type Repository interface {
	SaveUser(ctx context.Context, email, name, password string) error
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)

	SaveUserRole(ctx context.Context, userID, roleID int64) error
	RemoveUserRole(ctx context.Context, userID, roleID int64) error
}

type repo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Repository {
	return &repo{
		db: db,
	}
}

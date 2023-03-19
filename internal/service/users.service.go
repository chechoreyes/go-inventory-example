package service

import (
	"context"
	"errors"

	"github.com/chechoreyes/max-inventory/encryption"
	"github.com/chechoreyes/max-inventory/internal/models"
)

var (
	ErrUserAlreadyExists    = errors.New("user already exists")
	ErrorInvalidCredentials = errors.New("invalid credentials")
)

func (s *serv) RegisterUser(ctx context.Context, email, name, password string) error {
	u, _ := s.repo.GetUserByEmail(ctx, email)

	if u != nil {
		return ErrUserAlreadyExists
	}

	//HASH PASSWORD
	bb, err := encryption.Encrypt([]byte(password))
	if err != nil {
		return err
	}
	pass := encryption.ToBase64(bb)

	return s.repo.SaveUser(ctx, email, name, pass)
}

func (s *serv) LoginUser(ctx context.Context, email, password string) (*models.User, error) {

	u, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	//DECRYPT PASSWORD
	bb, err := encryption.FromBase64(u.Password)
	if err != nil {
		return nil, err
	}
	decrytedPassword, err := encryption.Decrypt(bb)
	if err != nil {
		return nil, err
	}

	if string(decrytedPassword) != password {
		return nil, ErrorInvalidCredentials
	}

	return &models.User{
		ID:    u.ID,
		Email: u.Email,
		Name:  u.Name,
	}, nil
}

func (s *serv) AddUserRole(ctx context.Context, userID, roleID int64) error {
	return s.repo.SaveUserRole(ctx, userID, roleID)
}

func (s *serv) RemoveUserRole(ctx context.Context, userID, roleID int64) error {
	return s.repo.RemoveUserRole(ctx, userID, roleID)
}

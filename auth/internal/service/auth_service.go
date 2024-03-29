package service

import (
	"context"

	"auth/internal/model"
	"auth/internal/store"
)

type AuthService struct {
	st store.Querier
}

func NewAuthService(st store.Querier) *AuthService {
	return &AuthService{
		st: st,
	}
}

type AuthServiceAPI interface {
	Login(ctx context.Context, in *model.LoginIn) (*model.LoginOut, error)
	Register(ctx context.Context, in *model.RegisterIn) (*model.RegisterOut, error)
}

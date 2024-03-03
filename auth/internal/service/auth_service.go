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
	Register(ctx context.Context, in *model.RegisterIn) (*model.RegisterOut, error)
}

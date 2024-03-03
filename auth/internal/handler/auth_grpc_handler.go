package handler

import (
	"context"

	"auth/internal/model"
	"auth/internal/service"
	"auth/pkg/rpc"
)

type AuthGRPCHandler struct {
	rpc.UnimplementedAuthServiceServer
	authSvc *service.AuthService
}

func NewAuthGRPCHandler(authSvc *service.AuthService) *AuthGRPCHandler {
	return &AuthGRPCHandler{authSvc: authSvc}
}

func (h *AuthGRPCHandler) Register(ctx context.Context, req *rpc.RegisterRequest) (*rpc.RegisterResponse, error) {
	in := &model.RegisterIn{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
		Role:     req.GetRole().String(),
	}

	resp, err := h.authSvc.Register(ctx, in)
	if err != nil {
		return nil, err
	}

	return &rpc.RegisterResponse{UserUid: resp.UID}, nil
}

func (h *AuthGRPCHandler) Login(ctx context.Context, req *rpc.LoginRequest) (*rpc.LoginResponse, error) {
	in := &model.LoginIn{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}

	resp, err := h.authSvc.Login(ctx, in)
	if err != nil {
		return nil, err
	}

	return &rpc.LoginResponse{Token: resp.Token}, nil
}

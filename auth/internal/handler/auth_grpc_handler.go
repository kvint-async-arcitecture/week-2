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
	}

	resp, err := h.authSvc.Register(ctx, in)
	if err != nil {
		return nil, err
	}

	return &rpc.RegisterResponse{UserUid: resp.UID}, nil
}

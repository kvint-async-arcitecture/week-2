package service

import (
	"gitlab.com/go-ledger/common/errors"
	"google.golang.org/grpc/codes"
)

var (
	ErrBadRequest           = errors.UserError("errBadRequest", errors.WithGRPCCode(codes.InvalidArgument))
	ErrAccountAlreadyExists = errors.UserError("errAccountAlreadyExists",
		errors.WithTemplate("account with email '%s' already exists"),
		errors.WithGRPCCode(codes.AlreadyExists))

	ErrPassword = errors.InternalError("errPassword", errors.WithGRPCCode(codes.Internal))
	ErrDB       = errors.InternalError("errDB", errors.WithGRPCCode(codes.Internal))
)

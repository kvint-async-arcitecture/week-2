package service

import (
	"context"
	"errors"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"

	"auth/internal/libs/password"
	"auth/internal/model"
	"auth/internal/store"
)

func (svc *AuthService) Register(ctx context.Context, in *model.RegisterIn) (*model.RegisterOut, error) {
	if err := in.Validate(); err != nil {
		return nil, ErrBadRequest.Wrap(err)
	}

	hashedPassword, err := password.Encrypt(in.Password)
	if err != nil {
		return nil, ErrPassword.Wrap(err)
	}

	userUID, err := svc.st.InsertUser(ctx, &store.InsertUserParams{
		Email:    in.Email,
		Password: hashedPassword,
	})
	if err != nil {
		var pgerr *pgconn.PgError
		if errors.As(err, &pgerr) {
			if pgerr.Code == pgerrcode.UniqueViolation {
				return nil, ErrAccountAlreadyExists.Wrap(err, in.Email)
			}
		}

		return nil, ErrDB.Wrap(err)
	}

	return &model.RegisterOut{UID: userUID}, nil
}

package service

import (
	"context"
	"encoding/base64"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"gitlab.com/go-ledger/common/log"
	"golang.org/x/crypto/ed25519"

	"auth/internal/config"
	"auth/internal/libs/password"
	"auth/internal/model"
)

func (svc *AuthService) Login(ctx context.Context, in *model.LoginIn) (*model.LoginOut, error) {
	if err := in.Validate(); err != nil {
		return nil, ErrBadRequest.Wrap(err)
	}

	user, err := svc.st.GetUser(ctx, in.Email)
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			log.Default().Info().Ctx(ctx).Msgf("user with email '%s' not found", in.Email)

			return nil, ErrPasswordMismatch.Wrap(err)
		default:
			return nil, ErrDB.Wrap(err)
		}
	}

	if err := password.CheckPasswordHash(in.Password, user.Password); err != nil {
		log.Default().Info().Ctx(ctx).Msgf("password mismath for user with email '%s'", in.Email)

		return nil, ErrPasswordMismatch.Wrap(err)
	}

	token, err := generateToken(in.Email, user.Role)
	if err != nil {
		return nil, ErrToken.Wrap(err)
	}

	return &model.LoginOut{Token: token}, nil
}

func generateToken(email, role string) (string, error) {
	createdAt := time.Now()

	tokenID, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}

	// https://habr.com/ru/articles/654191/

	claims := UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Audience:  []string{email},
			ID:        tokenID.String(),
			NotBefore: jwt.NewNumericDate(createdAt),
			ExpiresAt: jwt.NewNumericDate(createdAt.Add(config.Default().Credentials.ExpiresTimeout)),
			IssuedAt:  jwt.NewNumericDate(createdAt),
		},
		Email: email,
		Role:  role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodEdDSA, claims)

	privateKey, err := base64.StdEncoding.DecodeString(config.Default().Credentials.Token)
	if err != nil {
		return "", err
	}

	signedString, err := token.SignedString(ed25519.PrivateKey(privateKey))
	if err != nil {
		return "", err
	}

	return signedString, nil
}

type UserClaims struct {
	jwt.RegisteredClaims
	Email string `json:"email"`
	Role  string `json:"role"`
}

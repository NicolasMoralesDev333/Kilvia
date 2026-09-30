package service

import (
	"context"
	"errors"
	"net/mail"
	"strings"
	"time"

	"github.com/NicolasMoralesDev333/Kilvia/backend/internal/domain"
	"github.com/NicolasMoralesDev333/Kilvia/backend/internal/repository"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	UserID    string      `json:"uid"`
	CompanyID string      `json:"cid"`
	Role      domain.Role `json:"role"`
	jwt.RegisteredClaims
}

type AuthResult struct {
	Token string      `json:"token"`
	User  domain.User `json:"user"`
}

type Service struct {
	repository repository.Repository
	jwtKey     []byte
	tokenTTL   time.Duration
	now        func() time.Time
}

func New(repo repository.Repository, jwtSecret string, tokenTTL time.Duration) *Service {
	return &Service{repository: repo, jwtKey: []byte(jwtSecret), tokenTTL: tokenTTL, now: time.Now}
}

func (s *Service) Register(ctx context.Context, input domain.RegisterCompanyInput) (AuthResult, error) {
	if strings.TrimSpace(input.CompanyName) == "" || strings.TrimSpace(input.CUIT) == "" ||
		strings.TrimSpace(input.FirstName) == "" || strings.TrimSpace(input.LastName) == "" ||
		!validEmail(input.Email) || len(input.Password) < 10 {
		return AuthResult{}, ErrInvalidInput
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return AuthResult{}, err
	}
	user, err := s.repository.RegisterCompanyAdmin(ctx, input, string(hash))
	if err != nil {
		return AuthResult{}, translateRepositoryError(err)
	}
	token, err := s.issueToken(user)
	return AuthResult{Token: token, User: user}, err
}

func (s *Service) Login(ctx context.Context, input domain.LoginInput) (AuthResult, error) {
	user, err := s.repository.FindUserByEmail(ctx, input.Email)
	if err != nil {
		return AuthResult{}, ErrUnauthorized
	}
	if user.Status != "ACTIVO" || bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password)) != nil {
		return AuthResult{}, ErrUnauthorized
	}
	user.PasswordHash = ""
	token, err := s.issueToken(user)
	return AuthResult{Token: token, User: user}, err
}

func (s *Service) Me(ctx context.Context, identity domain.Identity) (domain.User, error) {
	user, err := s.repository.GetUser(ctx, identity.CompanyID, identity.UserID)
	return user, translateRepositoryError(err)
}

func (s *Service) ParseToken(raw string) (domain.Identity, error) {
	token, err := jwt.ParseWithClaims(raw, &Claims{}, func(token *jwt.Token) (any, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, ErrUnauthorized
		}
		return s.jwtKey, nil
	}, jwt.WithExpirationRequired(), jwt.WithIssuedAt())
	if err != nil || !token.Valid {
		return domain.Identity{}, ErrUnauthorized
	}
	claims, ok := token.Claims.(*Claims)
	if !ok || claims.UserID == "" || claims.CompanyID == "" {
		return domain.Identity{}, ErrUnauthorized
	}
	return domain.Identity{UserID: claims.UserID, CompanyID: claims.CompanyID, Role: claims.Role}, nil
}

func (s *Service) issueToken(user domain.User) (string, error) {
	now := s.now()
	claims := Claims{
		UserID: user.ID, CompanyID: user.CompanyID, Role: user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject: user.ID, IssuedAt: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(s.tokenTTL)),
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(s.jwtKey)
}

func validEmail(value string) bool {
	parsed, err := mail.ParseAddress(strings.TrimSpace(value))
	return err == nil && strings.EqualFold(parsed.Address, strings.TrimSpace(value))
}

func translateRepositoryError(err error) error {
	switch {
	case err == nil:
		return nil
	case errors.Is(err, repository.ErrNotFound):
		return ErrNotFound
	case errors.Is(err, repository.ErrConflict):
		return ErrConflict
	default:
		return err
	}
}

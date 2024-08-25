package service

import (
	"context"
	"crypto/sha1"
	"fmt"
	"kode-notes/internal/entity"
	"kode-notes/internal/repository"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthService struct {
	repo     repository.User
	SignKey  string
	TokenTTL time.Duration
	Salt     string
}

type AuthDependencies struct {
	userRepo repository.User
	signKey  string
	tokenTTL time.Duration
	salt     string
}

func NewAuthService(deps AuthDependencies) *AuthService {
	return &AuthService{
		repo:     deps.userRepo,
		SignKey:  deps.signKey,
		TokenTTL: deps.tokenTTL,
		Salt:     deps.salt,
	}
}

func (s *AuthService) Login(ctx context.Context, username, password string) (string, error) {
	if username == "" {
		return "", fmt.Errorf("username is required")
	}
	if password == "" {
		return "", fmt.Errorf("password is required")
	}
	user, err := s.repo.GetUser(ctx, username, s.generatePasswordHash(password))
	if err != nil {
		return "", fmt.Errorf("repository - GetUser - s.repo.GetUser: %w", err)
	}
	token, err := s.generateToken(user)
	if err != nil {
		return "", fmt.Errorf("service - Login - generateToken: %w", err)
	}
	return token, nil
}

func (s *AuthService) Register(ctx context.Context, username, password string) error {
	if username == "" {
		return fmt.Errorf("username is required")
	}
	if password == "" {
		return fmt.Errorf("password is required")
	}
	err := s.repo.CreateUser(ctx, username, s.generatePasswordHash(password))
	if err != nil {
		return fmt.Errorf("repository - CreateUser - s.repo.CreateUser: %w", err)
	}
	return nil
}

func (s *AuthService) generateToken(user entity.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"expires_at": time.Now().Add(s.TokenTTL).Unix(),
		"issued_at":  time.Now().Unix(),
		"user_id":    user.Id,
	})
	tokenString, err := token.SignedString([]byte(s.SignKey))
	if err != nil {
		return "", fmt.Errorf("service - Login - token.SignedString: %w", err)
	}
	return tokenString, nil
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.SignKey), nil
	})
	if err != nil {
		return 0, fmt.Errorf("service - ParseToken - jwt.Parse: %w", err)
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, fmt.Errorf("service - ParseToken - token.Claims: %w", err)
	}
	userId, ok := claims["user_id"].(float64)
	if !ok {
		return 0, fmt.Errorf("service - ParseToken - claims: %w", err)
	}
	return int(userId), nil
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(s.Salt)))
}

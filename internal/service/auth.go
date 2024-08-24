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

const (
	salt       = "hjqrhjqw124617ajfhajs"
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
	tokenTTL   = 12 * time.Hour
)

type AuthService struct {
	repo repository.User
}

func NewAuthService(userRepo repository.User) *AuthService {
	return &AuthService{
		repo: userRepo,
	}
}

func (s *AuthService) Login(ctx context.Context, username, password string) (string, error) {
	if username == "" {
		return "", fmt.Errorf("username is required")
	}
	if password == "" {
		return "", fmt.Errorf("password is required")
	}
	user, err := s.repo.GetUser(ctx, username, generatePasswordHash(password))
	if err != nil {
		return "", fmt.Errorf("repository - GetUser - s.repo.GetUser: %w", err)
	}
	token, err := generateToken(user)
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
	err := s.repo.CreateUser(ctx, username, generatePasswordHash(password))
	if err != nil {
		return fmt.Errorf("repository - CreateUser - s.repo.CreateUser: %w", err)
	}
	return nil
}

func generateToken(user entity.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"expires_at": time.Now().Add(tokenTTL).Unix(),
		"issued_at":  time.Now().Unix(),
		"user_id":    user.Id,
	})
	tokenString, err := token.SignedString([]byte(signingKey))
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
		return []byte(signingKey), nil
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

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

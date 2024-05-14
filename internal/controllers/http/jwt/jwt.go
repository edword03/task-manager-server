package jwt

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"task-manager/internal/config"
	"task-manager/internal/domain/entities"
	"time"
)

type tokenRepository interface {
	Get(string) (string, error)
	Set(key string, value string, expiration time.Duration) error
	Delete(string) error
}

type JWTService struct {
	cfg       *config.AppConfig
	tokenRepo tokenRepository
	maxAge    time.Duration
}

func NewJWTService(cfg *config.AppConfig, tokenRepo tokenRepository, maxAge time.Duration) *JWTService {
	return &JWTService{
		cfg:       cfg,
		tokenRepo: tokenRepo,
		maxAge:    maxAge,
	}
}

func (s *JWTService) GenerateAccessToken(user *entities.User) (string, error) {
	fmt.Println(time.Now().Add(s.cfg.AccessMaxAge).Unix())
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":        user.ID,
		"firstName": user.FirstName,
		"lastName":  user.LastName,
		"iat":       time.Now().Unix(),
		"exp":       time.Now().Add(s.cfg.AccessMaxAge).Unix(),
	})

	return accessToken.SignedString([]byte(s.cfg.SecretKey))
}

func (s *JWTService) ParseAccessToken(token string) (jwt.MapClaims, error) {
	var claims jwt.MapClaims
	parsedToken, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.cfg.SecretKey), nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %v", err)
	}

	if !parsedToken.Valid {
		return nil, errors.New(fmt.Sprintf("invalid token: %s", jwt.ErrTokenNotValidYet))
	}

	return claims, nil
}

func (s *JWTService) GenerateRefreshToken(id string) (string, error) {
	refreshToken := uuid.New().String()

	err := s.tokenRepo.Set(refreshToken, id, s.maxAge)
	if err != nil {
		return "", err
	}

	return refreshToken, nil
}

func (s *JWTService) CheckRefreshToken(refreshTokenString string) (string, error) {
	id, err := s.tokenRepo.Get(refreshTokenString)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (s *JWTService) DeleteRefreshToken(refreshTokenString string) error {
	err := s.tokenRepo.Delete(refreshTokenString)

	if err != nil {
		return err
	}

	return nil
}

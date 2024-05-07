package auth

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"task-manager/internal/domain/entities"
	"task-manager/internal/infrastructure/config"
	"task-manager/internal/infrastructure/database/redis/repositories"
	"time"
)

type IJWTService interface {
	GenerateAccessToken(user *entities.User) (string, error)
	ParseAccessToken(tokenString string) (*jwt.Token, error)
	GenerateRefreshToken(id string) (string, error)
	CheckRefreshToken(refreshTokenString string) (string, error)
	DeleteRefreshToken(refreshTokenString string) error
}

type JWTService struct {
	cfg       *config.AppConfig
	tokenRepo repositories.ITokenRepository
}

func NewJWTService(cfg *config.AppConfig, tokenRepo repositories.ITokenRepository) *JWTService {
	return &JWTService{
		cfg:       cfg,
		tokenRepo: tokenRepo,
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

func (s *JWTService) ParseAccessToken(token string) (*jwt.Token, error) {
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

	return parsedToken, nil
}

func (s *JWTService) GenerateRefreshToken(id string) (string, error) {
	refreshToken := uuid.New().String()

	err := s.tokenRepo.Set(refreshToken, id, maxAgeCookie*time.Duration(s.cfg.RefreshMaxAge))
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

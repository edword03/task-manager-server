package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"task-manager/internal/domain/entities"
	"task-manager/internal/infrastructure/config"
	"time"
)

type IJWTService interface {
	GenerateAccessToken(user *entities.User) (string, error)
	ParseAccessToken(tokenString string) (*jwt.Token, error)
	GenerateRefreshToken() string
}

type JWTService struct {
	cfg *config.AppConfig
}

func NewJWTService(cfg *config.AppConfig) *JWTService {
	return &JWTService{
		cfg: cfg,
	}
}

func (s *JWTService) GenerateAccessToken(user *entities.User) (string, error) {
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
	accessToken, _ := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.cfg.SecretKey), nil
	})

	return accessToken, nil
}

func (s *JWTService) GenerateRefreshToken() string {
	return uuid.New().String()
}

package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
	"todoApp/Entity"
	"todoApp/pkg/repository"
)

const (
	salt       = "jfirjosminvirp9ij48j29h7hf4w8jd98w1"
	signingKey = "kijdi#NJdeUwqF2401IJDInrfrfr5ij"
	tokenTTL   = 24 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int    `json:"user_id"`
	Name   string `json:"name"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user Entity.User) (int, error) {
	user.Password = s.generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(login Entity.Login) (string, error) {
	login.Password = s.generatePasswordHash(login.Password)
	user, err := s.repo.GetUser(login)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix()},
		UserId: user.Id,
		Name:   user.Name,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/markovk1n/spoty/internal/repository"
	"github.com/markovk1n/spoty/models"
)

// variables
const (
	salt       = "denj2n13842f197dfwodys1hg"
	signingKey = "uhfe134$E@#hfrew3@UE(*jrg"
	tokenTTL   = 12 * time.Hour
)

var ErrAuth = errors.New("auth error")

// structs
type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repository repository.Authorization
}

// funcs
func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repository: repo}
}

// Methods
func (s *AuthService) CreateUser(user models.User) error {
	// _, err := s.repository.GetUserByLogin(user.Username)
	// if err != nil {
	// 	if !errors.Is(err, sql.ErrNoRows) {
	// 		return fmt.Errorf("service: create user: %w", err)
	// 	}
	// } else {
	// 	return fmt.Errorf("serivce: create user: %w: username is taken", ErrAuth)
	// }

	if err := validUser(user); err != nil {
		return fmt.Errorf("service: create user: %w", err)
	}
	user.Password = generateHashPassword(user.Password)

	return s.repository.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repository.GetUser(username, generateHashPassword(password))
	if err != nil {

		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})
	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}
	return claims.UserId, nil
}

// // tools
func generateHashPassword(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func validUser(user models.User) error {
	for _, char := range user.Username {
		if char <= 32 || char >= 127 {
			return fmt.Errorf("%w: invalid username", ErrAuth)
		}
	}

	if len(user.Username) <= 3 || len(user.Username) >= 36 {
		return fmt.Errorf("%w: invalid username", ErrAuth)
	}
	// validPassword, err := regexp.MatchString(`(?=.*\d)(?=.*[a-z])(?=.*[A-Z]).{8,254}`, user.Password)
	// if err != nil {
	// 	return err
	// }
	// if !validPassword {
	// 	return fmt.Errorf("%w: invalid password", ErrAuth)
	// }
	// if user.Password != user.VerifyPassword {
	// 	return fmt.Errorf("%w: password dont match", ErrAuth)
	// }
	return nil
}

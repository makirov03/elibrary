package usecase

import (
	"context"
	"elibrary/domain/models"
	"elibrary/internal/repository"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

var jwtSecret = []byte("Directorsad@12")

type UserUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) *UserUsecase {
	return &UserUsecase{repo: repo}
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func generateToken(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"username": user.Username,
		"role":     user.Role,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func (u *UserUsecase) RegisterUser(ctx context.Context, user models.User) error {
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	err = u.repo.CreateUser(ctx, user)
	if err != nil {
		log.Println("Error registering user:", err)
		return err
	}

	return nil
}

func (u *UserUsecase) Login(ctx context.Context, username, password string) (string, error) {
	user, err := u.repo.GetUserByUsername(ctx, username)
	if err != nil {
		return "", errors.New("user not found")
	}

	if !checkPasswordHash(password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	token, err := generateToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}

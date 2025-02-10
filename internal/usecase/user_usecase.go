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

var jwtSecret = []byte("your-secret-key") // Change this to a secure key

type UserUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) *UserUsecase {
	return &UserUsecase{repo: repo}
}

// Hash password
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// Compare password
func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Generate JWT token
func generateToken(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"username": user.Username,
		"role":     user.Role,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func (u *UserUsecase) RegisterUser(ctx context.Context, user models.User) error {
	// Hash the password
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return err // ❌ If password hashing fails, we get an error
	}
	user.Password = hashedPassword

	// Store user
	err = u.repo.CreateUser(ctx, user)
	if err != nil {
		log.Println("Error registering user:", err) // ✅ Debug log
		return err
	}

	return nil
}

// Authenticate user and return JWT
func (u *UserUsecase) Login(ctx context.Context, username, password string) (string, error) {
	user, err := u.repo.GetUserByUsername(ctx, username)
	if err != nil {
		return "", errors.New("user not found")
	}

	// Verify password
	if !checkPasswordHash(password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	// Generate JWT
	token, err := generateToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}

package service

import (
	"cryptoserver/domain"
	"cryptoserver/errors"
	"cryptoserver/internal/auth"
	"cryptoserver/internal/repository"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (s *AuthService) Register(username, password string) (string, errors.CustomError) {
	const op = "register user"

	if username == "" {
		return "", errors.NewErrUserNameRequired("username is required field", http.StatusBadRequest, op)
	}
	if password == "" {
		return "", errors.NewErrPasswordRequired("password is required field", http.StatusBadRequest, op)
	}

	_, err := s.userRepo.GetByUsername(username)
	if err == nil {
		return "", errors.NewErrUserAlreadyExists("user with this username already exists", http.StatusConflict, op)
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.NewErrHashingPassword("fail to hash password", http.StatusInternalServerError, op)
	}

	user := domain.User{Username: username, PasswordHash: string(hash)}
	if err := s.userRepo.Create(&user); err != nil {
		return "", errors.NewErrCreateUser("fail to create user", http.StatusInternalServerError, op)
	}

	token, err := auth.GenerateToken(&user)
	if err != nil {
		return "", errors.NewErrGenerateToken("fail to generate JWT token", http.StatusInternalServerError, op)
	}

	return token, nil
}

func (s *AuthService) Login(username, password string) (string, errors.CustomError) {
	const op = "login user"

	if username == "" {
		return "", errors.NewErrUserNameRequired("username is required field", http.StatusBadRequest, op)
	}
	if password == "" {
		return "", errors.NewErrPasswordRequired("password is required field", http.StatusBadRequest, op)
	}

	user, err := s.userRepo.GetByUsername(username)
	if err != nil {
		return "", errors.NewErrUserNotFound("user not found", http.StatusNotFound, op)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return "", errors.NewErrInvalidCredentials("invalid credentials", http.StatusUnauthorized, op)
	}

	token, err := auth.GenerateToken(user)
	if err != nil {
		return "", errors.NewErrGenerateToken("failed to generate token", http.StatusInternalServerError, op)
	}

	return token, nil
}

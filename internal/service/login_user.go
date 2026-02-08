package service

import (
	"cryptoserver/internal/domain"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)


func (s *AuthService) Login(username, password string) (string, domain.CustomError) {
	const op = "login user"

	if username == "" {
		return "", domain.NewErrUserNameRequired("username is required field", http.StatusBadRequest, op)
	}
	if password == "" {
		return "", domain.NewErrPasswordRequired("password is required field", http.StatusBadRequest, op)
	}

	user, err := s.userRepo.GetByUsername(username)
	if err != nil {
		return "", domain.NewErrUserNotFound("user not found", http.StatusNotFound, op)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return "", domain.NewErrInvalidCredentials("invalid credentials", http.StatusUnauthorized, op)
	}

	token, err := GenerateToken(user)
	if err != nil {
		return "", domain.NewErrGenerateToken("failed to generate token", http.StatusInternalServerError, op)
	}

	return token, nil
}

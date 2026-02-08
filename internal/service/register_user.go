package service

import (
	"cryptoserver/internal/domain"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	
)




func (s *AuthService) RegisterUser(username, password string) (string, domain.CustomError) {
	const op = "register user"

	if username == "" {
		return "", domain.NewErrUserNameRequired("username is required field", http.StatusBadRequest, op)
	}
	if password == "" {
		return "", domain.NewErrPasswordRequired("password is required field", http.StatusBadRequest, op)
	}

	_, err := s.userRepo.GetByUsername(username)
	if err == nil {
		return "", domain.NewErrUserAlreadyExists("user with this username already exists", http.StatusConflict, op)
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", domain.NewErrHashingPassword("fail to hash password", http.StatusInternalServerError, op)
	}

	user := domain.User{Username: username, PasswordHash: string(hash)}
	if err := s.userRepo.Create(&user); err != nil {
		return "", domain.NewErrCreateUser("fail to create user", http.StatusInternalServerError, op)
	}

	token, err := GenerateToken(&user)
	if err != nil {
		return "", domain.NewErrGenerateToken("fail to generate JWT token", http.StatusInternalServerError, op)
	}

	return token, nil
}
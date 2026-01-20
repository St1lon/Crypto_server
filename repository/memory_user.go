package repository

import (
	"cryptoserver/domain"
	"errors"
)

type MemoryUserRepository struct{
	user map[string]*domain.User
}

func NewMemoryUserRepository() *MemoryUserRepository{
	return &MemoryUserRepository{
		user : make(map[string]*domain.User),
	}
}

func (r *MemoryUserRepository) Create(user *domain.User) error{
	_, exist := r.user[user.Username]
	if exist{
		return errors.New("user already exist")
	}
	r.user[user.Username] = user
	return nil
}
func (r *MemoryUserRepository) GetByUsername(username string) (*domain.User, error){
	user, exist := r.user[username]
	if !exist{
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (r *MemoryUserRepository) Update
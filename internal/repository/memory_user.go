package repository

import (
	"cryptoserver/domain"
	"errors"
	"sync"
)

type MemoryUserRepository struct{
	user map[string]*domain.User
	mu sync.RWMutex
}


func NewMemoryUserRepository() *MemoryUserRepository{
	return &MemoryUserRepository{
		user : make(map[string]*domain.User),
	}
}

func (r *MemoryUserRepository) Create(user *domain.User) error{
	r.mu.Lock()
	defer r.mu.Unlock()
	_, exist := r.user[user.Username]
	if exist{
		return errors.New("user already exist")
	}
	r.user[user.Username] = user
	return nil
}
func (r *MemoryUserRepository) GetByUsername(username string) (*domain.User, error){
	r.mu.RLock()
	defer r.mu.RUnlock()
	user, exist := r.user[username]
	if !exist{
		return nil, errors.New("user not found")
	}
	return user, nil
}



// func (r *MemoryUserRepository) Update(user *domain.User) error{
// 	r.mu.Lock()
// 	defer r.mu.Unlock()
// 	user, exist := r.user[user.Username]
// 	if !exist{
// 		return errors.New("user not found")
// 	}

// }
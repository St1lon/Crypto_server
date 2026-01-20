package repository

import (
	"cryptoserver/domain"
	"sync"
	"errors"
)

type MemoryCryptoRepository struct{
	cryptoList map[string]*domain.Crypto
	mu sync.RWMutex
}

func NewMemoryCryptoRepository() *MemoryCryptoRepository {
    return &MemoryCryptoRepository{
        cryptoList: make(map[string]*domain.Crypto),
    }
}

func (r *MemoryCryptoRepository) GetAll() ([]*domain.Crypto, error) {
    r.mu.RLock()
    defer r.mu.RUnlock()
	if len(r.cryptoList) == 0 {
		return nil, errors.New("no cryptocurrencies found")	
	}
    result := make([]*domain.Crypto, 0, len(r.cryptoList))
    for _, crypto := range r.cryptoList {
        result = append(result, crypto)
    }
    return result, nil
}

func (r* MemoryCryptoRepository)GetBySymbol(symbol string) (*domain.Crypto, error){
    r.mu.RLock()
    defer r.mu.RUnlock()
	crypto, exist := r.cryptoList[symbol]
	if !exist{
		return nil, errors.New("crypto not found")
	}
	return crypto, nil
}

func (r* MemoryCryptoRepository)Create(crypto *domain.Crypto) error{
    r.mu.Lock()
    defer r.mu.Unlock()
	_, exist := r.cryptoList[crypto.Symbol]
	if exist{
		return errors.New("crypto already exist")
	}
	r.cryptoList[crypto.Symbol] = crypto
	return nil
}

func (r* MemoryCryptoRepository)Update(crypto *domain.Crypto) error{
    r.mu.Lock()
    defer r.mu.Unlock()
	_, exist := r.cryptoList[crypto.Symbol]
	if !exist{
		return errors.New("crypto not found")
	}
	r.cryptoList[crypto.Symbol] = crypto
	return nil
}

    Delete(symbol string) error



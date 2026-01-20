package repository

import "cryptoserver/domain"

type UserRepository interface{
	GetByUsername(username string) (*domain.User, error)
	Create(user *domain.User) error
}

type CryptoRepository interface {
    GetAll() ([]*domain.Crypto, error)
    GetBySymbol(symbol string) (*domain.Crypto, error)
    Create(crypto *domain.Crypto) error
    Update(crypto *domain.Crypto) error
    Delete(symbol string) error
    AddPriceHistory(symbol string, record domain.PriceRecord) error
    GetPriceHistory(symbol string) ([]domain.PriceRecord, error)
}
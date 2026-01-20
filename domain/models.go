package domain

import "time"

// User - модель пользователя
type User struct {
    ID           string
    Username     string
    PasswordHash string
}

// Crypto - модель криптовалюты
type Crypto struct {
    Symbol       string    `json:"symbol"`
    CoinGeckoID  string    `json:"-"`
    Name         string    `json:"name"`
    CurrentPrice float64   `json:"current_price"`
    LastUpdated  time.Time `json:"last_updated"`
}

// PriceRecord - запись истории цен
type PriceRecord struct {
    Price     float64   `json:"price"`
    Timestamp time.Time `json:"timestamp"`
}
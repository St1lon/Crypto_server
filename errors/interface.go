package errors

// CustomError - интерфейс для всех кастомных ошибок
type CustomError interface {
	error
	GetCode() int
	GetMsg() string
	GetOp() string
}

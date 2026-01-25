package errors

import (
	"fmt"
)

type ErrUserNameRequired struct {
	Msg  string
	Code int
	Op   string
}

func (e ErrUserNameRequired) Error() string {
	return fmt.Sprintf("username required: %s, fail with code %d, operation: %s", e.Msg, e.Code, e.Op)
}

func (e ErrUserNameRequired) GetCode() int   { return e.Code }
func (e ErrUserNameRequired) GetMsg() string { return e.Msg }
func (e ErrUserNameRequired) GetOp() string  { return e.Op }

func NewErrUserNameRequired(msg string, code int, op string) CustomError {
	return ErrUserNameRequired{Msg: msg, Code: code, Op: op}
}

type ErrPasswordRequired struct {
	Msg  string
	Code int
	Op   string
}

func (e ErrPasswordRequired) Error() string {
	return fmt.Sprintf("password required: %s, fail with code %d, operation: %s", e.Msg, e.Code, e.Op)
}

func (e ErrPasswordRequired) GetCode() int   { return e.Code }
func (e ErrPasswordRequired) GetMsg() string { return e.Msg }
func (e ErrPasswordRequired) GetOp() string  { return e.Op }

func NewErrPasswordRequired(msg string, code int, op string) CustomError {
	return ErrPasswordRequired{Msg: msg, Code: code, Op: op}
}

type ErrUserAlreadyExists struct {
	Msg  string
	Code int
	Op   string
}

func (e ErrUserAlreadyExists) Error() string {
	return fmt.Sprintf("user already exists: %s, fail with code %d, operation: %s", e.Msg, e.Code, e.Op)
}

func (e ErrUserAlreadyExists) GetCode() int   { return e.Code }
func (e ErrUserAlreadyExists) GetMsg() string { return e.Msg }
func (e ErrUserAlreadyExists) GetOp() string  { return e.Op }

func NewErrUserAlreadyExists(msg string, code int, op string) CustomError {
	return ErrUserAlreadyExists{Msg: msg, Code: code, Op: op}
}


type ErrUserNotFound struct {
	Msg  string
	Code int
	Op   string
}
func (e ErrUserNotFound) Error() string {
	return fmt.Sprintf("user not found: %s, fail with code %d, operation: %s", e.Msg, e.Code, e.Op)
}
func (e ErrUserNotFound) GetCode() int   { return e.Code }
func (e ErrUserNotFound) GetMsg() string { return e.Msg }
func (e ErrUserNotFound) GetOp() string  { return e.Op }

func NewErrUserNotFound(msg string, code int, op string) CustomError {
	return ErrUserNotFound{Msg: msg, Code: code, Op: op}
}

type ErrInvalidCredentials struct {
	Msg  string
	Code int
	Op   string
}
func (e ErrInvalidCredentials) Error() string {
	return fmt.Sprintf("invalid credentials: %s, fail with code %d, operation: %s", e.Msg, e.Code, e.Op)
}
func (e ErrInvalidCredentials) GetCode() int   { return e.Code }
func (e ErrInvalidCredentials) GetMsg() string { return e.Msg }
func (e ErrInvalidCredentials) GetOp() string  { return e.Op }

func NewErrInvalidCredentials(msg string, code int, op string) CustomError {
	return ErrInvalidCredentials{Msg: msg, Code: code, Op: op}
}

type ErrTokenMissed struct {
	Msg  string
	Code int
	Op   string
}
func (e ErrTokenMissed) Error() string {
	return fmt.Sprintf("token missed: %s, fail with code %d, operation: %s", e.Msg, e.Code, e.Op)
}
func (e ErrTokenMissed) GetCode() int   { return e.Code }
func (e ErrTokenMissed) GetMsg() string { return e.Msg }
func (e ErrTokenMissed) GetOp() string  { return e.Op }

func NewErrTokenMissed(msg string, code int, op string) CustomError {
	return ErrTokenMissed{Msg: msg, Code: code, Op: op}
}

type ErrInvalidToken struct {
	Msg  string
	Code int
	Op   string
}

func (e ErrInvalidToken) Error() string {
	return fmt.Sprintf("invalid token: %s, fail with code %d, operation: %s", e.Msg, e.Code, e.Op)
}
func (e ErrInvalidToken) GetCode() int   { return e.Code }
func (e ErrInvalidToken) GetMsg() string { return e.Msg }
func (e ErrInvalidToken) GetOp() string  { return e.Op }

func NewErrInvalidToken(msg string, code int, op string) CustomError {
	return ErrInvalidToken{Msg: msg, Code: code, Op: op}
}
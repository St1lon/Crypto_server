package domain

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

type ErrWrongMethod struct {
	Msg    string
	Code   int
	Op     string
}

func (e ErrWrongMethod) Error() string {
	return fmt.Sprintf("wrong method: %s, fail with code %d, operation: %s", e.Msg, e.Code, e.Op)
}

func (e ErrWrongMethod) GetCode() int   { return e.Code }
func (e ErrWrongMethod) GetMsg() string { return e.Msg }
func (e ErrWrongMethod) GetOp() string  { return e.Op }

func NewErrWrongMethod(method string, code int, op string) CustomError {
	msg := fmt.Sprintf("wrong method: %s", method)
	return ErrWrongMethod{Code: code, Op: op, Msg: msg}
}

type ErrWrongCT struct {
	Msg  string
	Code int
	Op   string
}

func (e ErrWrongCT) Error() string {
	return fmt.Sprintf("wrong content type: %s, fail with code %d, operation: %s", e.Msg, e.Code, e.Op)
}

func (e ErrWrongCT) GetCode() int   { return e.Code }
func (e ErrWrongCT) GetMsg() string { return e.Msg }
func (e ErrWrongCT) GetOp() string  { return e.Op }

func NewErrWrongCT(msg string, code int, op string) CustomError {
	return ErrWrongCT{Msg: msg, Code: code, Op: op}
}

type ErrInvalidJSON struct {
	Msg  string
	Code int
	Op   string
}

func (e ErrInvalidJSON) Error() string {
	return fmt.Sprintf("invalid JSON: %s, fail with code %d, operation: %s", e.Msg, e.Code, e.Op)
}

func (e ErrInvalidJSON) GetCode() int   { return e.Code }
func (e ErrInvalidJSON) GetMsg() string { return e.Msg }
func (e ErrInvalidJSON) GetOp() string  { return e.Op }

func NewErrInvalidJSON(msg string, code int, op string) CustomError {
	return ErrInvalidJSON{Msg: msg, Code: code, Op: op}
}


type ErrHashingPassword struct {
	Msg  string
	Code int
	Op   string
}

func (e ErrHashingPassword) Error() string {
	return fmt.Sprintf("hashing password error: %s, fail with code %d, operation: %s", e.Msg, e.Code, e.Op)
}

func (e ErrHashingPassword) GetCode() int   { return e.Code }
func (e ErrHashingPassword) GetMsg() string { return e.Msg }
func (e ErrHashingPassword) GetOp() string  { return e.Op }

func NewErrHashingPassword(msg string, code int, op string) CustomError {
	return ErrHashingPassword{Msg: msg, Code: code, Op: op}
}

type ErrCreateUser struct {
	Msg  string
	Code int
	Op   string
}

func (e ErrCreateUser) Error() string {
	return fmt.Sprintf("create user error: %s, fail with code %d, operation: %s", e.Msg, e.Code, e.Op)
}

func (e ErrCreateUser) GetCode() int   { return e.Code }
func (e ErrCreateUser) GetMsg() string { return e.Msg }
func (e ErrCreateUser) GetOp() string  { return e.Op }

func NewErrCreateUser(msg string, code int, op string) CustomError {
	return ErrCreateUser{Msg: msg, Code: code, Op: op}
}

type ErrGenerateToken struct {
	Msg  string
	Code int
	Op   string
}

func (e ErrGenerateToken) Error() string {
	return fmt.Sprintf("generate token error: %s, fail with code %d, operation: %s", e.Msg, e.Code, e.Op)
}

func (e ErrGenerateToken) GetCode() int   { return e.Code }
func (e ErrGenerateToken) GetMsg() string { return e.Msg }
func (e ErrGenerateToken) GetOp() string  { return e.Op }

func NewErrGenerateToken(msg string, code int, op string) CustomError {
	return ErrGenerateToken{Msg: msg, Code: code, Op: op}
}


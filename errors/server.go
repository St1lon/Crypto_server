package errors

import (
	"fmt"
)

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

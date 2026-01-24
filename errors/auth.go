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

package errors

import (
	"fmt"
)

type ErrWrongMethod struct {
	Code   int
	Method string
	Op     string
}

func (e ErrWrongMethod) Error() string {
	return fmt.Sprintf("wrong method: %s, fail with code %d, operation: %s", e.Method, e.Code, e.Op)
}

func (e ErrWrongMethod) GetCode() int   { return e.Code }
func (e ErrWrongMethod) GetMsg() string { return e.Method }
func (e ErrWrongMethod) GetOp() string  { return e.Op }

func NewErrWrongMethod(method string, code int, op string) CustomError {
	return ErrWrongMethod{Method: method, Code: code, Op: op}
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

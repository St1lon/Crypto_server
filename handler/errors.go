package handler

import (
	"fmt"
	"net/http"
)

type ErrWrongMethod struct {
	Code int
	Method string
	Op   string
}

func (e ErrWrongMethod) Error() string {
	return fmt.Sprintf("wrong method: %s, fail with code %d, operation: %s", e.Method, e.Code, e.Op)
}

func NewErrWrongMethod(method string, code int, op string) error{
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
func NewErrWrongCT(msg string, code int, op string) error{
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
func NewErrInvalidJSON(msg string, code int, op string) error{
	return ErrInvalidJSON{Msg: msg, Code: code, Op: op}
}
package utils

import "errors"

var (
	ErrNotFound     = errors.New("registro no encontrado")
	ErrValidation   = errors.New("error de validacion")
	ErrDatabase     = errors.New("error de base de datos")
	ErrConflict     = errors.New("conflicto")
	ErrUnauthorized = errors.New("no autorizado")
	ErrInternal     = errors.New("error interno del servidor")
)

type AppError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func NewAppError(code, message string) *AppError {
	return &AppError{Code: code, Message: message}
}

func (e *AppError) Error() string {
	return e.Message
}

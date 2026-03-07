package errors

import "net/http"

func BadRequest(msg string, err error) *AppError {
	return New("BAD_REQUEST", msg, http.StatusBadRequest, err)
}

func NotFound(msg string, err error) *AppError {
	return New("NOT_FOUND", msg, http.StatusNotFound, err)
}

func Conflict(msg string, err error) *AppError {
	return New("CONFLICT", msg, http.StatusConflict, err)
}

func Internal(msg string, err error) *AppError {
	return New("INTERNAL_ERROR", msg, http.StatusInternalServerError, err)
}

func Unauthorized(msg string, err error) *AppError {
	return New("UNAUTHORIZED", msg, http.StatusUnauthorized, err)
}
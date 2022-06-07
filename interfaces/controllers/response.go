package controllers

import (
	"CoopeLunch-api/domain"
	"encoding/json"
	"net/http"
)

func response(w http.ResponseWriter, err error, body map[string]interface{}) error {
	status := getStatusCode(err)
	w.WriteHeader(status)
	data, _ := json.Marshal(body)
	w.Write(data)
	return err
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}
	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrForbidden:
		return http.StatusForbidden
	case domain.ErrUnauthorized:
		return http.StatusUnauthorized
	case domain.ErrBadRequest:
		return http.StatusBadRequest
	case domain.StatusCreated:
		return http.StatusCreated
	case domain.ErrUnknownType:
		return http.StatusUnsupportedMediaType
	case domain.ErrExistingEmail:
		return http.StatusConflict
	case domain.ErrIncorrectPassword:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}

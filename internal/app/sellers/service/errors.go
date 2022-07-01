package service

import (
	"errors"
	"net/http"
)

func errorHandler(err error) (int, error) {
	switch err.Error() {
	case "Error 1364: Field 'locality_id' doesn't have a default value":
		return http.StatusBadRequest, errors.New("invalid locality_id")
	default:
		return http.StatusInternalServerError, errors.New("request could not be processed")
	}
}

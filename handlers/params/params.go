package params

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
)

// GetUInt .
func GetUInt(r *http.Request, key string) (uint, error) {
	fID, err := strconv.Atoi(chi.URLParam(r, key))
	return uint(fID), err
}

// GetInt .
func GetInt(r *http.Request, key string) (int, error) {
	return strconv.Atoi(chi.URLParam(r, key))
}

// GetString .
func GetString(r *http.Request, key string) (string, error) {
	str := chi.URLParam(r, key)
	if str == "" {
		return "", errors.New("Missing")
	}

	return str, nil
}

// GetDate .
func GetDate(r *http.Request, key string) (time.Time, error) {
	str := chi.URLParam(r, key)

	return time.Parse("02-01-2006", str)
}

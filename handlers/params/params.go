package params

import (
	"net/http"
	"strconv"

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

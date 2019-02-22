package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/praveen001/go-boilerplate/app"
	"github.com/praveen001/go-boilerplate/models"
)

// UserHandler .
type UserHandler struct {
	user   *models.UserService
	logger *app.Logger
}

// NewUserHandler creates a new `UserHandler`
//
// It picks what it needs from application context and keeps it locally
//
// Easy to know what this handler group is using
func NewUserHandler(c *app.Context) *UserHandler {
	return &UserHandler{
		user:   c.DB.User,
		logger: c.Logger,
	}
}

// Create .
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var u *models.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		h.logger.Error("Unable to decode request body", err)
		return
	}

	if err := h.user.New(u); err != nil {
		h.logger.Error("Unable to create new user", err)
		return
	}

	json.NewEncoder(w).Encode(u)

}

// List .
func (h *UserHandler) List(w http.ResponseWriter, r *http.Request) {
	u, err := h.user.All()
	if err != nil {
		h.logger.Error("Unable to fetch users", err)
		return
	}

	json.NewEncoder(w).Encode(u)
}

// DeleteAll .
func (h *UserHandler) DeleteAll(w http.ResponseWriter, r *http.Request) {
	if err := h.user.DeleteAll(); err != nil {
		h.logger.Error("Unable to delete users", err)
		return
	}
	fmt.Fprintln(w, "Delete all users")
}

// Get ..
func (h *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	u := r.Context().Value("user")

	json.NewEncoder(w).Encode(u)
}

// Update ..
func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	var u models.User

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		h.logger.Error("Unable to decode request body", err)
		return
	}

	if err := h.user.Update(&u); err != nil {
		h.logger.Error("Unable to update user", err)
		return
	}

	json.NewEncoder(w).Encode(u)
}

// Delete ..
func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	u := r.Context().Value("user").(*models.User)

	if err := h.user.Delete(u.ID); err != nil {
		h.logger.Error("Unable to delete user", err)
		return
	}

	json.NewEncoder(w).Encode(u)
}

// Preload .
func (h *UserHandler) Preload(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rawUserID := chi.URLParam(r, "userID")

		userID, err := strconv.Atoi(rawUserID)
		if err != nil {
			h.logger.Error("Bad Request", err)
			return
		}

		u, err := h.user.Find(uint(userID))
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			h.logger.Error("Unable to find user", err)
			return
		}

		ctx := context.WithValue(r.Context(), "user", u)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

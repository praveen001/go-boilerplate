package handlers

import (
	"fmt"
	"net/http"

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
	fmt.Fprintln(w, "New User")
}

// List .
func (h *UserHandler) List(w http.ResponseWriter, r *http.Request) {
	u, _ := h.user.FindAll()
	h.logger.Info(u)
	fmt.Fprintln(w, "List all users")
}

// DeleteAll .
func (h *UserHandler) DeleteAll(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Delete all users")
}

// Get ..
func (h *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Get User By ID", chi.URLParam(r, "userID"))
}

// Update ..
func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Update User By ID", chi.URLParam(r, "userID"))
}

// Delete ..
func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Delete User By ID", chi.URLParam(r, "userID"))
}

// Preload .
func (h *UserHandler) Preload(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Loading users from middleware")
		next.ServeHTTP(w, r)
	})
}

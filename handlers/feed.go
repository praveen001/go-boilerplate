package handlers

import (
	"fmt"
	"net/http"

	"github.com/praveen001/go-boilerplate/app"

	"github.com/go-chi/chi"
)

// FeedHandler .
type FeedHandler struct {
	logger *app.Logger
}

// NewFeedHandler .
func NewFeedHandler(c *app.Context) *FeedHandler {
	return &FeedHandler{
		logger: c.Logger,
	}
}

// Create .
func (h *FeedHandler) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "New user feed", chi.URLParam(r, "feedID"))
}

// List .
func (h *FeedHandler) List(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Get User Feeds", chi.URLParam(r, "feedID"))
}

// DeleteAll .
func (h *FeedHandler) DeleteAll(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Delete all feeds")
}

// Get feed by id
func (h *FeedHandler) Get(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Get feed by id", chi.URLParam(r, "feedID"))
}

// Update .
func (h *FeedHandler) Update(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Update feed by id", chi.URLParam(r, "feedID"))
}

// Delete .
func (h *FeedHandler) Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w)
}

// Preload .
func (h *FeedHandler) Preload(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Loading feeds from middleware")
		next.ServeHTTP(w, r)
	})
}

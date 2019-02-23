package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/praveen001/go-boilerplate/app"
	"github.com/praveen001/go-boilerplate/models"
	"github.com/praveen001/go-boilerplate/repository"

	"github.com/go-chi/chi"
)

// FeedHandler .
type FeedHandler struct {
	feed   *repository.FeedRepository
	user   *repository.UserRepository
	logger *app.Logger
}

// NewFeedHandler creates a new `FeedHandler`
//
// It picks what it needs from application context and keeps it locally
//
// Easy to know what this handler group is using
func NewFeedHandler(c *app.Context) *FeedHandler {
	return &FeedHandler{
		feed:   c.DB.Feed,
		user:   c.DB.User,
		logger: c.Logger,
	}
}

// Create .
func (h *FeedHandler) Create(w http.ResponseWriter, r *http.Request) {
	userID := 1

	var f *models.Feed
	if err := json.NewDecoder(r.Body).Decode(&f); err != nil {
		h.logger.Error("Unable to decode request body", err)
		return
	}
	f.UserID = uint(userID)

	if err := h.feed.New(f); err != nil {
		h.logger.Error("Unable to create new feed", err)
		return
	}

	json.NewEncoder(w).Encode(f)
}

// List .
func (h *FeedHandler) List(w http.ResponseWriter, r *http.Request) {
	f, err := h.feed.All()
	if err != nil {
		h.logger.Error("Unable to fetch feeds", err)
		return
	}

	json.NewEncoder(w).Encode(f)
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

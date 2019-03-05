package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/praveen001/go-boilerplate/app"
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

// List .
func (h *FeedHandler) List(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(uint)

	user, err := h.user.Find(userID)
	if err != nil {
		h.logger.Error("Unable to fetch feeds", err)
		return
	}

	json.NewEncoder(w).Encode(user.Feeds)
}

// Get feed by ID
func (h *FeedHandler) Get(w http.ResponseWriter, r *http.Request) {

}

// Preload .
func (h *FeedHandler) Preload(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rawFeedID := chi.URLParam(r, "feedID")
		userID := r.Context().Value("userID").(int)

		feedID, err := strconv.Atoi(rawFeedID)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		feed, err := h.feed.Find(uint(feedID))
		if err != nil || !feed.BelongsTo(uint(userID)) {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		ctx := context.WithValue(r.Context(), "feed", feed)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

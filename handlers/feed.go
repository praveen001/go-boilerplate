package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/praveen001/go-boilerplate/repository"

	"github.com/praveen001/go-boilerplate/app"

	"github.com/praveen001/go-boilerplate/handlers/ctx"
	"github.com/praveen001/go-boilerplate/handlers/params"
)

// FeedHandler .
type FeedHandler struct {
	logger *app.Logger

	feed *repository.FeedRepository
}

// NewFeedHandler .
func NewFeedHandler(c *app.Context) *FeedHandler {
	return &FeedHandler{
		logger: c.Logger,

		feed: c.DB.Feed,
	}
}

// List .
func (h *FeedHandler) List(w http.ResponseWriter, r *http.Request) {
	user := ctx.GetUser(r.Context())

	feeds, err := h.feed.FindUserFeeds(user.ID)
	if err != nil {
		h.logger.Error("unable to fetch feeds", err.Error())
		return
	}

	json.NewEncoder(w).Encode(feeds)
}

// Get feed by ID
func (h *FeedHandler) Get(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(ctx.GetFeed(r.Context()))
}

// Preload .
func (h *FeedHandler) Preload(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := ctx.GetUser(r.Context())

		feedID, err := params.GetInt(r, "feedID")
		if err != nil {
			h.logger.Error("Invalid feed id", feedID, err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if feed, err := h.feed.FindUserFeed(user.ID, feedID); err != nil {
			h.logger.Error("Feed not found", feedID, err.Error())
			w.WriteHeader(http.StatusNotFound)
		} else {
			c := ctx.SetFeed(r.Context(), feed)
			next.ServeHTTP(w, r.WithContext(c))
		}

	})
}

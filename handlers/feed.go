package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/praveen001/go-boilerplate/app"
	"github.com/praveen001/go-boilerplate/handlers/ctx"
	"github.com/praveen001/go-boilerplate/handlers/params"
	"github.com/praveen001/go-boilerplate/repository"
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
	user := ctx.GetUser(r.Context())

	user, err := h.user.Find(user.ID)
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
		user := ctx.GetUser(r.Context())

		feedID, err := params.GetUInt(r, "feedID")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		feed, err := h.feed.Find(feedID)
		if err != nil || !feed.BelongsTo(user.ID) {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		c := ctx.SetFeed(r.Context(), feed)
		next.ServeHTTP(w, r.WithContext(c))
	})
}

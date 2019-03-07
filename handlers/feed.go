package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/praveen001/go-boilerplate/handlers/ctx"
	"github.com/praveen001/go-boilerplate/handlers/params"
	"github.com/praveen001/go-boilerplate/models"
)

// ListFeeds .
func (h *Handler) ListFeeds(w http.ResponseWriter, r *http.Request) {
	user := ctx.GetUser(r.Context())

	if err := models.FindUserFeeds(h.DB, user, &user.Feeds); err != nil {
		h.Logger.Error("unable to fetch feeds", err.Error())
		return
	}

	json.NewEncoder(w).Encode(user.Feeds)
}

// GetFeed by ID
func (h *Handler) GetFeed(w http.ResponseWriter, r *http.Request) {
	feed := ctx.GetFeed(r.Context())

	json.NewEncoder(w).Encode(feed)
}

// PreloadFeed .
func (h *Handler) PreloadFeed(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := ctx.GetUser(r.Context())

		feedID, err := params.GetInt(r, "feedID")
		if err != nil {
			h.Logger.Error("Invalid feed id", feedID, err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		feed := &models.Feed{ID: feedID}
		if err := models.FindUserFeed(h.DB, user, feed); err != nil {
			h.Logger.Error("Feed not found", feedID, err.Error())
			w.WriteHeader(http.StatusNotFound)
		} else {
			c := ctx.SetFeed(r.Context(), feed)
			next.ServeHTTP(w, r.WithContext(c))
		}

	})
}

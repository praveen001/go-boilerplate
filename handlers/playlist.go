package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/praveen001/go-boilerplate/handlers/ctx"
	"github.com/praveen001/go-boilerplate/handlers/params"

	"github.com/praveen001/go-boilerplate/repository"

	"github.com/praveen001/go-boilerplate/models"

	"github.com/praveen001/go-boilerplate/app"
)

// PlaylistHandler ..
type PlaylistHandler struct {
	playlist *repository.PlaylistRepository
	item     *repository.ItemRepository
	logger   *app.Logger
}

// NewPlaylistHandler .
func NewPlaylistHandler(c *app.Context) *PlaylistHandler {
	return &PlaylistHandler{
		playlist: c.DB.Playlist,
		item:     c.DB.Item,
		logger:   c.Logger,
	}
}

// Create .
func (h *PlaylistHandler) Create(w http.ResponseWriter, r *http.Request) {
	feed := ctx.GetFeed(r.Context())

	playlist := &models.Playlist{}
	if err := json.NewDecoder(r.Body).Decode(playlist); err != nil {
		h.logger.Error("Invalid playlist", err.Error())
		return
	}

	playlist.FeedID = feed.ID
	playlist.Status = models.PlaylistStatusNew

	if err := h.playlist.Save(playlist); err != nil {
		h.logger.Error("Unable to save playlist", err.Error())
		return
	}
}

// Get .
func (h *PlaylistHandler) Get(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(ctx.GetPlaylist(r.Context()))
}

// Update .
func (h *PlaylistHandler) Update(w http.ResponseWriter, r *http.Request) {
	playlist := ctx.GetPlaylist(r.Context())

	oldItems := playlist.Items
	playlist.Items = []*models.Item{}

	if err := json.NewDecoder(r.Body).Decode(playlist); err != nil {
		h.logger.Error("Invalid playlist", err.Error())
		return
	}
	h.playlist.Save(playlist)

	go h.item.DeleteMulti(oldItems)
}

// Delete .
func (h *PlaylistHandler) Delete(w http.ResponseWriter, r *http.Request) {

}

// GetByDate .
func (h *PlaylistHandler) GetByDate(w http.ResponseWriter, r *http.Request) {
	date, err := params.GetInt(r, "date")
	if err != nil {
		h.logger.Error("Invalid date", err.Error())
		return
	}

	p, err := h.playlist.FindByDate(uint64(date))
	if err != nil {
		h.logger.Error("Unable to find playlists", err.Error())
		return
	}

	json.NewEncoder(w).Encode(p)
}

// Preload .
func (h *PlaylistHandler) Preload(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		feedID, _ := params.GetUInt(r, "feedID")

		playlistID, err := params.GetUInt(r, "playlistID")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		playlist, err := h.playlist.Find(playlistID)
		if err != nil || playlist.FeedID != feedID {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		c := ctx.SetPlaylist(r.Context(), playlist)
		next.ServeHTTP(w, r.WithContext(c))
	})
}

package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

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
	feed := r.Context().Value("feed").(*models.Feed)

	playlist := &models.Playlist{}
	if err := json.NewDecoder(r.Body).Decode(playlist); err != nil {
		h.logger.Error("Invalid playlist", err.Error())
		return
	}

	playlist.FeedID = feed.ID
	playlist.Status = models.PlaylistStatusNew
	playlist.GenerateGroupID()

	if err := h.playlist.Save(playlist); err != nil {
		h.logger.Error("Unable to save playlist", err.Error())
		return
	}
}

// Get .
func (h *PlaylistHandler) Get(w http.ResponseWriter, r *http.Request) {
	playlist := r.Context().Value("playlist")
	json.NewEncoder(w).Encode(playlist)
}

// Update .
func (h *PlaylistHandler) Update(w http.ResponseWriter, r *http.Request) {
	playlist := r.Context().Value("playlist").(*models.Playlist)

	playlist.Items = []*models.Item{}
	groupID := playlist.GroupID

	if err := json.NewDecoder(r.Body).Decode(playlist); err != nil {
		h.logger.Error("Invalid playlist", err.Error())
		return
	}
	playlist.GenerateGroupID()
	h.playlist.Save(playlist)

	go h.item.DeleteByGroupID(groupID)
}

// Delete .
func (h *PlaylistHandler) Delete(w http.ResponseWriter, r *http.Request) {

}

// GetByDate .
func (h *PlaylistHandler) GetByDate(w http.ResponseWriter, r *http.Request) {
	rawDate := chi.URLParam(r, "date")
	date, err := strconv.Atoi(rawDate)
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
		rawPlaylistID := chi.URLParam(r, "playlistID")

		feedID, _ := strconv.Atoi(chi.URLParam(r, "feedID"))

		playlistID, err := strconv.Atoi(rawPlaylistID)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		playlist, err := h.playlist.Find(uint(playlistID))
		if err != nil || playlist.FeedID != uint(feedID) {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		ctx := context.WithValue(r.Context(), "playlist", playlist)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

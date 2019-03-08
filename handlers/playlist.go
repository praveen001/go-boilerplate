package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/praveen001/go-boilerplate/app"
	"github.com/praveen001/go-boilerplate/repository"

	"github.com/praveen001/go-boilerplate/handlers/ctx"
	"github.com/praveen001/go-boilerplate/handlers/params"
	"github.com/praveen001/go-boilerplate/models"
)

// PlaylistHandler .
type PlaylistHandler struct {
	logger *app.Logger

	playlist *repository.PlaylistRepository
}

// NewPlaylistHandler .
func NewPlaylistHandler(c *app.Context) *PlaylistHandler {
	return &PlaylistHandler{
		logger: c.Logger,

		playlist: c.DB.Playlist,
	}
}

// Create .
func (h *PlaylistHandler) Create(w http.ResponseWriter, r *http.Request) {
	feed := ctx.GetFeed(r.Context())

	playlist := &models.Playlist{
		FeedID: feed.ID,
		Status: models.PlaylistStatusNew,
	}
	if err := json.NewDecoder(r.Body).Decode(playlist); err != nil {
		h.logger.Error("Invalid playlist", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	date, err := params.GetDate(r, "date")
	if err != nil {
		h.logger.Error("Invalid date", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	playlist.PlayOn = date

	if err := h.playlist.Create(playlist); err != nil {
		h.logger.Error("Unable to create playlist", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// Read .
func (h *PlaylistHandler) Read(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(ctx.GetPlaylist(r.Context()))
}

// Update .
func (h *PlaylistHandler) Update(w http.ResponseWriter, r *http.Request) {
	// playlist := ctx.GetPlaylist(r.Context())

	// oldItems := playlist.Items
	// playlist.Items = []*models.Item{}

	// if err := json.NewDecoder(r.Body).Decode(playlist); err != nil {
	// 	h.logger.Error("Invalid playlist", err.Error())
	// 	return
	// }
	// h.playlist.Save(playlist)

	// go h.item.DeleteMulti(oldItems)
}

// Delete .
func (h *PlaylistHandler) Delete(w http.ResponseWriter, r *http.Request) {
	playlist := ctx.GetPlaylist(r.Context())

	if err := h.playlist.Delete(playlist); err != nil {
		h.logger.Error("Unable to delete playlist", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// ReadByDate .
func (h *PlaylistHandler) ReadByDate(w http.ResponseWriter, r *http.Request) {
	feed := ctx.GetFeed(r.Context())

	date, err := params.GetDate(r, "date")
	if err != nil {
		h.logger.Error("Invalid date", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	playlists, err := h.playlist.FindPlaylistByDate(feed.ID, date)
	if err != nil {
		h.logger.Error("Unable to find playlists for date", date, err.Error())
		return
	}

	json.NewEncoder(w).Encode(playlists)
}

// Preload .
func (h *PlaylistHandler) Preload(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		feed := ctx.GetFeed(r.Context())

		playlistID, err := params.GetInt(r, "playlistID")
		if err != nil {
			h.logger.Error("Invalid playlist id", playlistID, err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		playlist, err := h.playlist.Read(playlistID)
		if err != nil || playlist.FeedID != feed.ID {
			h.logger.Error("Unable to find playlist", playlistID, err.Error())
			w.WriteHeader(http.StatusNotFound)
			return
		}

		c := ctx.SetPlaylist(r.Context(), playlist)
		next.ServeHTTP(w, r.WithContext(c))
	})
}

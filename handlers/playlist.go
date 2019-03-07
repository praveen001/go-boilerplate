package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/praveen001/go-boilerplate/handlers/ctx"
	"github.com/praveen001/go-boilerplate/handlers/params"
	"github.com/praveen001/go-boilerplate/models"
)

// CreatePlaylist .
func (h *Handler) CreatePlaylist(w http.ResponseWriter, r *http.Request) {
	feed := ctx.GetFeed(r.Context())

	playlist := &models.Playlist{
		FeedID: feed.ID,
		Status: models.PlaylistStatusNew,
	}
	if err := json.NewDecoder(r.Body).Decode(playlist); err != nil {
		h.Logger.Error("Invalid playlist", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := playlist.Create(h.DB); err != nil {
		h.Logger.Error("Unable to create playlist", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// GetPlaylist .
func (h *Handler) GetPlaylist(w http.ResponseWriter, r *http.Request) {
	playlist := ctx.GetPlaylist(r.Context())

	json.NewEncoder(w).Encode(playlist)
}

// UpdatePlaylist .
func (h *Handler) UpdatePlaylist(w http.ResponseWriter, r *http.Request) {
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

// DeletePlaylist .
func (h *Handler) DeletePlaylist(w http.ResponseWriter, r *http.Request) {
	playlist := ctx.GetPlaylist(r.Context())

	if err := playlist.Delete(h.DB); err != nil {
		h.Logger.Error("Unable to delete playlist", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// GetPlaylistByDate .
func (h *Handler) GetPlaylistByDate(w http.ResponseWriter, r *http.Request) {
	feed := ctx.GetFeed(r.Context())

	date, err := params.GetInt(r, "date")
	if err != nil {
		h.Logger.Error("Invalid date", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	playlists, err := models.FindPlaylistByDate(h.DB, date, feed.ID)
	if err != nil {
		h.Logger.Error("Unable to find playlists for date", date, err.Error())
		return
	}

	json.NewEncoder(w).Encode(playlists)
}

// PreloadPlaylist .
func (h *Handler) PreloadPlaylist(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		feed := ctx.GetFeed(r.Context())

		playlistID, err := params.GetInt(r, "playlistID")
		if err != nil {
			h.Logger.Error("Invalid playlist id", playlistID, err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		playlist := &models.Playlist{ID: playlistID, FeedID: feed.ID}
		if err := playlist.Find(h.DB); err != nil {
			h.Logger.Error("Unable to find playlist", playlistID, err.Error())
			w.WriteHeader(http.StatusNotFound)
			return
		}

		c := ctx.SetPlaylist(r.Context(), playlist)
		next.ServeHTTP(w, r.WithContext(c))
	})
}

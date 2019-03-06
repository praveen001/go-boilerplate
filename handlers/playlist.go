package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/praveen001/go-boilerplate/repository"

	"github.com/praveen001/go-boilerplate/models"

	"github.com/praveen001/go-boilerplate/app"
)

// PlaylistHandler ..
type PlaylistHandler struct {
	playlist *repository.PlaylistRepository
	logger   *app.Logger
}

// NewPlaylistHandler .
func NewPlaylistHandler(c *app.Context) *PlaylistHandler {
	return &PlaylistHandler{
		playlist: c.DB.Playlist,
		logger:   c.Logger,
	}
}

// Create .
func (h *PlaylistHandler) Create(w http.ResponseWriter, r *http.Request) {
	feed := r.Context().Value("feed").(*models.Feed)

	// Arrange

	// Validate

	// Create
	playlist := &models.Playlist{
		FeedID: feed.ID,
		PlayOn: time.Now(),
		Status: models.PlaylistStatusNew,
		Type:   models.PlaylistTypeNormal,
		Items: []*models.Item{
			&models.Item{
				AssetID: "Test",
			},
		},
		GroupID: "123-123",
	}

	h.playlist.New(playlist)
}

// Get .
func (h *PlaylistHandler) Get(w http.ResponseWriter, r *http.Request) {

}

// Update .
func (h *PlaylistHandler) Update(w http.ResponseWriter, r *http.Request) {

}

// Delete .
func (h *PlaylistHandler) Delete(w http.ResponseWriter, r *http.Request) {

}

// GetByDate .
func (h *PlaylistHandler) GetByDate(w http.ResponseWriter, r *http.Request) {

}

// Preload .
func (h *PlaylistHandler) Preload(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Playlist preloading")
		next.ServeHTTP(w, r)
	})
}

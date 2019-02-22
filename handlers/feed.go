package handlers

import (
	"fmt"
	"net/http"

	"github.com/praveen001/go-boilerplate/app"

	"github.com/go-chi/chi"
)

// FeedHandler .
type FeedHandler struct {
	logger *app.Logger
}

// NewFeedHandler .
func NewFeedHandler(c *app.Context) *FeedHandler {
	return &FeedHandler{
		logger: c.Logger,
	}
}

// Get feed by id
func (fh *FeedHandler) Get(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Get feed by id", chi.URLParam(r, "feedID"))
}

// Update .
func (fh *FeedHandler) Update(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Update feed by id", chi.URLParam(r, "feedID"))
}

// Delete .
func (fh *FeedHandler) Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w)
}

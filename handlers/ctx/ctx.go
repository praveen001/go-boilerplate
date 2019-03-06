package ctx

import (
	"context"

	"github.com/praveen001/go-boilerplate/models"
)

type contextKey struct {
	name string
}

var (
	userCtxKey     = contextKey{"User"}
	feedCtxKey     = contextKey{"Feed"}
	playlistCtxKey = contextKey{"Playlist"}
)

// SetUser .
func SetUser(ctx context.Context, f *models.User) context.Context {
	return context.WithValue(ctx, userCtxKey, f)
}

// GetUser .
func GetUser(c context.Context) *models.User {
	if id := c.Value(userCtxKey); id != nil {
		return id.(*models.User)
	}

	return nil
}

// SetFeed .
func SetFeed(ctx context.Context, f *models.Feed) context.Context {
	return context.WithValue(ctx, feedCtxKey, f)
}

// GetFeed .
func GetFeed(ctx context.Context) *models.Feed {
	if f := ctx.Value(feedCtxKey); f != nil {
		return f.(*models.Feed)
	}

	return nil
}

// SetPlaylist .
func SetPlaylist(ctx context.Context, p *models.Playlist) context.Context {
	return context.WithValue(ctx, playlistCtxKey, p)
}

// GetPlaylist .
func GetPlaylist(ctx context.Context) *models.Playlist {
	if f := ctx.Value(playlistCtxKey); f != nil {
		return f.(*models.Playlist)
	}

	return nil
}

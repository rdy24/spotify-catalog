package tracks

import (
	"context"

	"github.com/rdy24/spotify-catalog/internal/models/trackactivities"
	"github.com/rdy24/spotify-catalog/internal/repository/spotify"
)

//go:generate mockgen -source=service.go -destination=service_mock_test.go -package=tracks
type SpotifyOutbound interface {
	Search(ctx context.Context, query string, limit, offset int) (*spotify.SpotifySearchResponse, error)
}

type trackActivitiesRepository interface {
	Create(ctx context.Context, model trackactivities.TrackActivity) error
	Update(ctx context.Context, model trackactivities.TrackActivity) error
	Get(ctx context.Context, userID uint, spotifyID string) (*trackactivities.TrackActivity, error)
	GetBulkSpotifyIDs(ctx context.Context, userID uint, spotifyIDs []string) (map[string]trackactivities.TrackActivity, error)
}

type service struct {
	spotifyOutbound     SpotifyOutbound
	trackActivitiesRepo trackActivitiesRepository
}

func NewService(spotifyOutbound SpotifyOutbound, trackActivitiesRepo trackActivitiesRepository) *service {
	return &service{
		spotifyOutbound:     spotifyOutbound,
		trackActivitiesRepo: trackActivitiesRepo,
	}
}

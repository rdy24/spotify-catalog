package tracks

import (
	"context"

	"github.com/rdy24/spotify-catalog/internal/models/spotify"
	"github.com/rdy24/spotify-catalog/internal/models/trackactivities"
	spotifyRepo "github.com/rdy24/spotify-catalog/internal/repository/spotify"
	"github.com/rs/zerolog/log"
)

func (s *service) GetRecommendation(ctx context.Context, userID uint, limit int, trackID string) (*spotify.RecommendationResponse, error) {
	trackDetails, err := s.spotifyOutbound.GetRecommendation(ctx, limit, trackID)
	if err != nil {
		log.Error().Err(err).Msg("error get recommendation from spotify outbound")
		return nil, err
	}

	trackIDs := make([]string, len(trackDetails.Tracks))
	for idx, item := range trackDetails.Tracks {
		trackIDs[idx] = item.ID
	}

	trackActivities, err := s.trackActivitiesRepo.GetBulkSpotifyIDs(ctx, userID, trackIDs)
	if err != nil {
		log.Error().Err(err).Msg("error get track activities from database")
		return nil, err
	}

	return modelToRecommendationResponse(trackDetails, trackActivities), nil
}

func modelToRecommendationResponse(data *spotifyRepo.SpotifyRecommendationResponse, mapTrackActivities map[string]trackactivities.TrackActivity) *spotify.RecommendationResponse {
	if data == nil {
		return nil
	}

	items := make([]spotify.SpotifyTrackObject, 0)

	for _, item := range data.Tracks {
		artistsName := make([]string, len(item.Artists))
		for idx, artist := range item.Artists {
			artistsName[idx] = artist.Name
		}

		imageUrls := make([]string, len(item.Album.Images))
		for idx, image := range item.Album.Images {
			imageUrls[idx] = image.URL
		}

		items = append(items, spotify.SpotifyTrackObject{
			// album related fields
			AlbumType:        item.Album.AlbumType,
			AlbumTotalTracks: item.Album.TotalTracks,
			AlbumImagesURL:   imageUrls,
			AlbumName:        item.Album.Name,
			// artist related fields
			ArtistsName: artistsName,
			// track related fields
			Explicit: item.Explicit,
			ID:       item.ID,
			Name:     item.Name,
			IsLiked:  mapTrackActivities[item.ID].IsLiked,
		})
	}

	return &spotify.RecommendationResponse{
		Items: items,
	}
}

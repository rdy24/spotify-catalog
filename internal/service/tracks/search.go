package tracks

import (
	"context"

	"github.com/rdy24/spotify-catalog/internal/models/spotify"
	"github.com/rdy24/spotify-catalog/internal/models/trackactivities"
	spotifyRepository "github.com/rdy24/spotify-catalog/internal/repository/spotify"
	"github.com/rs/zerolog/log"
)

func (s *service) Search(ctx context.Context, query string, pageSize, pageIndex int, userID uint) (*spotify.SearchResponse, error) {
	limit := pageSize
	offset := (pageIndex - 1) * pageSize

	trackDetails, err := s.spotifyOutbound.Search(ctx, query, limit, offset)

	if err != nil {
		log.Error().Err(err).Msg("failed to search tracks")
		return nil, err
	}

	trackIDs := make([]string, len(trackDetails.Tracks.Items))

	for idx, track := range trackDetails.Tracks.Items {
		trackIDs[idx] = track.ID
	}

	trackActivities, err := s.trackActivitiesRepo.GetBulkSpotifyIDs(ctx, userID, trackIDs)

	if err != nil {
		log.Error().Err(err).Msg("failed to get track activities")
		return nil, err
	}

	return modelToResponse(trackDetails, trackActivities), nil

}

func modelToResponse(data *spotifyRepository.SpotifySearchResponse, mapTrackActivities map[string]trackactivities.TrackActivity) *spotify.SearchResponse {
	if data == nil {
		return nil
	}

	items := make([]spotify.SpotifyTrackObject, 0)

	for _, item := range data.Tracks.Items {

		artistNames := make([]string, len(item.Artists))

		for idx, artist := range item.Artists {
			artistNames[idx] = artist.Name
		}

		imageUrls := make([]string, len(item.Album.Images))

		for idx, image := range item.Album.Images {
			imageUrls[idx] = image.URL
		}

		items = append(items, spotify.SpotifyTrackObject{
			AlbumType:        item.Album.AlbumType,
			AlbumTotalTracks: item.Album.TotalTracks,
			AlbumName:        item.Album.Name,
			AlbumImagesURL:   imageUrls,
			ArtistsName:      artistNames,
			Explicit:         item.Explicit,
			ID:               item.ID,
			Name:             item.Name,
			IsLiked:          mapTrackActivities[item.ID].IsLiked,
		})
	}

	return &spotify.SearchResponse{
		Limit:  data.Tracks.Limit,
		Offset: data.Tracks.Offset,
		Total:  data.Tracks.Total,
		Items:  items,
	}
}

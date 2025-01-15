package tracks

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rdy24/spotify-catalog/internal/models/spotify"
	"github.com/rdy24/spotify-catalog/pkg/jwt"
	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

func TestHandler_GetRecommendation(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockSvc := NewMockservice(mockCtrl)

	tests := []struct {
		name               string
		expectedStatusCode int
		expectedBody       *spotify.RecommendationResponse
		wantErr            bool
		mockFn             func()
	}{
		{
			name:               "success",
			expectedStatusCode: http.StatusOK,
			expectedBody: &spotify.RecommendationResponse{
				Items: []spotify.SpotifyTrackObject{
					{
						AlbumType:        "album",
						AlbumTotalTracks: 22,
						AlbumImagesURL:   []string{"https://i.scdn.co/image/ab67616d0000b273e8b066f70c206551210d902b", "https://i.scdn.co/image/ab67616d00001e02e8b066f70c206551210d902b", "https://i.scdn.co/image/ab67616d00004851e8b066f70c206551210d902b"},
						AlbumName:        "Bohemian Rhapsody (The Original Soundtrack)",
						ArtistsName:      []string{"Queen"},
						Explicit:         false,
						ID:               "3z8h0TU7ReDPLIbEnYhWZb",
						Name:             "Bohemian Rhapsody",
					},
					{
						AlbumType:        "album",
						AlbumTotalTracks: 12,
						AlbumImagesURL:   []string{"https://i.scdn.co/image/ab67616d0000b273e319baafd16e84f0408af2a0", "https://i.scdn.co/image/ab67616d00001e02e319baafd16e84f0408af2a0", "https://i.scdn.co/image/ab67616d00004851e319baafd16e84f0408af2a0"},
						AlbumName:        "A Night At The Opera (2011 Remaster)",
						ArtistsName:      []string{"Queen"},
						Explicit:         false,
						ID:               "4u7EnebtmKWzUH433cf5Qv",
						Name:             "Bohemian Rhapsody - Remastered 2011",
					},
				},
			},
			wantErr: false,
			mockFn: func() {
				mockSvc.EXPECT().GetRecommendation(gomock.Any(), uint(1), 10, "trackID").Return(&spotify.RecommendationResponse{
					Items: []spotify.SpotifyTrackObject{
						{
							AlbumType:        "album",
							AlbumTotalTracks: 22,
							AlbumImagesURL:   []string{"https://i.scdn.co/image/ab67616d0000b273e8b066f70c206551210d902b", "https://i.scdn.co/image/ab67616d00001e02e8b066f70c206551210d902b", "https://i.scdn.co/image/ab67616d00004851e8b066f70c206551210d902b"},
							AlbumName:        "Bohemian Rhapsody (The Original Soundtrack)",
							ArtistsName:      []string{"Queen"},
							Explicit:         false,
							ID:               "3z8h0TU7ReDPLIbEnYhWZb",
							Name:             "Bohemian Rhapsody",
						},
						{
							AlbumType:        "album",
							AlbumTotalTracks: 12,
							AlbumImagesURL:   []string{"https://i.scdn.co/image/ab67616d0000b273e319baafd16e84f0408af2a0", "https://i.scdn.co/image/ab67616d00001e02e319baafd16e84f0408af2a0", "https://i.scdn.co/image/ab67616d00004851e319baafd16e84f0408af2a0"},
							AlbumName:        "A Night At The Opera (2011 Remaster)",
							ArtistsName:      []string{"Queen"},
							Explicit:         false,
							ID:               "4u7EnebtmKWzUH433cf5Qv",
							Name:             "Bohemian Rhapsody - Remastered 2011",
						},
					},
				}, nil)
			},
		},

		{
			name:               "failed",
			expectedStatusCode: http.StatusBadRequest,
			expectedBody:       nil,
			wantErr:            true,
			mockFn: func() {
				mockSvc.EXPECT().GetRecommendation(gomock.Any(), uint(1), 10, "trackID").Return(nil, assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn()
			api := gin.New()

			h := &Handler{
				Engine:  api,
				service: mockSvc,
			}
			h.RegisterRoutes()
			w := httptest.NewRecorder()

			endpoint := `/tracks/recommendations?limit=10&trackID=trackID`

			req, err := http.NewRequest(http.MethodGet, endpoint, nil)
			assert.NoError(t, err)
			token, err := jwt.CreateToken(1, "username", "")
			assert.NoError(t, err)
			req.Header.Set("Authorization", token)

			h.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)

			if !tt.wantErr {
				res := w.Result()
				defer res.Body.Close()

				response := spotify.RecommendationResponse{}
				err = json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)

				assert.Equal(t, tt.expectedBody, &response)
			}
		})
	}
}

package memberships

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rdy24/spotify-catalog/internal/models/memberships"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestHandler_SignUp(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockSvc := NewMockservice(ctrlMock)

	tests := []struct {
		name               string
		h                  *Handler
		mockFn             func()
		expectedStatusCode int
	}{
		{
			name: "success",
			mockFn: func() {
				mockSvc.EXPECT().SignUp(
					memberships.SignUpRequest{
						Email:    "test@mail.com",
						Password: "password",
						Username: "testuser",
					},
				).Return(nil)
			},
			expectedStatusCode: 201,
		},
		{
			name: "failed",
			mockFn: func() {
				mockSvc.EXPECT().SignUp(
					memberships.SignUpRequest{
						Email:    "test@mail.com",
						Password: "password",
						Username: "testuser",
					},
				).Return(errors.New("failed to sign up user"))
			},
			expectedStatusCode: 400,
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

			endpoint := `/memberships/sign-up`

			model := memberships.SignUpRequest{
				Email:    "test@mail.com",
				Password: "password",
				Username: "testuser",
			}

			val, err := json.Marshal(model)

			assert.NoError(t, err)

			body := bytes.NewReader(val)

			req, err := http.NewRequest("POST", endpoint, body)
			assert.NoError(t, err)

			h.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)
		})
	}
}

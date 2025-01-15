package memberships

import (
	"errors"

	"github.com/rdy24/spotify-catalog/internal/models/memberships"
	"github.com/rdy24/spotify-catalog/pkg/jwt"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (s *service) Login(request memberships.LoginRequest) (string, error) {
	userDetail, err := s.repository.GetUser(request.Email, "", 0)

	if err != nil && err != gorm.ErrRecordNotFound {
		log.Error().Err(err).Msg("failed to get user from database")
		return "", err
	}

	if userDetail == nil {
		return "", errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(userDetail.Password), []byte(request.Password))

	if err != nil {
		log.Error().Err(err).Msg("failed to compare password")
		return "", errors.New("invalid password")
	}

	accessToken, err := jwt.CreateToken(userDetail.ID, userDetail.Username, s.cfg.Service.SecretJWT)

	if err != nil {
		log.Error().Err(err).Msg("failed to create access token")
		return "", err
	}

	return accessToken, nil

}

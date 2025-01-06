package memberships

import (
	"errors"

	"github.com/rdy24/spotify-catalog/internal/models/memberships"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (s *service) SignUp(request memberships.SignUpRequest) error {
	existingUser, err := s.repository.GetUser(request.Email, request.Username, 0)

	if err != nil && err != gorm.ErrRecordNotFound {
		log.Error().Err(err).Msg("failed to get user")
		return err
	}

	if existingUser != nil {
		return errors.New("user already exists")
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)

	if err != nil {
		log.Error().Err(err).Msg("failed to generate password hash")
		return err
	}

	model := memberships.User{
		Email:     request.Email,
		Username:  request.Username,
		Password:  string(pass),
		CreatedBy: "system",
		UpdatedBy: "system",
	}

	return s.repository.CreateUser(model)
}

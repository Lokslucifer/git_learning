package services

import (
	"context"
	"database/sql"
	"errors"
	customerrors "large_fss/internals/customErrors"
	"large_fss/internals/dto"
	"large_fss/internals/models"
)

func (s *Service) Signup(c context.Context, signupRequest *dto.SignUpDTO) (string, error) {
	// We check here if the user is already there
	_, err := s.repo.FindUserByEmail(c, signupRequest.Email)

	if err == nil {
		return "", customerrors.ErrUserAlreadyExists

	} else if !errors.Is(err, sql.ErrNoRows) {

		return "", err
	}

	userModel := models.User{
		Email:     signupRequest.Email,
		Password:  signupRequest.Password,
		FirstName: signupRequest.FirstName,
		LastName:  signupRequest.LastName,
	}

	userId, err := s.repo.CreateUser(c, userModel)
	if err != nil {
		return "", err
	}

	jwt, err := s.JwtService.CreateJWT(userId)

	if err != nil {

		return "", err
	}
	return jwt, nil

}

func (s *Service) Login(c context.Context, loginRequest *dto.LoginDTO) (string, error) {
	user, err := s.repo.FindUserByEmail(c, loginRequest.Email)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {

			return "", customerrors.ErrUserNotFound
		} else {

			return "", err
		}

	}

	if loginRequest.Password != user.Password {

		return "", customerrors.ErrInvalidPassword
	}

	jwt, err := s.JwtService.CreateJWT(user.ID)

	if err != nil {

		return "", err
	}
	return jwt, nil
}

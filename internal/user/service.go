package user

import (
	"ajher-server/utils"
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(input RegisterUserInput) (User, error)
	Login(input LoginUserInput) (User, error)
	GetUserById(userId int) (User, error)
	GoogleAuth(input GoogleOAuthInput) (User, error)
	// GetProfile(ID int) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Register(input RegisterUserInput) (User, error) {
	user := User{}
	user.Email = input.Email
	user.Username = input.Username

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	if err != nil {
		return user, err
	}

	user.Password = string(passwordHash)

	newUser, err := s.repository.Save(user)

	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) Login(input LoginUserInput) (User, error) {
	user, err := s.repository.FindByEmail(input.Email)

	if err != nil {
		return user, errors.New("user with that email doesn't exist")
	}

	return user, nil
}

func (s *service) GetUserById(userId int) (User, error) {
	user, err := s.repository.GetById(userId)

	if err != nil {
		return user, errors.New("user with that id doesn't exist")
	}

	return user, nil
}

func (s *service) GoogleAuth(input GoogleOAuthInput) (User, error) {
	var user User

	response, err := utils.VerifyIdToken(input.OAuthAccessToken)

	if err != nil {
		return user, errors.New("error validation token")
	}

	var googleUser utils.GoogleUser

	googleUser.Id = response.UserId
	googleUser.Email = response.Email
	googleUser.Picture = response.Audience
	googleUser.VerifiedEmail = response.VerifiedEmail

	isUserExist, err := s.repository.FindByEmail(googleUser.Email)

	if err == nil {
		return isUserExist, nil
	}

	userName := strings.Split(googleUser.Email, "@")[0]

	user.Email = googleUser.Email
	user.Username = userName
	user.FullName = userName
	user.Picture = googleUser.Picture

	newUser, err := s.repository.Save(user)

	if err != nil {
		return newUser, err
	}

	return newUser, nil

}

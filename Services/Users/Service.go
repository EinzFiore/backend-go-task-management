package users

import (
	Model "crowdfunding/Model"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterInput) (Model.User, error)
	Login(input LoginInput) (Model.User, error)
	IsEmailAvailable(input CheckEmailInput) (bool, error)
	SaveAvatar(ID int, fileLocation string) (Model.User, error)
	GetUserByID(ID int) (Model.User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterInput) (Model.User, error) {
	user := Model.User{}
	user.Fullname = input.Fullname
	user.Email = input.Email
	user.Avatar = input.Avatar
	user.Description = input.Description
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.Password = string(passwordHash)
	user.RoleId = 1
	user.CreatedAt = time.Now()

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) Login(input LoginInput) (Model.User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) IsEmailAvailable(input CheckEmailInput) (bool, error) {
	email := input.Email
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return false, err
	}

	if user.Id != 0 {
		return true, nil
	}

	return true, nil
}

func (s *service) SaveAvatar(id int, fileLocation string) (Model.User, error) {
	user, err := s.repository.FindByID(id)
	if err != nil {
		return user, err
	}

	user.Avatar = fileLocation

	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}

func (s *service) GetUserByID(id int) (Model.User, error) {
	user, err := s.repository.FindByID(id)
	if err != nil {
		return user, err
	}

	return user, nil
}

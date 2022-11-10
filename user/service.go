package user

import "golang.org/x/crypto/bcrypt"

type Service interface {
	RegisterUser(input RegisterInput) (User, error)
}

type service struct {
	// dependence
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterInput) (User, error) {
	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.Password = string(password)

	user.Status = true

	newUser, err := s.repository.Save(user)
	if err != nil {
		return user, err
	}

	return newUser, nil
}

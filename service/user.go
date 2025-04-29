package service

import (
	"tatar/book/.gen/mydb/model"
	"tatar/book/repository"

	"github.com/google/uuid"
)

type UserServiceImpl struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		repo: repo,
	}
}

func (s *UserServiceImpl) GetUsers(limit int, offset int) ([]*model.User, error) {
	users, err := s.repo.GetUsers(limit, offset)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserServiceImpl) GetUserByUUID(id string) (*model.User, error) {
	user, err := s.repo.GetUserByUUID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserServiceImpl) CreateUser(user *model.User) (*model.User, error) {
	user.UUID = uuid.New().String()
	createdUser, err := s.repo.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return createdUser, nil
}

func (s *UserServiceImpl) UpdateUser(user *model.User) (*model.User, error) {
	updatedUser, err := s.repo.UpdateUser(user)
	if err != nil {
		return nil, err
	}
	return updatedUser, nil
}

func (s *UserServiceImpl) DeleteUser(id string) error {
	err := s.repo.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}

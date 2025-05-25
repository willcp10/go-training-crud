package repository

import (
	"errors"

	"go-training-crud/data_source"
	"go-training-crud/internal/domain"
)

var _ UserRepository = new(UserRepositoryImpl)

type UserRepository interface {
	CreateUser(u domain.User) error
	FindAll() []domain.User
	FindByID(domain.ID) (domain.User, error)
	DeleteUser(id domain.ID) error
	UpdateUser(u domain.User) error
}

type UserRepositoryImpl struct {
	userDataSource data_source.UserDataSource
}

func (ur *UserRepositoryImpl) CreateUser(u domain.User) error {
	userModel := FromDomainMapper(u)

	err := ur.userDataSource.Insert(data_source.UserDataInput(userModel.ID, userModel.Name, userModel.Age, userModel.DocNumber))
	if err != nil {
		return errors.Join(errors.New("error creating user"), err)
	}

	return nil
}

func (ur *UserRepositoryImpl) FindAll() []domain.User {
	var userSlice []domain.User
	for _, u := range ur.userDataSource.SelectAll() {
		var userModel UserModel
		userModel.ID, userModel.Name, userModel.Age, userModel.DocNumber = data_source.UserDataOutput(u)
		userSlice = append(userSlice, ToDomainMapper(userModel))
	}

	return userSlice
}

func (ur *UserRepositoryImpl) FindByID(id domain.ID) (domain.User, error) {
	var userModel UserModel
	u, found := ur.userDataSource.Select(id.Int64())
	if !found {
		return domain.User{}, errors.New("user not found")
	}

	userModel.ID, userModel.Name, userModel.Age, userModel.DocNumber = data_source.UserDataOutput(u)
	return ToDomainMapper(userModel), nil
}

func (ur *UserRepositoryImpl) DeleteUser(id domain.ID) error {
	err := ur.userDataSource.Delete(id.Int64())
	if err != nil {
		return errors.Join(errors.New("error deleting user"), err)
	}

	return nil
}

func (ur *UserRepositoryImpl) UpdateUser(u domain.User) error {
	userModel := FromDomainMapper(u)

	err := ur.userDataSource.Update(data_source.UserDataInput(userModel.ID, userModel.Name, userModel.Age, userModel.DocNumber))
	if err != nil {
		return errors.Join(errors.New("error updating user"), err)
	}

	return nil
}

func NewUserRepository(userDataSource data_source.UserDataSource) UserRepository {
	return &UserRepositoryImpl{
		userDataSource: userDataSource,
	}
}

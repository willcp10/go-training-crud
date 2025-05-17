package cmd

import (
	"go-training-crud/internal/adapters/repository"
	"go-training-crud/internal/domain"
)

type UserCommandService interface {
	UserCreator
	UserUpdater
	UserDeleter
}

type UserCreator interface {
	CreateUser(cmd CreateUserCommand) error
}

type UserUpdater interface {
	UpdateUser(cmd UpdateUserCommand) error
}

type UserDeleter interface {
	DeleteUser(cmd DeleteUserCommand) error
}

var _ UserCommandService = new(UserCommandServiceImpl)

type UserCommandServiceImpl struct {
	userRepository repository.UserRepository
}

func NewUserCommandService(userRepository repository.UserRepository) UserCommandService {
	return &UserCommandServiceImpl{
		userRepository: userRepository,
	}
}

type CreateUserCommand struct {
	User domain.User
}

func NewCreateUserCommand(user domain.User) CreateUserCommand {
	return CreateUserCommand{
		User: user,
	}
}

type UpdateUserCommand struct {
	User domain.User
}

func NewUpdateUserCommand(user domain.User) UpdateUserCommand {
	return UpdateUserCommand{
		User: user,
	}
}

type DeleteUserCommand struct {
	ID domain.ID
}

func NewDeleteUserCommand(id domain.ID) DeleteUserCommand {
	return DeleteUserCommand{
		ID: id,
	}
}
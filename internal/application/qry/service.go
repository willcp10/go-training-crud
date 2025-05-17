package qry

import (
	"go-training-crud/internal/adapters/repository"
	"go-training-crud/internal/domain"
)

type UserQueryService interface {
	AllUsersFinder
	ByIDUserFinder
}

type AllUsersFinder interface {
	FindAllUsers() []domain.User
}

type ByIDUserFinder interface {
	FindUserByID(id domain.ID) (domain.User, error)
}

var _ UserQueryService = new(UserQueryServiceImpl)

type UserQueryServiceImpl struct {
	userRepository repository.UserRepository
}

func NewUserQueryService(userRepository repository.UserRepository) UserQueryService {
	return &UserQueryServiceImpl{
		userRepository: userRepository,
	}
}
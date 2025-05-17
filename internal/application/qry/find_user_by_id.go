package qry

import "go-training-crud/internal/domain"

func (u *UserQueryServiceImpl) FindUserByID(id domain.ID) (domain.User, error) {
	return u.userRepository.FindByID(id)
}
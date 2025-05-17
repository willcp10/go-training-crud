package qry

import "go-training-crud/internal/domain"

func (u *UserQueryServiceImpl) FindAllUsers() []domain.User {
	return u.userRepository.FindAll()
}
package repository

import "go-training-crud/internal/domain"

type UserModel struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Age       int    `json:"age"`
	DocNumber string `json:"doc_number"`
}

func FromDomainMapper(user domain.User) UserModel {
	return UserModel{
		ID: user.ID().Int64(),
		Name: user.Name().String(),
		Age: user.Age().Int(),
		DocNumber: user.DocNumber().String(),
	}
}

func ToDomainMapper(userModel UserModel) domain.User {
	return domain.NewUser(
		domain.NewID(userModel.ID),
		domain.NewName(userModel.Name),
		domain.NewAge(userModel.Age),
		domain.NewDocNumber(userModel.DocNumber),
	)
}
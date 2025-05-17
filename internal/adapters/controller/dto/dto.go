package dto

import "go-training-crud/internal/domain"

type UserDTO struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Age       int    `json:"age"`
	DocNumber string `json:"doc_number"`
}

func (dto UserDTO) ToDomain() domain.User {
	return domain.NewUser(
		domain.NewID(dto.ID),
		domain.NewName(dto.Name),
		domain.NewAge(dto.Age),
		domain.NewDocNumber(dto.DocNumber),
	)
}

func FromDomain(user domain.User) UserDTO {
	return UserDTO{
		ID:        user.ID().Int64(),
		Name:      user.Name().String(),
		Age:       user.Age().Int(),
		DocNumber: user.DocNumber().String(),
	}
}

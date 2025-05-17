package domain

type User struct {
	id        ID
	name      Name
	age       Age
	docNumber DocNumber
}

func (u *User) ID() ID {
	return u.id.ID()
}

func (u *User) Name() Name {
	return u.name.Name()
}

func (u *User) Age() Age {
	return u.age.Age()
}

func (u *User) DocNumber() DocNumber {
	return u.docNumber.DocNumber()
}

func NewUser(id ID, name Name, age Age, docNumber DocNumber) User {
	return User{
		id: id,
		name: name,
		age: age,
		docNumber: docNumber,
	}
}
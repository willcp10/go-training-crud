package domain

type ID int64

func (id ID) ID() ID {
	return id
}

func (id ID) Int64() int64 {
	return int64(id)
}

func NewID(id int64) ID {
	return ID(id)
}
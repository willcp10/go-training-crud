package domain

type Name string

func (name Name) Name() Name {
	return name
}

func (name Name) String() string {
	return string(name)
}

func NewName(name string) Name {
	return Name(name)
}
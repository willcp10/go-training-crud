package domain

type Age int

func (age Age) Age() Age {
	return age
}

func (age Age) Int() int {
	return int(age)
}

func NewAge(age int) Age {
	return Age(age)
}
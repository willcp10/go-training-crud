package domain

type DocNumber string

func (docNumber DocNumber) DocNumber() DocNumber {
	return docNumber
}

func (docNumber DocNumber) String() string {
	return string(docNumber)
}

func NewDocNumber(docNumber string) DocNumber {
	return DocNumber(docNumber)
}
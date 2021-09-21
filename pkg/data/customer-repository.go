package data

type CustomerRepository interface {
	GetAll(code string) ([]*Customer, error)
}

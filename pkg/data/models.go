package data

import "database/sql"

type Models struct {
	Customers CustomerRepository
}

func NewModels(db *sql.DB) Models {
	return Models{
		Customers: CustomerModel{DB: db},
	}
}

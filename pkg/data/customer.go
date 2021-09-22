package data

import (
	"database/sql"
	"fmt"
)

var CountryCodes = map[string]string{
	"(237)": "Cameron",
	"(251)": "Ethiopia",
	"(212)": "Morocco",
	"(258)": "Mozambique",
	"(256)": "Uganda",
}

type Customer struct {
	ID          int
	Name        string
	PhoneNumber string
}

type CustomerModel struct {
	DB *sql.DB
}

func (m CustomerModel) GetAll(code string) ([]*Customer, error) {

	stmt := `SELECT id, name, phone FROM customer WHERE phone LIKE ?`
	rows, err := m.DB.Query(stmt, code+"%")
	if err != nil {
		return nil, err
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	var customers []*Customer

	for rows.Next() {
		customer := &Customer{}

		err := rows.Scan(&customer.ID, &customer.Name, &customer.PhoneNumber)
		if err != nil {
			return nil, err
		}

		customers = append(customers, customer)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return customers, err
}

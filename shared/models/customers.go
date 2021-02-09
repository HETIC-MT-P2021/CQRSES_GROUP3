package models

import (
	"errors"
	"unicode"

	"github.com/asaskevich/govalidator"
)

// Customer is our struct for users
type Customer struct {
	ID             uint64 `gorm:"primary_key"`
	Name           string `gorm:"size:255"`
	Email          string `gorm:"size:255; unique"`
	HashedPassword string
}

// CustomerForm is our struct to handle new users requests
type CustomerForm struct {
	Name     string `gorm:"size:255"`
	Email    string `gorm:"size:255" valid:"email~Invalid email address"`
	Password string `gorm:"size:255"`
}

// CustomerJSON is the struct to return user without the hash password
type CustomerJSON struct {
	ID          uint64
	Name, Email string
}

// ValidateCustomer takes a customer form as parameter and check if its properties are valid
func ValidateCustomer(customer *CustomerForm) error {
	_, err := govalidator.ValidateStruct(customer)
	if err != nil {
		return err
	}
	if customer.Name == "" {
		return errors.New("Invalid name")
	}

	if valid := verifyPassword(customer.Password); valid == false {
		return errors.New("Password must be composed of at least 8 character, one uppercase letter and one number")
	}

	return nil
}

func verifyPassword(s string) bool {
	number := false
	upper := false

	for _, c := range s {
		switch {
		case unicode.IsNumber(c):
			number = true
		case unicode.IsUpper(c):
			upper = true
		default:
		}
	}

	return number && upper && len(s) >= 8
}

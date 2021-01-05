package repositories

import (
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/database"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/models"
)

// PersistCustomer is used to persist user objects to database
func PersistCustomer(customer *models.Customer) error {
	
	if err := database.Db.Debug().Create(&customer).Error; err != nil {
		return err
	}

	return nil
}

// FindCustomerByEmail receives a pointer to a customer and try to find a customer with the email address corresponding
func FindCustomerByEmail(customer *models.Customer) error {
	if err := database.Db.Debug().Where("email = ?", customer.Email).First(&customer).Error; err != nil {
		return err
	}

	return nil
}
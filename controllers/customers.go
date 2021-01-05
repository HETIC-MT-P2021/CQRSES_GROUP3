package controllers

import (
	"github.com/edwinvautier/go-boilerplate/repositories"
	"github.com/edwinvautier/go-boilerplate/services"
	"net/http"

	"github.com/edwinvautier/go-boilerplate/models"
	"github.com/gin-gonic/gin"
)

// CreateCustomer is the controller to create a new customer
func CreateCustomer(c *gin.Context) {
	var customerForm models.CustomerForm
	if err := c.ShouldBindJSON(&customerForm); err != nil {
		c.JSON(http.StatusBadRequest, "Invalid informations provided, please refer to the documentation.")
		return
	}

	// Validate user structure (email, password)
	if err := models.ValidateCustomer(&customerForm); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// Create customer object and persist it to db
	customer := models.Customer {
		Name: customerForm.Name,
		Email: customerForm.Email,
		HashedPassword: services.HashPassword(customerForm.Password),
	}

	if err := repositories.PersistCustomer(&customer); err != nil {
		c.JSON(http.StatusInternalServerError, "Couldn't create user. Try again.")
		return
	}

	c.JSON(http.StatusOK, models.CustomerJSON{ID: customer.ID, Name: customer.Name, Email: customer.Email})
}

func TestController(c *gin.Context) {
	c.JSON(http.StatusOK, "OKKKKKK")
}
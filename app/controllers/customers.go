package controllers

import (
	"net/http"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/app/services"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/shared/database"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/shared/models"
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
	customer := models.Customer{
		Name:           customerForm.Name,
		Email:          customerForm.Email,
		HashedPassword: services.HashPassword(customerForm.Password),
	}
	
	repository := database.Init()

	if err := repository.PersistCustomer(&customer); err != nil {
		c.JSON(http.StatusInternalServerError, "Couldn't create user. Try again.")
		return
	}

	c.JSON(http.StatusOK, models.CustomerJSON{ID: customer.ID, Name: customer.Name, Email: customer.Email})
}

func TestController(c *gin.Context) {
	c.JSON(http.StatusOK, "OKKKKKK")
}

package controllers

import (
	"net/http"
	"strconv"

	"github.com/edwinvautier/go-boilerplate/config"
	"github.com/edwinvautier/go-boilerplate/models"
	"github.com/edwinvautier/go-boilerplate/repositories"
	"github.com/edwinvautier/go-boilerplate/services"
	"github.com/gin-gonic/gin"
)

// Login takes customer email and password a json params and returns a token or an error depending on the credentials given.
func Login(c *gin.Context) {
	var customerForm models.CustomerForm
	if err := c.ShouldBindJSON(&customerForm); err != nil {
		c.JSON(http.StatusBadRequest, "Incorrect user informations")
		return
	}

	// Try to find user with this address
	customer := models.Customer{
		Email: customerForm.Email,
	}
	if err := repositories.FindCustomerByEmail(&customer); err != nil {
		c.JSON(http.StatusUnauthorized, "incorrect email or password.")
		return
	}

	// Verify password
	hashedPwd := services.HashPassword(customerForm.Password)
	if hashedPwd != customer.HashedPassword {
		c.JSON(http.StatusUnauthorized, "incorrect email or password.")
		return
	}

	// Generate connection token
	token, err := services.GenerateToken(customer.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Couldn't create your authorization")
		return
	}
	validTime, _ := strconv.ParseInt(config.GoDotEnvVariable("TOKEN_VALID_DURATION"), 10, 64)

	c.SetCookie("token", token, 60*int(validTime), "/", config.GoDotEnvVariable("DOMAIN"), false, false)
	c.JSON(http.StatusOK, "Logged in successfully")
}

package tests

import (
	"testing"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/shared/models"
)

func TestControllers(t *testing.T) {
	initDB()
	initRouter()

	credentials := models.CustomerForm{
		Email:    "john@doe.com",
		Name:     "John Doe",
		Password: "Test123456",
	}

	testRegister(t, credentials)
	token := testLogin(t, credentials)

	testRoot(t, token)
}

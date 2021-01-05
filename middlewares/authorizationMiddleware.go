package middlewares

import (
	"net/http"
	"strings"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/models"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/repositories"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/services"
	"github.com/gin-gonic/gin"
)

// CheckAuthorization is the middleware used to find token and check it from request authorization header
func CheckAuthorization(c *gin.Context) {
	// Get and try to split the Bearer token
	bearer := c.GetHeader("Authorization")
	if !strings.HasPrefix(bearer, "Bearer") {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Unauthorized")
		return
	}

	token := strings.TrimPrefix(bearer, "Bearer ")

	// validate and decode token to get its informations
	_, claims, err := services.DecodeToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Unauthorized")
		return
	}

	customer := models.Customer{
		Email: claims.Email,
	}

	if err := repositories.FindCustomerByEmail(&customer); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Unauthorized")
		return
	}

	c.Next()
}

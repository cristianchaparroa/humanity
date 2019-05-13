package api

import (
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

// AuthRequired verifies  the user in the system
// still loged.
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user")
		if user == nil {
			// You'd normally redirect to login page
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
		} else {
			// Continue down the chain to handler etc
			c.Next()
		}
	}
}

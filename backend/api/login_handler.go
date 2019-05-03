package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

// UserLogin is an structure to check the login
type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginHandler makes a basic login with user an password
func LoginHandler(c *gin.Context) {
	fmt.Println("LoginHandler")
	session := sessions.Default(c)

	var u UserLogin
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	username := u.Email
	password := u.Password

	fmt.Println(username)
	fmt.Println(password)

	if strings.Trim(username, " ") == "" || strings.Trim(password, " ") == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Parameters can't be empty"})
		return
	}

	// TODO: use a repository to compare
	if username == "cristianchaparroa@gmail.com" && password == "12345" {
		session.Set("user", username) //In real world usage you'd set this to the users ID
		err := session.Save()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate session token"})
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "Successfully authenticated user"})
		}
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
	}
}

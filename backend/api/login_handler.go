package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/cristianchaparroa/humanity/backend/repositories"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

// UserLogin is an structure to check the login
type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginHandler makes a basic login with user an password
func LoginHandler(db *sql.DB) gin.HandlerFunc {

	fn := func(c *gin.Context) {
		fmt.Println("LoginHandler")
		session := sessions.Default(c)

		var u UserLogin
		if err := c.ShouldBindJSON(&u); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		username := u.Email
		password := u.Password

		if strings.Trim(username, " ") == "" || strings.Trim(password, " ") == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Parameters can't be empty"})
			return
		}

		// TODO: use a repository to compare
		if isLogin(db, u) {
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

	return gin.HandlerFunc(fn)
}

// TODO: move it to a service layer
func isLogin(db *sql.DB, u UserLogin) bool {
	ar := repositories.NewAccountRepository(db)
	account, err := ar.FindByEmail(u.Email)

	if err != nil {
		return false
	}

	fmt.Println(u.Email)
	fmt.Println(u.Password)
	fmt.Println(account.ID)

	if u.Password == account.Password {
		return true
	}

	return false
}

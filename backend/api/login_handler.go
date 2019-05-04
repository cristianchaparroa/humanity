package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/cristianchaparroa/humanity/backend/services"
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

		if strings.Trim(u.Email, " ") == "" || strings.Trim(u.Password, " ") == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Parameters can't be empty"})
			return
		}

		as := services.NewAccountService(db)
		isLogin, acc := as.Login(u.Email, u.Password)

		fmt.Printf("Is user in db:%v", (acc != nil))

		if isLogin {
			session.Set("email", u.Email) //In real world usage you'd set this to the users ID
			session.Set("userId", acc.ID)
			session.Set("nickname", acc.Nickname)

			err := session.Save()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate session token"})
			} else {
				c.JSON(http.StatusOK, gin.H{"message": "Successfully authenticated user", "user_id": acc.ID, "nickname": acc.Nickname})
			}
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		}
	}

	return gin.HandlerFunc(fn)
}

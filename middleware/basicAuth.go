package middleware

import (
	"fmt"
	"gotify-project/database"
	"gotify-project/repository"

	"github.com/gin-gonic/gin"
)

func BasicAuth(c *gin.Context) {
	// Get the Basic Authentication credentials
	username, password, hasAuth := c.Request.BasicAuth()

	if hasAuth {

		users, err := repository.Login(database.DbConnection, username, password)

		if err != nil || len(users) == 0 {
			fmt.Println("username atau password salah")
			c.Abort()
			c.Writer.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
			c.JSON(403, gin.H{"error": "username atau password salah"})
		} else {

			if username == users[0].Username && password == users[0].Password {
				fmt.Println("User authenticated")
			} else {
				fmt.Println("username atau password salah")
				c.Abort()
				c.Writer.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
				c.JSON(403, gin.H{"error": "username atau password salah"})
			}
		}
	} else {
		fmt.Println("empty basic auth")
		c.Abort()
		c.Writer.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
		c.JSON(401, gin.H{"error": "empty basic auth"})
	}
}

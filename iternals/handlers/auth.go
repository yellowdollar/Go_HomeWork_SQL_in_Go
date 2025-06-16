package handlers

import (
	"home_work_sql_gin/iternals"
	"home_work_sql_gin/iternals/models"
	"home_work_sql_gin/iternals/repositories"

	"github.com/gin-gonic/gin"
)

func SignUpHandler(c *gin.Context) {
	login := c.Query("login")
	password := c.Query("password")

	userData := models.AuthUser{
		Login:    login,
		Password: password,
	}

	err := repositories.SignUp(iternals.Db, userData)

	if err != nil {
		c.JSON(400, gin.H{"status_code": 400, "message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"status_code": 200, "message": "User Added"})
	return
}

func SignInHandler(c *gin.Context) {

	login := c.Query("login")
	password := c.Query("password")

	token := repositories.SignIn(login, password, iternals.Db)

	c.JSON(200, gin.H{"status_code": 200, "access_token": token, "token_type": "Bearer"})

}

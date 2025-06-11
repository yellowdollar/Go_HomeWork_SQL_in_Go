package handlers

import (
	"home_work_sql_gin/iternals"
	"home_work_sql_gin/iternals/models"
	"home_work_sql_gin/iternals/repositories"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateUTable(c *gin.Context) {
	repositories.CreateUsersTable(iternals.Db)

	c.JSON(200, gin.H{"status_code": 200, "message": "Table was created"})
	return
}

func AddUser(c *gin.Context) {
	name := c.Query("name")
	email := c.Query("email")

	user := models.User{
		Name:  name,
		Email: email,
	}

	err := repositories.AddNewUser(user, iternals.Db)

	if err != nil {
		c.JSON(400, gin.H{"status_code": 400, "message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"status_code": 200, "message": "Row Added"})
	return
}

func GetAllU(c *gin.Context) {
	users, err := repositories.GetAllUsers(iternals.Db)

	if err != nil {
		c.JSON(400, gin.H{"status_code": 400, "message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"status_code": 200, "data": users})
	return
}

func GetUById(c *gin.Context) {
	idStr := c.Query("id")

	id, errConv := strconv.Atoi(idStr)

	if errConv != nil {
		c.JSON(400, gin.H{"status_code": 400, "message": errConv.Error()})
		return
	}

	user, err := repositories.GetUserById(id, iternals.Db)

	if err != nil {
		c.JSON(400, gin.H{"status_code": 400, "message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"status_code": 200, "data": user})
}

func UpdateUEmail(c *gin.Context) {
	idStr := c.Query("id")
	userEmail := c.Query("email")

	id, errConv := strconv.Atoi(idStr)

	if errConv != nil {
		c.JSON(400, gin.H{"status_code": 400, "message": errConv.Error()})
		return
	}

	err := repositories.UpdateUserEmail(id, userEmail, iternals.Db)

	if err != nil {
		c.JSON(400, gin.H{"status_code": 400, "message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"status_code": 200, "message": "Row Updated"})
	return
}

func DeleteUByName(c *gin.Context) {
	userName := c.Query("name")

	err := repositories.DeleteUserByName(userName, iternals.Db)

	if err != nil {
		c.JSON(400, gin.H{"status_code": 400, "message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"status_code": 200, "message": "Row Deleted"})
	return
}

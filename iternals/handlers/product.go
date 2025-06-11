package handlers

import (
	"home_work_sql_gin/iternals"
	"home_work_sql_gin/iternals/models"
	"home_work_sql_gin/iternals/repositories"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateTable(c *gin.Context) {
	repositories.CreateTable(iternals.Db)
}

func Add(c *gin.Context) {
	name := c.Query("name")
	priceStr := c.Query("price")

	price, err := strconv.Atoi(priceStr)

	if err != nil {
		c.JSON(400, gin.H{"status_code": 400, "error_message": "Введите корректное числовое значение в параметр Price"})
	}

	var item = models.Product{
		Name:  name,
		Price: price,
	}

	addNewProduct, err := repositories.AddNewProduct(iternals.Db, item)

	if err != nil {
		c.JSON(400, gin.H{"status_code": 400, "error_message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"status_code": 200, "data": addNewProduct})
	return
}

func GetAll(c *gin.Context) {
	items, err := repositories.GetAllProducts(iternals.Db)

	if err != nil {
		c.JSON(400, gin.H{"status_code": 400, "error_message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"status_code": 200, "data": items})
	return
}

func GetById(c *gin.Context) {
	idStr := c.Query("id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.JSON(400, gin.H{"status_code": 400, "error_message": err.Error()})
		return
	}

	item, err1 := repositories.GetProductById(id, iternals.Db)

	if err1 != nil {
		c.JSON(400, gin.H{"status_code": 400, "error_message": err1.Error()})
		return
	}

	c.JSON(200, gin.H{"status_code": 200, "data": item})
	return
}

func UpdateById(c *gin.Context) {
	idStr := c.Query("id")
	priceStr := c.Query("price")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.JSON(400, gin.H{"status_code": 400, "error_message": err.Error()})
		return
	}

	price, err1 := strconv.Atoi(priceStr)

	if err1 != nil {
		c.JSON(400, gin.H{"status_code": 400, "error_message": err1.Error()})
		return
	}

	err2 := repositories.UpdateProductPriceById(id, price, iternals.Db)

	if err2 != nil {
		c.JSON(400, gin.H{"status_code": 400, "error_message": err2.Error()})
		return
	}

	c.JSON(200, gin.H{"status_code": 200, "message": "Item was updated"})
	return
}

func DeleteById(c *gin.Context) {
	idStr := c.Query("id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.JSON(400, gin.H{"status_code": 400, "message": err.Error()})
		return
	}

	err1 := repositories.DeleteProductById(id, iternals.Db)

	if err1 != nil {
		c.JSON(400, gin.H{"status_code": 400, "message": err1.Error()})
		return
	}

	c.JSON(200, gin.H{"status_code": 200, "message": "Row deleted"})
	return
}

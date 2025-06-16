package main

import (
	"home_work_sql_gin/iternals"
	"home_work_sql_gin/iternals/handlers"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {

	r := gin.Default()

	iternals.InitDB()

	productsHandler := r.Group("/products")
	{
		productsHandler.POST("/addProduct", handlers.Add)
		productsHandler.GET("/getProducts", handlers.GetAll)
		productsHandler.GET("/getProduct", handlers.GetById)
		productsHandler.PUT("/updateProduct", handlers.UpdateById)
		productsHandler.DELETE("/deleteProduct", handlers.DeleteById)
	}

	usersHandler := r.Group("/users")
	{
		usersHandler.POST("/createTable", handlers.CreateUTable)
		usersHandler.POST("/addUser", handlers.AddUser)
		usersHandler.GET("/getUsers", handlers.GetAllU)
		usersHandler.GET("/getUser", handlers.GetUById)
		usersHandler.PUT("/updateUser", handlers.UpdateUEmail)
		usersHandler.DELETE("deleteUser", handlers.DeleteUByName)
	}

	authHandler := r.Group("/auth")
	{
		authHandler.POST("/sign_up", handlers.SignUpHandler)
		authHandler.POST("/sign_in", handlers.SignInHandler)
	}

	r.Run(":8080")
}

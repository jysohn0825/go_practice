package main

import (
	"toy/controllers"
	"toy/models"

	_ "toy/docs"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title 			Book API
// @version 		1.0
// @description		This is a Book api server for test gin, gorm
// @contact.name 	jy.son toy project API
// @contact.email 	jy.son@konai.com
// @host localhost
// @BasePath /api
func main() {
	r := gin.Default()

	// Connect to database
	models.ConnectDatabase()

	// Routes
	r.GET("/books", controllers.FindBooks)
	r.GET("/book/:id", controllers.FindBook)
	r.POST("/book", controllers.CreateBook)
	r.PATCH("/book/:id", controllers.UpdateBook)
	r.DELETE("/book/:id", controllers.DeleteBook)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// Run the server
	r.Run()
}

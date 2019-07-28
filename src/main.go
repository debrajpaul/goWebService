package main

import (
	"goWebService/src/handler"
	"goWebService/src/modules/inventory"

	"github.com/gin-gonic/gin"
)

func main() {
	inventory.InitialMigration()
	product := inventory.NewInstance()
	r := gin.Default()

	r.GET("/items", handler.GetItems())
	r.GET("/item/:name", handler.GetItem(product))
	r.POST("/item", handler.PostItems(product))
	r.DELETE("/item/:name", handler.DeleteItem(product))

	r.Run() // listen and serve on 0.0.0.0:8080
}

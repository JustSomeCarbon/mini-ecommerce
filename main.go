package main

import (
	"log"
	"os"

	"github.com/JustSomeCarbon/mini-ecommerce/controllers"
	"github.com/JustSomeCarbon/mini-ecommerce/database"
	"github.com/JustSomeCarbon/mini-ecommerce/middleware"
	"github.com/JustSomeCarbon/mini-ecommerce/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.InstantBut())

	log.Fatal(router.Run(":" + port))
}

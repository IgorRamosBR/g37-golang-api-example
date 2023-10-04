package main

import (
	"postech/controllers"
	"postech/repositories"
	"postech/services"

	"github.com/gin-gonic/gin"
)

func main() {
	customerRepository := repositories.NewCustomerRepository()
	customerService := services.NewCustomerService(customerRepository)
	customerController := controllers.CustomerController{
		CustomerService: customerService,
	}

	router := gin.Default()
	router.GET("/customers", customerController.GetAllCustomers)
	router.POST("/customers", customerController.CreateCustomer)
	router.Run() // listen and serve on 0.0.0.0:8080
}

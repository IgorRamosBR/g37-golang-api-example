package controllers

import (
	"fmt"
	"postech/models"
	"postech/services"

	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	CustomerService services.CustomerService
}

func (c CustomerController) GetAllCustomers(ctx *gin.Context) {
	customers := c.CustomerService.GetAllCustomers()

	ctx.JSON(200, customers)
}

func (c CustomerController) CreateCustomer(ctx *gin.Context) {
	var customer models.Customer
	err := ctx.BindJSON(&customer)
	if err != nil {
		ctx.JSON(400, fmt.Sprintf("error to bind customer, error: %s", err.Error()))
		return
	}

	err = c.CustomerService.SaveCustomer(customer)
	if err != nil {
		ctx.JSON(400, fmt.Sprintf("error to save customer, error: %s", err.Error()))
		return
	}

	ctx.JSON(200, "customer created.")
}

package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"macus/database"
	"macus/model"

	"github.com/gin-gonic/gin"
)

type CustomerHandler struct {
	CustomerDB database.CustomerRepository
}

func (h CustomerHandler) RegisterApi(engine *gin.Engine) {
	v1 := engine.Group("/api/v1")
	{
		customers := v1.Group("/customers")
		{
			customers.GET(":id", h.GetCustomer)
			customers.GET("", h.ListCustomers)
			customers.POST("", h.CreateCustomer)
			customers.DELETE(":id", h.DeleteCustomer)
			customers.PATCH(":id", h.UpdateCustomer)
		}
	}
}

// GetCustomer godoc
// @summary Get Customer
// @description  Get customer by id
// @tags customers
// @security ApiKeyAuth
// @id GetCustomer
// @accept json
// @produce json
// @param id path int true "id of customer to be gotten"
// @response 200 {object} model.Customer "OK"
// @response 400 {object} model.Response "Bad Request"
// @response 401 {object} model.Response "Unauthorized"
// @response 409 {object} model.Response "Conflict"
// @response 500 {object} model.Response "Internal Server Error"
// @Router /api/v1/customers/:id [get]
func (h CustomerHandler) GetCustomer(c *gin.Context) {
	//Logic goes here
	fmt.Println("GetCustomer")

	customerID := c.Param("id")
	fmt.Println(customerID)

	// management := database.PostgresDB{
	// 	DB: h.DB,
	// }
	// customer, err := management.GetCustomerByID(customerID)
	customer, err := h.CustomerDB.GetCustomerByID(customerID)
	if err != nil {
		fmt.Println("Handlers GetProductByID error: ", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, customer)
}

// ListCustomers godoc
// @summary List Customers
// @description List all customers
// @tags customers
// @security ApiKeyAuth
// @id ListCustomers
// @accept json
// @produce json
// @response 200 {object} model.Customers "OK"
// @response 400 {object} model.Response "Bad Request"
// @response 401 {object} model.Response "Unauthorized"
// @response 409 {object} model.Response "Conflict"
// @response 500 {object} model.Response "Internal Server Error"
// @Router /api/v1/customers [get]
func (h CustomerHandler) ListCustomers(c *gin.Context) {
	fmt.Println("Logic goes here")
	//Logic goes here
	resp, err := http.Get("https://reqres.in/api/users?page=2")
	if err != nil {
		fmt.Println("No response from request")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	c.JSON(http.StatusOK, string(body))
}

// CreateCustomer godoc
// @summary Create Customer
// @description Create new customer
// @tags customers
// @security ApiKeyAuth
// @id CreateCustomer
// @accept json
// @produce json
// @param Customer body model.CustomerForCreate true "Customer data to be created"
// @response 200 {object} model.Response "OK"
// @response 400 {object} model.Response "Bad Request"
// @response 401 {object} model.Response "Unauthorized"
// @response 500 {object} model.Response "Internal Server Error"
// @Router /api/v1/customers [post]
func (h CustomerHandler) CreateCustomer(c *gin.Context) {
	fmt.Println("CreateNewProduct ShouldBindJSON ")
	var newCustomer model.NewCustomer
	err := c.ShouldBindJSON(&newCustomer)
	if err != nil {
		fmt.Println("CreateNewProduct ShouldBindJSON error: ", err)
		c.Status(http.StatusBadRequest)
		return
	}
	var customerEntity = database.CustomerEntity{
		Name: newCustomer.Name,
	}
	// management := database.PostgresDB{
	// 	DB: h.DB,
	// }
	// newItemp, err := management.CreateNewCustomer(customerEntity, time.Now())
	newItemp, err := h.CustomerDB.CreateNewCustomer(customerEntity, time.Now())
	if err != nil {
		fmt.Println("Handlers CreateNewProduct error: ", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	fmt.Print("new", newItemp)
	c.JSON(http.StatusOK, newItemp)
}

// DeleteCustomer godoc
// @summary Delete Customer
// @description Delete customer by id
// @tags customers
// @security ApiKeyAuth
// @id DeleteCustomer
// @accept json
// @produce json
// @param id path int true "id of customer to be deleted"
// @response 200 {object} model.Response "OK"
// @response 400 {object} model.Response "Bad Request"
// @response 401 {object} model.Response "Unauthorized"
// @response 500 {object} model.Response "Internal Server Error"
// @Router /api/v1/customers/:id [delete]
func (h CustomerHandler) DeleteCustomer(c *gin.Context) {
	//Logic goes here
}

// UpdateCustomer godoc
// @summary Update Customer
// @description Update customer by id
// @tags customers
// @security ApiKeyAuth
// @id UpdateCustomer
// @accept json
// @produce json
// @param id path int true "id of customer to be updated"
// @param Customer body model.CustomerForUpdate true "Customer data to be updated"
// @response 200 {object} model.Response "OK"
// @response 400 {object} model.Response "Bad Request"
// @response 401 {object} model.Response "Unauthorized"
// @response 500 {object} model.Response "Internal Server Error"
// @Router /api/v1/customers/:id [patch]
func (h CustomerHandler) UpdateCustomer(c *gin.Context) {
	//Logic goes here
}

package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"macus/data"
	"macus/entities"
	"macus/models"
	"macus/services"

	"github.com/gin-gonic/gin"
)

type CustomerHandler struct {
	customerRepository services.CustomerService
}

func NewCustomerHandler(db data.DbProvider) *CustomerHandler {
	repo := services.NewCustomerRepository(db)
	return &CustomerHandler{customerRepository: repo}
}

func (h CustomerHandler) GetCustomer(c *gin.Context) {
	fmt.Println("GetCustomer")
	customerID := c.Param("id")
	fmt.Println(customerID)
	customer, err := h.customerRepository.GetCustomerByID(customerID)
	if err != nil {
		fmt.Println("Handlers GetProductByID error: ", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, customer)
}

func (h CustomerHandler) ListCustomers(c *gin.Context) {
	resp, err := http.Get("https://reqres.in/api/users?page=2")
	if err != nil {
		fmt.Println("No response from request")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	c.JSON(http.StatusOK, string(body))
}

func (h CustomerHandler) CreateCustomer(c *gin.Context) {
	var newCustomer models.NewCustomer
	err := c.ShouldBindJSON(&newCustomer)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	var customerEntity = entities.CustomerEntity{
		Name: newCustomer.Name,
	}
	newItemp, err := h.customerRepository.CreateNewCustomer(customerEntity, time.Now())
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	fmt.Print("new", newItemp)
	c.JSON(http.StatusOK, newItemp)
}

func (h CustomerHandler) DeleteCustomer(c *gin.Context) {
	//Logic goes here
}

func (h CustomerHandler) UpdateCustomer(c *gin.Context) {
	//Logic goes here
}

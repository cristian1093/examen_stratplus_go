package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/models"
)

type Order struct {
	ID         int       `json:"id"`
	UserId     int       `json:"user_id"`
	BondId     int       `json:"bond_id"`
	OrderType  string    `json:"order_type"`
	NumBonds   int       `json:"num_bonds"`
	Price      float64   `json:"price"`
	Status     string    `json:"status"`
	Expiration time.Time `json:"expiration"`
}

var (
	OrderBook []Order
	mutex     sync.RWMutex
)

func validateOrder(order Order) error {
	// Perform validation based on your requirements
	// For example, check if NumBonds and PricePerBond are within specified ranges
	// You can also implement additional validation logic

	if order.NumBonds < 1 || order.NumBonds > 10000 {
		return fmt.Errorf("Number of bonds must be between 1 and 10,000")
	}

	if order.Price < 0.0000 || order.Price > 100000000.0000 {
		return fmt.Errorf("Price per bond must be between 0.0000 and 100,000,000.0000")
	}

	return nil
}

// createOrder handles the creation of a new limit order
func CreateOrder(c *gin.Context) {

	UserId := c.GetInt("userID")

	fmt.Println("id user", UserId)
	// Parse request body to get order details
	var newOrder Order
	if err := c.ShouldBindJSON(&newOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Validate order details
	if err := validateOrder(newOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set expiration time (one hour from now)
	newOrder.Expiration = time.Now().Add(time.Hour)

	// Use a mutex to ensure atomic access to the OrderBook
	mutex.Lock()
	OrderBook = append(OrderBook, newOrder)
	mutex.Unlock()

	respuesta := models.CreateOrder(UserId, newOrder.BondId, newOrder.OrderType, newOrder.NumBonds, newOrder.Price, newOrder.Status, newOrder.Expiration)

	if respuesta == false {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	// Respond with the created order
	c.JSON(200, gin.H{"messages": "Successful registration"})
}

func UpdateOrder(c *gin.Context) {
	var (
		request Order
	)

	if err := c.ShouldBindJSON(&request); err != nil {
		BadRequest(c, []string{err.Error()})
		return
	}
	// Validate order details
	if err := validateOrder(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	respuesta := models.UpdateOrder(request.UserId, request.BondId, request.OrderType, request.NumBonds, request.Price, request.Status, idInt)

	//returns responses depending on the response from the record
	if respuesta == true {
		c.JSON(200, gin.H{"messages": "Registro se actualizo correctamente"})
	} else {
		c.JSON(500, gin.H{"messages": "Error al modificar los datos"})
	}

}

func GetUserOrders(c *gin.Context) {
	mutex.RLock()
	defer mutex.RUnlock()

	UserID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}

	respuesta, err := models.GetUserOrders(UserID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user orders"})
		return
	}
	if err != nil {
		log.Printf("Error getting user orders: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user orders"})
		return
	}

	log.Printf("User orders for UserID %d: %v", UserID, respuesta)

	c.JSON(200, gin.H{"messages": respuesta})

}

func GetOrderBook(c *gin.Context) {

	// Use RLock for concurrent read access
	mutex.RLock()
	defer mutex.RUnlock()

	respuesta, err := models.GetOrderBook(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user orders"})
		return
	}

	c.JSON(200, gin.H{"OrderBook": respuesta})
}

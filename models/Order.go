package models

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// "github.com/gin-gonic/gin"
type UserInfo struct {
	Id    int    `json:"id"`
	User  string `json:"user"`
	Email string `json:"email"`
	Phone int    `json:"phone"`
}

type BondInfo struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	CurrentPrice int    `json:"current_price"`
}

type Order struct {
	ID         int        `json:"id"`
	UserId     int        `json:"user_id"`
	BondId     int        `json:"bond_id"`
	OrderType  string     `json:"order_type"`
	NumBonds   int        `json:"num_bonds"`
	Price      float64    `json:"price"`
	Status     string     `json:"status"`
	Expiration time.Time  `json:"expiration"`
	User       []UserInfo `json:"user"`
	Bond       []BondInfo `json:"bond"`
}

// Modelo de refistro de usuario

func CreateOrder(UserId int, BondId int, OrderType string, NumBonds int, Price float64, Status string, Expiration time.Time) bool {

	createContactSQL := ("insert into test.orders(user_id, bond_id, order_type, num_bonds, price, status, expiration) values (?,?,?,?,?,?,?)")

	respuesta, err := database.db.Query(createContactSQL, UserId, BondId, OrderType, NumBonds, Price, Status, Expiration)
	fmt.Println(respuesta)
	if err != nil {
		log.Println(err)
		return false

	} else {
		return true
	}

}

// GetUserOrders retrieves orders along with user information for a specific user ID
func GetUserOrders(UserID int) ([]Order, error) {
	var (
		orders []Order
		user   UserInfo
		bond   BondInfo
	)

	getCatalogoSQL := "call UpdateAndShowOrders(?)"

	log.Printf("Executing SQL query: %s\n", getCatalogoSQL)
	rows, err := database.db.Query(getCatalogoSQL, UserID)
	if err != nil {
		log.Printf("Error querying database: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var expirationStr string
		var order Order

		err := rows.Scan(
			&order.ID, &order.UserId, &order.BondId, &order.OrderType, &order.NumBonds, &order.Price, &order.Status,
			&expirationStr, &user.Id, &user.User, &user.Email, &user.Phone, &bond.Id, &bond.Name, &bond.Description, &bond.CurrentPrice)
		if err != nil {
			log.Printf("Error scanning row: %v\n", err)
			return nil, err
		}

		// Parse the expiration string into a time.Time value
		expirationTime, err := time.Parse("2006-01-02 15:04:05", expirationStr)
		if err != nil {
			log.Printf("Error parsing expiration time: %v\n", err)
			return nil, err
		}
		order.Expiration = expirationTime

		// Initialize the Users slice and append the current user
		order.User = append(order.User, user)
		order.Bond = append(order.Bond, bond)

		// Append the order to the orders slice
		orders = append(orders, order)
	}

	return orders, nil
}

func GetOrderBook(c *gin.Context) ([]Order, error) {
	var (
		orders []Order
		user   UserInfo
		bond   BondInfo
	)

	getCatalogoSQL := "call UpdateAndShowOrders(NULL)"

	log.Printf("Executing SQL query: %s\n", getCatalogoSQL)
	rows, err := database.db.Query(getCatalogoSQL)
	if err != nil {
		log.Printf("Error querying database: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var expirationStr string
		var order Order

		err := rows.Scan(
			&order.ID, &order.UserId, &order.BondId, &order.OrderType, &order.NumBonds, &order.Price, &order.Status,
			&expirationStr, &user.Id, &user.User, &user.Email, &user.Phone, &bond.Id, &bond.Name, &bond.Description, &bond.CurrentPrice)
		if err != nil {
			log.Printf("Error scanning row: %v\n", err)
			return nil, err
		}

		// Parse the expiration string into a time.Time value
		expirationTime, err := time.Parse("2006-01-02 15:04:05", expirationStr)
		if err != nil {
			log.Printf("Error parsing expiration time: %v\n", err)
			return nil, err
		}
		order.Expiration = expirationTime

		// Initialize the Users slice and append the current user
		order.User = append(order.User, user)
		order.Bond = append(order.Bond, bond)

		// Append the order to the orders slice
		orders = append(orders, order)
	}

	return orders, nil
}

func UpdateOrder(UserId int, BondId int, OrderType string, NumBonds int, Price float64, Status string, Id int) bool {

	createContactSQL := ("UPDATE test.orders SET  user_id=?, bond_id=?, order_type=?, num_bonds=?, price=?, status=? where id=?")

	_, err := database.db.Query(createContactSQL, UserId, BondId, OrderType, NumBonds, Price, Status, Id)

	if err != nil {
		log.Println(err)
		return false

	} else {
		return true
	}

}

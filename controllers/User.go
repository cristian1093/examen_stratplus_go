package controllers

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/models"
)

// User Structure
type User struct {
	User     string `json:"user"`
	Email    string `json:"email"`
	Phone    int    `json:"phone"`
	Password string `json:"password"`
}

// validationRequest
func (request *User) validationRequest() (messages []string, err error) {

	if len(request.User) < 1 {
		messages = append(messages, "The Name is required")
		err = fmt.Errorf("The Name is required")
	}

	if request.Phone < 1 {
		messages = append(messages, "The Phone is required")
		err = fmt.Errorf("The Phone is required")
	}

	if len(request.Email) < 1 {
		messages = append(messages, "The Email is required")
		err = fmt.Errorf("The Email is required")
	}

	if len(request.Password) < 1 {
		messages = append(messages, "The Password is required")
		err = fmt.Errorf("The Password is required")
	}
	return
}

func CreateUser(c *gin.Context) {
	var request User

	if err := c.ShouldBindJSON(&request); err != nil {
		BadRequest(c, []string{err.Error()})
		return
	}

	if messages, err := request.validationRequest(); err != nil {
		BadRequest(c, messages)
		return
	}

	// Insert the data into the users table with the hashed password
	response := models.UserAdd(request.User, request.Email, request.Phone, request.Password)

	if response {
		c.JSON(200, gin.H{"message": "Successful registration"})

	} else {
		log.Println("Error inserting data into the database")
		c.JSON(200, gin.H{"message": "Error inserting data"})

	}
}

func ShowUser(c *gin.Context) {

	//insert the data into the users table
	response := models.Show_Users(c)

	//returns responses depending on the response from the record
	if response != nil {
		c.JSON(200, gin.H{"users": response})
	} else {
		c.JSON(500, gin.H{"users": "There is no data"})
	}

}

func UpdateUser(c *gin.Context) {
	var (
		request User
	)

	if err := c.ShouldBindJSON(&request); err != nil {
		BadRequest(c, []string{err.Error()})
		return
	}
	if messages, err := request.validationRequest(); err != nil {
		BadRequest(c, messages)
		return

	}

	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	respuesta := models.UserUpdate(request.User, request.Email, request.Phone, request.Password, idInt)

	//returns responses depending on the response from the record

	if respuesta == true {
		c.JSON(200, gin.H{"messages": "Registry was updated correctly"})
	} else {
		c.JSON(500, gin.H{"messages": "Error modifying data"})
	}

}

func DeleteUser(c *gin.Context) {

	idDelete := c.Param("id")
	idDeleteInt, _ := strconv.Atoi(idDelete)
	response := models.DeleteUser(idDeleteInt)

	//returns responses depending on the response from the record
	if response == true {
		c.JSON(200, gin.H{"messages": "User deleted successfully"})
	} else {
		c.JSON(500, gin.H{"messages": "Error deleting user"})
	}

}

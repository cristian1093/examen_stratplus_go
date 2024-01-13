package controllers

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/models"
)

// User Structure
type Bond struct {
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	CurrentPrice float64 `json:"current_price"`
}

// validationRequest
func (request *Bond) validationRequest() (messages []string, err error) {

	if len(request.Name) < 1 {
		messages = append(messages, "The Name is required")
		err = fmt.Errorf("The Name is required")
	}

	if len(request.Description) < 1 {
		messages = append(messages, "The Description is required")
		err = fmt.Errorf("The Description is required")
	}

	if request.CurrentPrice < 1 {
		messages = append(messages, "The CurrentPrice is required")
		err = fmt.Errorf("The CurrentPrice is required")
	}

	return
}

func CreateBond(c *gin.Context) {
	var request Bond

	if err := c.ShouldBindJSON(&request); err != nil {
		BadRequest(c, []string{err.Error()})
		return
	}

	if messages, err := request.validationRequest(); err != nil {
		BadRequest(c, messages)
		return
	}

	// Insert the data into the users table with the hashed password
	response := models.BondAdd(request.Name, request.Description, request.CurrentPrice)

	if response {
		c.JSON(200, gin.H{"message": "Successful registration"})

	} else {
		log.Println("Error inserting data into the database")
		c.JSON(200, gin.H{"message": "Error inserting data"})

	}
}

func ShowBond(c *gin.Context) {

	//insert the data into the users table
	response := models.Show_Bonds(c)

	//returns responses depending on the response from the record
	if response != nil {
		c.JSON(200, gin.H{"users": response})
	} else {
		c.JSON(500, gin.H{"users": "There is no data"})
	}

}

func UpdateBond(c *gin.Context) {
	var (
		request Bond
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
	respuesta := models.UpdateBond(request.Name, request.Description, request.CurrentPrice, idInt)

	//returns responses depending on the response from the record

	if respuesta == true {
		c.JSON(200, gin.H{"messages": "Registry was updated correctly"})
	} else {
		c.JSON(500, gin.H{"messages": "Error modifying data"})
	}

}

func DeleteBond(c *gin.Context) {

	idDelete := c.Param("id")
	idDeleteInt, _ := strconv.Atoi(idDelete)
	response := models.DeleteBond(idDeleteInt)

	//returns responses depending on the response from the record
	if response == true {
		c.JSON(200, gin.H{"messages": "User deleted successfully"})
	} else {
		c.JSON(500, gin.H{"messages": "Error deleting user"})
	}

}

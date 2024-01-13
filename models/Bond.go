package models

import (
	"log"

	"github.com/gin-gonic/gin"
)

// "github.com/gin-gonic/gin"

type Bond struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	CurrentPrice float64 `json:"current_price"`
}

// User registration model
func BondAdd(name string, description string, current_price float64) bool {

	createUserSQL := ("insert into test.bonds(name, description, current_price) values (?,?,?)")

	_, err := database.db.Query(createUserSQL, name, description, current_price)

	if err != nil {
		log.Println(err)
		return false

	} else {
		return true
	}

}

// Show User registration model
func Show_Bonds(c *gin.Context) []Bond {
	var (
		bond  Bond
		bonds []Bond
	)

	getUsersSQL := "select id, name,description,current_price from test.bonds"

	rows, err := database.db.Query(getUsersSQL)

	if err != nil {
		log.Println(err)
		panic(err)
	}

	for rows.Next() {
		err := rows.Scan(&bond.Id, &bond.Name, &bond.Description, &bond.CurrentPrice)
		bonds = append(bonds, bond)
		if err != nil {
			log.Println(err)
		}
	}
	defer rows.Close()

	return bonds

}

// User registration model
func UpdateBond(name string, description string, current_price float64, id int) bool {

	updateBondSQL := ("UPDATE test.bonds SET  name=?, description=?, current_price=? where id=?")

	_, err := database.db.Query(updateBondSQL, name, description, current_price, id)

	if err != nil {
		log.Println(err)
		return false
	}

	return true

}

func DeleteBond(id int) bool {

	success := false
	deleteBondSQL := "DELETE FROM test.bonds WHERE id = ?"
	_, err := database.db.Query(deleteBondSQL, id)
	if err != nil {
		log.Println(err)
		panic(err)
	} else {

		success = true
	}

	return success

}

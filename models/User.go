package models

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

// "github.com/gin-gonic/gin"

type User struct {
	Id    int    `json:"id"`
	User  string `json:"user"`
	Email string `json:"email"`
	Phone int    `json:"phone"`
}

// User registration model
func UserAdd(user string, email string, phone int, password string) bool {

	createUserSQL := ("insert into test.users(user, email, phone, password) values (?,?,?,?)")

	_, err := database.db.Query(createUserSQL, user, email, phone, password)

	if err != nil {
		log.Println(err)
		return false

	} else {
		return true
	}

}

// Show User registration model
func Show_Users(c *gin.Context) []User {
	var (
		user  User
		users []User
	)

	getUsersSQL := "select id, user,email,phone from test.users"

	rows, err := database.db.Query(getUsersSQL)

	if err != nil {
		log.Println(err)
		panic(err)
	}

	for rows.Next() {
		err := rows.Scan(&user.Id, &user.User, &user.Email, &user.Phone)
		users = append(users, user)
		if err != nil {
			log.Println(err)
		}
	}
	defer rows.Close()

	return users

}

// User registration model
func UserUpdate(user string, email string, phone int, password string, id int) bool {

	updateUserSQL := ("UPDATE test.users SET  user=?, email=?, phone=?,password=? where id=?")

	respuesta, err := database.db.Query(updateUserSQL, user, email, phone, password, id)
	fmt.Println(respuesta)
	if err != nil {
		log.Println(err)
		return false
	}

	return true

}

func DeleteUser(id int) bool {

	success := false
	deleteUserSQL := "DELETE FROM test.users WHERE id = ?"
	_, err := database.db.Query(deleteUserSQL, id)
	fmt.Println(database.db.Query(deleteUserSQL, id))
	if err != nil {
		log.Println(err)
		panic(err)
	} else {

		success = true
	}

	return success

}

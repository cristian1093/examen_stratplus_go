package models

import (
	"log"
)

type UserData struct {
	Id    int    `json:"id"`
	User  string `json:"user"`
	Email string `json:"email"`
	Phone int    `json:"phone"`
}

// Login Model
func LoginUser(user string, password string) string {
	var result string
	createContactSQL := ("select count(*) from test.users where (user=? or email=?) and (password=?)")

	err := database.db.QueryRow(createContactSQL, user, user, password).Scan(&result)
	if err != nil {
		log.Println(err)
	}
	return result
}

// InfoUser Model get the user's data to add it to the token, for example the id
func InfoUser(user_name string, password string) []UserData {
	var (
		user  UserData
		users []UserData
	)

	getCatalogSQL := "select id, user, email, phone  from test.users where (user=? or email=?) and (password=?)"
	log.Println(getCatalogSQL)
	rows, err := database.db.Query(getCatalogSQL, user_name, user_name, password)

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

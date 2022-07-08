package models

import (
	"log"
)

//Modelo de Login
func LoginUser(usuario string, contrasena string) string {
	var result string
	createContactSQL := ("select count(*) from examen_stratplus.usuarios where (correo=? or usuario=?) and (contrasena=?)")

	err := database.db.QueryRow(createContactSQL, usuario, usuario, contrasena).Scan(&result)
	if err != nil {
		log.Println(err)
	}
	return result
}

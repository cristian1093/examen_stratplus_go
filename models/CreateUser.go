package models

import (
	"fmt"
	"log"
)

// "github.com/gin-gonic/gin"

// Modelo de refistro de usuario
func UserAdd(usuario string, correo string, telefono int, contrasena string) bool {

	createContactSQL := ("insert into examen_stratplus.usuarios (usuario, correo, telefono, contrasena) values (?,?,?,?)")

	respuesta, err := database.db.Query(createContactSQL, usuario, correo, telefono, contrasena)
	fmt.Println(respuesta)
	if err != nil {
		log.Println(err)
		return false

	} else {
		return true
	}

}

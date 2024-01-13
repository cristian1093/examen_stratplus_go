package models

import "log"

// Modelo que valida de correo o teléfono ya se encuentran registrados
func Validate(correo string, telefono int) string {
	var result string
	createContactSQL := ("select count(*) from test.usuarios where correo=? or telefono=?")

	err := database.db.QueryRow(createContactSQL, correo, telefono).Scan(&result)
	if err != nil {
		log.Println(err)
	}
	return result
}

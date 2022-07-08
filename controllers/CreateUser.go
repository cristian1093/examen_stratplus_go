package controllers

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/models"
)

//Estructura de Usuario
type User struct {
	Usuario    string `json:"usuario"`
	Correo     string `json:"correo"`
	Telefono   int    `json:"telefono"`
	Contrasena string `json:"contraseña"`
}

//validationRequest
func (request *User) validationRequest() (messages []string, err error) {

	//valida que el usuario sea requerido
	if len(request.Usuario) < 1 {
		messages = append(messages, "El usuario es requerido")
		err = fmt.Errorf("El usuario es requerido")
	}

	//valida que el teléfono sea requerido
	if request.Telefono < 1 {
		messages = append(messages, "El teléfono es requerido")
		err = fmt.Errorf("El teléfono es requerido")

	} else {

		// valida que el télefono  cuente con una estructra de 10 dígitos
		if isTelefono, _ := regexp.MatchString(`([0-9]*).{10,10}$`, strconv.Itoa(request.Telefono)); !isTelefono {
			messages = append(messages, "El teléfono no cuenta con un formato correcto")
			err = fmt.Errorf("El teléfono no cuenta con un formato correcto")
		}
	}

	//valida que el correo sea requerido
	if len(request.Correo) < 1 {
		messages = append(messages, "El correo es requerido")
		err = fmt.Errorf("El correo es requerido")

	} else {
		//valida que contenga el formato de correo
		if isCorreo, _ := regexp.MatchString(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`, request.Correo); !isCorreo {
			messages = append(messages, "El correo no cuenta con un formato correcto")
			err = fmt.Errorf("El correo no cuenta con un formato correcto")
		}
	}

	//valida que la contraseña sea requerida
	if len(request.Contrasena) < 1 {
		messages = append(messages, "La contraseña es requerida")
		err = fmt.Errorf("La contraseña es requerida")

	} else {

		//valida que la contraseña debe ser mayor a 6 caracteres
		if len(request.Contrasena) < 6 {
			messages = append(messages, "La contraseña no debe ser mayor a 6 caracteres")
			err = fmt.Errorf("La contraseña no debe ser mayor a 6 caracteres")
		}

		//valida que la contraseña debe ser menor a 12 caracteres
		if len(request.Contrasena) > 12 {
			messages = append(messages, "La contraseña no debe ser menor a 12 caracteres")
			err = fmt.Errorf("La contraseña no debe ser menor a 12 caracteres")
		}

		//valida que la contraseña debe contener al menos una mayúscula
		if isContrasena, _ := regexp.MatchString(`[A-Z]`, request.Contrasena); !isContrasena {
			messages = append(messages, "La contraseña debe contener al menos una mayúscula")
			err = fmt.Errorf("La contraseña debe contener al menos una mayúscula")
		}

		//valida que la contraseña debe contener al menos una minúscula
		if isContrasena, _ := regexp.MatchString(`[a-z]`, request.Contrasena); !isContrasena {
			messages = append(messages, "La contraseña debe contener al menos una minúscula")
			err = fmt.Errorf("La contraseña debe contener al menos una minúscula")
		}

		//valida que la contraseña debe contener al menos un número
		if isContrasena, _ := regexp.MatchString(`[0-9]`, request.Contrasena); !isContrasena {
			messages = append(messages, "La contraseña debe contener al menos un número")
			err = fmt.Errorf("La contraseña debe contener al menos un número")
		}

		//valida que la contraseña debe contener al menos un caracter especial
		if isContrasena, _ := regexp.MatchString(`[@$!%*?&]`, request.Contrasena); !isContrasena {
			messages = append(messages, "La contraseña debe contener al menos un caracter especial")
			err = fmt.Errorf("La contraseña debe contener al menos un caracter especial")
		}

	}

	return
}

func CreateUser(c *gin.Context) {
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

	//hace una llamada al modelo, que valida si el correo o el teléfono se encuentran registrados
	result := models.Validate(request.Correo, request.Telefono)

	result_int, _ := strconv.Atoi(result)

	//si el resultado de la validacion encuentra un orreo o el teléfono registrado retorna un mensaje de error
	if result_int >= 1 {
		c.JSON(400, gin.H{"messages": "La contraseña/teléfono ya se encuentra registrado"})
	} else {

		//inserta los datos en la tabla usuarios
		respuesta := models.UserAdd(request.Usuario, request.Correo, request.Telefono, request.Contrasena)

		//regresa respuestas dependiendo de la respuesta de el registro
		if respuesta == true {
			c.JSON(200, gin.H{"messages": "Registro correcto"})
		} else {
			c.JSON(500, gin.H{"messages": "Error al insertar los datos"})
		}
	}

}

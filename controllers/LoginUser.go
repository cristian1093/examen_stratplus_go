package controllers

import (
	"fmt"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/models"
)

type LoginUserData struct {
	Usuario    string `json:"usuario"`
	Contrasena string `json:"contraseña"`
}

var mySigningKey = []byte("mysupersecret")

//Función JWT

func GenerateJWT(usuario string) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = usuario
	claims["exp"] = time.Now().Add(time.Minute * 30)

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		fmt.Errorf("Something went wrong: %s", err.Error())
		return "", err

	}

	return tokenString, nil

}

//valida el request del login
func (request *LoginUserData) ValidationLoginRequest() (messages []string, err error) {

	//valida que el usuario es requerido
	if len(request.Usuario) < 1 {
		messages = append(messages, "El campo usuario es requerido")
		err = fmt.Errorf("El campo usuario es requerido")

	}

	//valida que la contraseña es requerida
	if len(request.Contrasena) < 1 {
		messages = append(messages, "El campo contraseña es requerido")
		err = fmt.Errorf("El campo contraseña es requerido")

	}
	return
}

//Funcion general del login
func LoginUser(c *gin.Context) {

	var (
		request LoginUserData
	)

	if err := c.ShouldBindJSON(&request); err != nil {
		BadRequest(c, []string{err.Error()})
		return
	}
	if messages, err := request.ValidationLoginRequest(); err != nil {
		BadRequest(c, messages)
		return
	}

	// valida en que el login del usuario
	response := models.LoginUser(request.Usuario, request.Contrasena)

	response_int, _ := strconv.Atoi(response)

	// retorna un error si el usuario o contraseña son incorrectos
	if response_int == 0 {
		c.JSON(401, gin.H{"messages": "usuario / contraseña incorrecto"})
	} else {

		//se invoca la funcion de jwt para retornar un token
		tokenString, err := GenerateJWT(request.Usuario)

		if err != nil {
			fmt.Println("Errror generating token string")
		}

		// retorna el token del login
		c.JSON(200, gin.H{"token": tokenString})
	}

}

package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/models"
)

type LoginUserData struct {
	Id       string `json:"id"`
	User     string `json:"user"`
	Password string `json:"password"`
}

var mySigningKey = []byte("mysupersecret")

//JWT function

func GenerateJWT(user string, id int) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = user
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Minute * 30)

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		fmt.Errorf("Something went wrong: %s", err.Error())
		return "", err

	}

	return tokenString, nil

}

func TokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Check for the "Bearer " prefix and strip it
		if len(tokenString) >= 7 && strings.ToUpper(tokenString[0:7]) == "BEARER " {
			tokenString = tokenString[7:]
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return mySigningKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Extract user ID from the token claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		userID, ok := claims["id"].(float64)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Set the user ID in the Gin context
		c.Set("userID", int(userID))

		c.Next()
	}
}

// validate the login request
func (request *LoginUserData) ValidationLoginRequest() (messages []string, err error) {

	//validates that the user is required
	if len(request.User) < 1 {
		messages = append(messages, "The user field is required")
		err = fmt.Errorf("The user field is required")

	}

	//validates that the password is required
	if len(request.Password) < 1 {
		messages = append(messages, "Password field is required")
		err = fmt.Errorf("Password field is required")

	}
	return
}

// General login function
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

	// validates the user's login
	response := models.LoginUser(request.User, request.Password)

	response_int, _ := strconv.Atoi(response)

	// returns an error if the username or password is incorrect
	if response_int == 0 {
		c.JSON(401, gin.H{"messages": "user / contraseña incorrecto"})
	} else {
		response := models.InfoUser(request.User, request.Password)

		//the jwt function is invoked to return a token
		tokenString, err := GenerateJWT(request.User, response[0].Id)

		if err != nil {
			fmt.Println("Errror generating token string")
		}

		// retorna el token del login
		c.JSON(200, gin.H{"token": tokenString})
	}

}

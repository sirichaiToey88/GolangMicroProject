package login

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"project/models/users"
	"project/schemas"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func generateSigningKey(length int) (string, error) {
	key := make([]byte, length)
	_, err := rand.Read(key)
	if err != nil {
		return "", err
	}

	// Convert the random bytes to a base64-encoded string
	encodedKey := base64.URLEncoding.EncodeToString(key)

	return encodedKey, nil
}

func LoginHandle(c echo.Context) error {
	// signingKey, err := generateSigningKey(32)
	// dataUser := []users.User_login{}
	body := &users.User_login{}
	if err := c.Bind(body); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request payload")
	}

	fmt.Println("Email =", body.Email)
	fmt.Println("Mobile =", body.Mobile)

	db, err := schemas.Connectdb()
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query(`SELECT id, user_name, email,mobile,User_id FROM t_user_template WHERE email=? AND mobile=? AND del=0`, body.Email, body.Mobile)

	dataUser := []users.User_login{}
	for rows.Next() {
		var m users.User_login
		if err := rows.Scan(&m.Id, &m.User_name, &m.Email, &m.Mobile, &m.User_id); err != nil {
			log.Fatal("error select into DB", err)
		}

		dataUser = append(dataUser, m)

		fmt.Println("Data Response=>", dataUser)
		// Create JWT token
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["email"] = m.Email
		claims["exp"] = time.Now().Add(time.Minute * 5).Unix() // Token expiration time

		// Sign the token
		tokenString, err := token.SignedString([]byte("Sirichai"))
		if err != nil {
			return echo.ErrInternalServerError
		}
		// Return the token as JSON response
		// response := map[string]string{"token": tokenString, "user": jsonArray}
		response := map[string]interface{}{
			"token": tokenString,
			"data":  dataUser,
		}

		return c.JSON(http.StatusOK, response)
	}
	return c.JSON(http.StatusBadRequest, err.Error())
}

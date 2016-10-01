/*
https://godoc.org/github.com/dgrijalva/jwt-go#example-New--Hmac
https://dinosaurscode.xyz/go/2016/06/17/golang-jwt-authentication/

 */
package main

import (
	"fmt"
	"time"
	jwt "github.com/dgrijalva/jwt-go"
)

func main() {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	hmacSampleSecret := []byte("secret string")
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(hmacSampleSecret)

	fmt.Println(tokenString, err)
}

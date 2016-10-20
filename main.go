package main

import (
	"github.com/auth0/go-jwt-middleware"
	"github.com/codegangsta/negroni"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"net/http"
	"time"
	"github.com/Zhanat87/stack-auth/controllers"
	"github.com/Zhanat87/stack-auth/common"
)

func main() {

	StartServer()

}

func StartServer() {
	r := mux.NewRouter()

	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return mySigningKey, nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})

	r.HandleFunc("/ping", controllers.PingHandler)
	r.Handle("/secured/ping", negroni.New(
		negroni.HandlerFunc(jwtMiddleware.HandlerWithNext),
		negroni.Wrap(http.HandlerFunc(controllers.SecuredPingHandler)),
	))

	r.Handle("/get-token", GetTokenHandler).Methods("GET")

	http.Handle("/", r)
	http.ListenAndServe(":8082", nil)
}


// Глобальный секретный ключ
var mySigningKey = []byte("secret")

var GetTokenHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){

	// Создаем новый токен
	// Устанавливаем набор параметров для токена
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"admin": "true",
		"name": "username",
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	// Подписываем токен нашим секретным ключем
	tokenString, _ := token.SignedString(mySigningKey)

	// Отдаем токен клиенту
	common.respondJson(tokenString, w)
})
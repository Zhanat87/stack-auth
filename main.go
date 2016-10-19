package main

import (
	"encoding/json"
	"github.com/auth0/go-jwt-middleware"
	"github.com/codegangsta/negroni"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"net/http"
	"time"
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

	r.HandleFunc("/ping", PingHandler)
	r.Handle("/secured/ping", negroni.New(
		negroni.HandlerFunc(jwtMiddleware.HandlerWithNext),
		negroni.Wrap(http.HandlerFunc(SecuredPingHandler)),
	))

	r.Handle("/get-token", GetTokenHandler).Methods("GET")

	http.Handle("/", r)
	http.ListenAndServe(":8082", nil)
}

type Response struct {
	Text string `json:"text"`
}

func respondJson(text string, w http.ResponseWriter) {
	response := Response{text}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
	respondJson("All good. You don't need to be authenticated to call this", w)
}

func SecuredPingHandler(w http.ResponseWriter, r *http.Request) {
	respondJson("All good. You only get this message if you're authenticated", w)
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
	respondJson(tokenString, w)
})
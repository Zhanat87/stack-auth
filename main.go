package main

import (
	//"encoding/json"
	"fmt"
	//"net/http"
	"os"
	//"strings"
	"net/http"
	//"encoding/json"
)

//type EnvVarsModel struct {
//	vars []string `json:"vars"`
//	mongoDbHost string `json:"mongoDbHost"`
//	a string `json:"a"`
//}

func main() {
	/*
	@link https://gobyexample.com/environment-variables
	 */
	//fmt.Println("Something went wrong!", os.Getenv("MYVAR"))
	//fmt.Println()
	//for _, e := range os.Environ() {
	//	fmt.Println(e)
	//	//pair := strings.Split(e, "=")
	//	//fmt.Println(pair[0])
	//	//fmt.Println(pair[1])
	//}
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		//output, err := json.Marshal(EnvVarsModel{vars: os.Environ(), mongoDbHost: os.Getenv("STACK_MONGODB_PORT_27017_TCP_ADDR"), a: "b"})
		//if err != nil {
		//	fmt.Println("Something went wrong!")
		//}
		//w.Header().Set("Content-Type", "application/json")
		//w.Write(output)
		for _, e := range os.Environ() {
			fmt.Println(e)
			//pair := strings.Split(e, "=")
			//fmt.Println(pair[0])
			//fmt.Println(pair[1])
		}

	})
	http.ListenAndServe(":8082", nil)
}
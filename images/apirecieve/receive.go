package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var u User
		json.NewDecoder(r.Body).Decode(&u)
		fmt.Printf("%+v\n", u)
	})

	http.ListenAndServe(":8080", nil)
}

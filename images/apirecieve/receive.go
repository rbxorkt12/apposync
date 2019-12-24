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
	fmt.Printf("Receieve server start!\n")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var u User
		json.NewDecoder(r.Body).Decode(&u)
		fmt.Printf("%+v\n", u)
		fmt.Printf("%s:%s", u.Id, u.Name)
	})

	http.ListenAndServe(":10080", nil)
}

package main

import (
	"encoding/json"
	"fmt"
)

type User struct{
	UserName string `json:"username"`
	Email []string `json:"email"`
}

func main(){
	user := &User{UserName: "blackhawk", Email: []string{"blackhawk@mail.com","abc@mail.com"}}

	j, _ := json.Marshal(user)

	fmt.Println(string(j))
}
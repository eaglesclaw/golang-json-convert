package main

import (
	"encoding/json"
	"fmt"
)

func main(){
	myJson := `{"person1": 45, "person2": 89}`

	byteToJson := []byte(myJson)

	var myTeam map[string]int

	err := json.Unmarshal(byteToJson, &myTeam)
	if err != nil { fmt.Println(err) }
	fmt.Println(myTeam)
}
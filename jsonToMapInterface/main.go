package main

import (
	"encoding/json"
	"fmt"
)

func main(){

	myJson := `{"birinci":1, "ikinci":["abc","def"]}`

	var myMap map[string]interface{}

	err := json.Unmarshal([]byte(myJson), &myMap)
	if err != nil { fmt.Println(err)}

	fmt.Println(myMap["ikinci"].([]interface{})[0])
}
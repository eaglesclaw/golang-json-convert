package main

import (
	"encoding/json"
	"fmt"
)

func main(){
	abc := map[string]int{"hasan":12, "mehmet":19}

	getValue, _ := json.Marshal(abc)

	fmt.Println(getValue)
}
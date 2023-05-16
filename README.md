# golang-json-convert

	a := `{"person1":23, "person2":76}`

	o := []byte(a)

	var data map[string]int

	err := json.Unmarshal(o, &data)

	if err != nil{
		fmt.Println(err)
	}
	fmt.Printf("Value: %v\nType: %T",data,data)

	a:= map[string]int{"person1":55,"person2":54}

	j,_:= json.Marshal(a)

	fmt.Println(string(j))

    --

    a := `{"person1":23, "person2":["56","55"]}`
	o := []byte(a)

	var data map[string]interface{}

	err := json.Unmarshal(o, &data)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)

	fmt.Println(data["person2"].([]interface{})[0])
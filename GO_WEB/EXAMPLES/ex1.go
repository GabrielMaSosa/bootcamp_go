package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	Name      string
	Price     float64
	Published string
}

func main() {
	var p2 Product
	jsndata := `{"Name": "Mac Book2","Price": 1502.0, "Published": "2022"}`
	p := Product{
		Name:      "Mac Book",
		Price:     1500.0,
		Published: "2023",
	}

	jsondata, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsondata))

	if err2 := json.Unmarshal([]byte(jsndata), &p2); err2 != nil {
		panic(err)

	}
	fmt.Println(p)
}

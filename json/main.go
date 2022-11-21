package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Expense struct {
	Name      string  `json:"name"`
	Value     float64 `json:"value"`
	Type      string  `json:"type"`
	Necessity string  `json:"necessity"`
}

func main() {
	expense := Expense{
		Name:      "Gasoline",
		Value:     150.00,
		Type:      "necessary",
		Necessity: "high",
	}

	res, err := json.Marshal(expense)
	if err != nil {
		println(err)
	}

	println(string(res))

	err = json.NewEncoder(os.Stdout).Encode(expense)
	if err != nil {
		println(err)
	}

	pureJson := []byte(`{"name": "Json to struct", "type": "test"}`)
	var expense2 Expense
	err = json.Unmarshal(pureJson, &expense2)
	if err != nil {
		println(err)
	}

	fmt.Printf("Name: %s, type: %s", expense2.Name, expense2.Type)
}

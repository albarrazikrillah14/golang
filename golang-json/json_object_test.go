package golang_json

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

type Customer struct {
	Firstname  string   `json:"firstname"`
	Middlename string   `json:"middlename"`
	Lastname   string   `json:"lastname"`
	Hobbies    []string `json:"hobbies"`
}

func ConvertJson(data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

func TestJsonObject(t *testing.T) {
	data := Customer{
		Firstname:  "Albarra",
		Middlename: "-",
		Lastname:   "Zikrillah",
		Hobbies:    []string{"Basket", "Membaca", "Ngoding"},
	}

	ConvertJson(data)
}

func TestDecodeJson(t *testing.T) {
	jsonRequest := `{"Firstname":"Albarra","Middlename":"-","Lastname":"Zikrillah"}`
	jsonBytes := []byte(jsonRequest)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer.Firstname)
}

func TestDecoder(t *testing.T) {
	reader, err := os.Open("Customer.json")
	if err != nil {
		panic(err)
	}

	decoder := json.NewDecoder(reader)
	customer := &Customer{}

	err = decoder.Decode(customer)

	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
}

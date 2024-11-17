package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ImageUrl string `json:"imageUrl"`
}

func TestJsonTag(t *testing.T) {
	product := Product{
		Id:       "P001",
		Name:     "Mac Book",
		ImageUrl: "http://localhost:8080/products/images/1",
	}

	byte, err := json.Marshal(product)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(byte))
}

func TestDecode(t *testing.T) {
	jsonString := `{"id":"P001","name":"Mac Book","imageUrl":"http://localhost:8080/products/images/1"}`
	jsonBytes := []byte(jsonString)

	product := &Product{}
	err := json.Unmarshal(jsonBytes, product)

	if err != nil {
		panic(err)
	}

	fmt.Println(product)
}

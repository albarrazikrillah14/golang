package main

import "medomeckz/auth-api/src/Infrastructures/http"

func main() {
	err := http.NewCreateServer("localhost", 3000)
	if err != nil {
		panic(err)
	}
}

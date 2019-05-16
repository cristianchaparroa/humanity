package main

import (
	"github.com/cristianchaparroa/humanity/backend/api"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	server := api.NewChatServer()
	server.Run()

	defer server.Close()
}

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
	server.SetupDB()
	server.SetupRoutes()
	server.Run()

	defer server.Close()
}

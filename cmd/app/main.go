package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/mhcodev/fake_store_api/internal/repository"
)

func main() {

	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbType := "postgres"

	var dbRepo *repository.DBRepository

	switch dbType {
	case "postgres":
		postgresRepo, conn := repository.InitPosgresRepositories()
		dbRepo = postgresRepo
		defer conn.Close()
	}

	users, err := dbRepo.UserRepository.GetUsersByParams()

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	for _, v := range users {
		fmt.Println(v.Name)
	}

}

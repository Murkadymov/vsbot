package commands

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GetToken() string {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Got error with env: ", err)
	}
	TOKEN := os.Getenv("TOKEN")

	return TOKEN
}

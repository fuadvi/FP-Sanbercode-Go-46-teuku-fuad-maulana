package utilitis

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func Getenv(key, fallback string) string {
	err := godotenv.Load()

	fmt.Println(err)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if value, ok := os.LookupEnv(key); ok {
		fmt.Println(value)
		return value
	}
	return fallback
}

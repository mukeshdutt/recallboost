package helpers

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func ConfigByKey(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err.Error())
	}
	value := os.Getenv(key)
	return value
}

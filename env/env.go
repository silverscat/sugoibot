package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv .env取得
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// Debug デバッグモード
func Debug() bool {
	LoadEnv()
	debugStr := os.Getenv("DEBUG")
	if debugStr == "true" {
		return true
	}
	return false
}

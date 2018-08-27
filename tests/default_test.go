package tests

import (
	"log"
	"goBoilterplate/config"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {

	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := config.Database()
	defer db.Close()

	config.Redis()

	os.Exit(m.Run())
}

package osugoi

import (
	"fmt"
	"testing"
	"os"

	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Warting: .env file not found or failed to load")
	}

	os.Exit(m.Run())
}

func TestNewClient(t *testing.T) {
	_, err := NewClient()
	if err != nil {
		t.Errorf("No errors are expected. Got %v", err)
	}
}

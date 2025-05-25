package osugoi

import (
	"fmt"
	"testing"
	"os"

	"github.com/joho/godotenv"
)

const baseURL = "https://osu.ppy.sh"

func TestMain(m *testing.M) {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Warting: .env file not found or failed to load")
	}

	os.Exit(m.Run())
}

func TestAuthenticate(t *testing.T) {
	client := NewClient(baseURL)

	if err := client.Authenticate(os.Getenv("CLIENT_ID"),
			os.Getenv("CLIENT_SECRET")); err != nil {
		t.Errorf("Authorize: unexpected error: %v", err)
	}
}

package osugoi

import (
	"fmt"
	"testing"
	"os"

	"github.com/joho/godotenv"
)

const baseURL = "https://osu.ppy.sh"
var testClient *Client

func TestMain(m *testing.M) {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Warting: .env file not found or failed to load")
	}

	testClient = NewClient(baseURL)

	err := testClient.Authenticate(
		os.Getenv("CLIENT_ID"),
		os.Getenv("CLIENT_SECRET"),
	)
	if err != nil {
		fmt.Errorf("Error creating client for testing: %w", err)
		os.Exit(1)
	}

	os.Exit(m.Run())
}

func TestAuthenticate(t *testing.T) {
	client := NewClient(baseURL)

	err := client.Authenticate(
		os.Getenv("CLIENT_ID"),
		os.Getenv("CLIENT_SECRET"),
	)
	if err != nil {
		t.Errorf("Authorize: unexpected error: %v", err)
	}
}

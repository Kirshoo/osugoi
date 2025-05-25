package osugoi

import (
	"testing"
	"os"
	"log"

	"github.com/joho/godotenv"
)

const baseURL = "https://osu.ppy.sh"
var testClient *Client

func TestMain(m *testing.M) {
	if err := godotenv.Load(); err != nil {
		log.Println("Warting: .env file not found or failed to load")
	}

	testClient = NewClient(baseURL)

	err := testClient.Authenticate(
		os.Getenv("CLIENT_ID"),
		os.Getenv("CLIENT_SECRET"),
	)
	if err != nil {
		log.Fatalf("Error creating client for testing: %v", err)
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

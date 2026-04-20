package customkeys_test

import (
	"fmt"
	"log"
	"os"

	customkeys "github.com/cksxe/sdk-go"
)

func Example() {
	client, err := customkeys.New(
		customkeys.WithToken(os.Getenv("CUSTOMKEYS_TOKEN")),
		customkeys.WithEnv(os.Getenv("CUSTOMKEYS_ENV_ID")),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// Fetch a single secret.
	dbURL, ok := client.Get("DATABASE_URL")
	if ok {
		fmt.Println("Database URL loaded")
		_ = dbURL
	}

	// Fetch with a default fallback.
	host := client.GetOrDefault("CACHE_HOST", "localhost:6379")
	_ = host

	// Fetch all secrets at once.
	all := client.GetAll()
	fmt.Printf("Loaded %d secrets\n", len(all))
}

func Example_mustGet() {
	client, err := customkeys.New(
		customkeys.WithToken(os.Getenv("CUSTOMKEYS_TOKEN")),
		customkeys.WithEnv(os.Getenv("CUSTOMKEYS_ENV_ID")),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// MustGet panics if the secret is missing — good for required config.
	_ = client.MustGet("DATABASE_URL")
	_ = client.MustGet("JWT_SECRET")
}

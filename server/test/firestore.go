package main

import (
	"cloud.google.com/go/firestore"
	"context"
	"log"
	"os"
)

func main() {
	ctx := context.Background()
	os.Setenv("FIRESTORE_EMULATOR_HOST", "localhost:8080")
	//client, err := firestore.NewClient(ctx, "dessage", option.WithCredentialsFile("dessage-c3b5c95267fb.json"))
	client, err := firestore.NewClientWithDatabase(ctx, "dessage", "dessage") // 使用适当的项目 ID
	if err != nil {
		log.Fatalf("Failed to create Firestore client: %v", err)
	}
	defer client.Close()

	_, err = client.Collection("twitter-user").Doc("Ada").Set(ctx, map[string]interface{}{
		"id":       "Ada",
		"name":     "Lovelace",
		"username": "1815",
	})
	if err != nil {
		log.Fatalf("Failed to add a new user: %v", err)
	}
}

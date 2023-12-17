package main

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()

	client, err := firestore.NewClient(ctx, "dessage", option.WithCredentialsFile("dessage-c3b5c95267fb.json"))
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

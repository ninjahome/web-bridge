package main

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"google.golang.org/api/option"
	"log"
	"os"
)

func main() {
	queryNoOfCollection("ninja-user")
}
func test() {
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

func updateField() {
	ctx := context.Background()
	saPath := "dessage-c3b5c95267fb.json"
	client, err := firestore.NewClientWithDatabase(ctx, "dessage",
		"dessage-release", option.WithCredentialsFile(saPath))
	if err != nil {
		log.Fatalf("Failed to create Firestore client: %v", err)
	}
	defer client.Close()

	// 假设有一个文档 "document-id" 在集合 "your-collection" 中
	docRef := client.Collection("tweets-posted").Doc("1714188650865")

	// 更新字段，添加回车换行
	newContent := "Hello World, now we are coming!\nLink the world, break down barriers, individualism means everything!" // 这里 "\n" 是换行符
	_, err = docRef.Update(ctx, []firestore.Update{{Path: "text", Value: newContent}})
	if err != nil {
		log.Fatalf("Failed to update document: %v", err)
	}

	log.Println("Document updated successfully with newline")
}

func queryNoOfCollection(table string) {
	ctx := context.Background()
	saPath := "dessage-c3b5c95267fb.json"
	client, err := firestore.NewClientWithDatabase(ctx, "dessage",
		"dessage-release", option.WithCredentialsFile(saPath))
	if err != nil {
		log.Fatalf("Failed to create Firestore client: %v", err)
	}
	defer client.Close()

	iter := client.Collection(table).Documents(ctx)
	defer iter.Stop()

	count := 0
	for {
		_, err := iter.Next()
		if err != nil {
			fmt.Println(err)
			break
		}
		count++
	}

	fmt.Println("total count:", count)
}

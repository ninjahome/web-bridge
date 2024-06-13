package main

import (
	"cloud.google.com/go/firestore"
	"context"
	"errors"
	"fmt"
	"github.com/ninjahome/web-bridge/database"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"os"
	"strings"
)

func main() {
	//queryNoOfCollection("ninja-user")
	err := refactorPoints("dessage", "dessage-alpha")
	if err != nil {
		panic(err)
	}
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

func refactorPoints(projectID, databaseID string) error {
	ctx := context.Background()
	saPath := "dessage-c3b5c95267fb.json"
	client, err := firestore.NewClientWithDatabase(ctx, projectID,
		databaseID, option.WithCredentialsFile(saPath))
	if err != nil {
		return err
	}
	defer client.Close()

	iter := client.Collection(database.DBTableNJUser).Documents(ctx)
	defer iter.Stop()

	for {
		doc, err := iter.Next()
		if err != nil {
			if errors.Is(err, iterator.Done) {
				break
			}
			return err
		}

		points := doc.Data()["points"].(int64)
		addr := strings.ToLower(doc.Data()["eth_addr"].(string))

		docRef := client.Collection(database.DBTableUserPoints).Doc(addr)
		_, err = docRef.Get(ctx)
		if err == nil {
			fmt.Println("no need updating for user:", addr)
			continue
		}

		if status.Code(err) != codes.NotFound {
			return err
		}

		sp := &database.SysPoints{
			EthAddr: addr,
			Points:  float32(points),
		}

		_, err = docRef.Set(ctx, sp)
		if err != nil {
			return err
		}

		fmt.Println("success update for user:", addr, points)
	}

	fmt.Println("process success")
	return nil
}

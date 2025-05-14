package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/bucketeer-io/go-server-sdk/pkg/bucketeer"
	"github.com/bucketeer-io/go-server-sdk/pkg/bucketeer/user"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	client, err := bucketeer.NewSDK(
		ctx,
		bucketeer.WithAPIKey("test"),
		bucketeer.WithHost("test"),
		bucketeer.WithTag("test"),
		bucketeer.WithEnableLocalEvaluation(true),          // <--- Enable the local evaluation
		bucketeer.WithCachePollingInterval(10*time.Minute), // <--- Change the default interval if needed
	)
	if err != nil {
		log.Fatalf("Failed initialize the new client: %v", err)
	}

	testUser := user.NewUser(
		"x1",
		nil,
	)
	showNewFeature := client.StringVariation(ctx, testUser, "fcm-feature", "value-2")
	if showNewFeature == "value-1" {
		fmt.Println("New feature is enabled")
	} else {
		fmt.Println("New feature is being tested")
	}

	showNewFeature2 := client.StringVariation(ctx, testUser, "test-211067", "value-2")
	if showNewFeature2 == "value-1" {
		fmt.Println("New feature is enabled")
	} else {
		fmt.Println("New feature is being tested")
	}
}

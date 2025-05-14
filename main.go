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
	showNewFeature := client.StringVariation(ctx, testUser, "feature-go-server-e2e-string", "value-2")
	fmt.Printf("Show new feature: %v\n", showNewFeature)

	if showNewFeature == "value-1" {
		fmt.Println("New feature is enabled")
	} else {
		fmt.Println("New feature is being tested")
	}
}

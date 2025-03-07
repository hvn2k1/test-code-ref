package main

import (
	"context"
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
		"END_USER_ID",
		nil,
	)
	showNewFeature := client.BoolVariation(ctx, testUser, "feature-go-server-e2e-string", false)
	if showNewFeature {
		// The Application code to show the new feature
	} else {
		// The code to run when the feature is off
	}
}

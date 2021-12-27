package main

import (
	"context"
	"log"
	"os"

	download_api "download-api"

	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
)

func main() {

	// Point Cloud Storage SDK to local emulator
	_ = os.Setenv("STORAGE_EMULATOR_HOST", "localhost:9000")

	ctx := context.Background()
	if err := funcframework.RegisterHTTPFunctionContext(ctx, "/", download_api.DownloadFile); err != nil {
		log.Fatalf("funcframework.RegisterHTTPFunctionContext: %v\n", err)
	}
	// Use PORT environment variable, or default to 8080.
	port := "8080"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}
	if err := funcframework.Start(port); err != nil {
		log.Fatalf("funcframework.Start: %v\n", err)
	}
}

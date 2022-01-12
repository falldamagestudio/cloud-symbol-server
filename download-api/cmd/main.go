package main

import (
	"context"
	"log"
	"os"

	download_api "github.com/falldamagestudio/cloud-symbol-server/download-api"

	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
)

func main() {

	ctx := context.Background()
	if err := funcframework.RegisterHTTPFunctionContext(ctx, "/", download_api.DownloadAPI); err != nil {
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

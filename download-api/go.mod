module github.com/falldamagestudio/cloud-symbol-server/download-api

go 1.13

replace github.com/falldamagestudio/cloud-symbol-server/admin-api => ../admin-api

require (
	cloud.google.com/go/firestore v1.6.1
	cloud.google.com/go/storage v1.18.2
	github.com/GoogleCloudPlatform/functions-framework-go v1.5.2
	github.com/falldamagestudio/cloud-symbol-server/admin-api v0.0.0-00010101000000-000000000000 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.1
	github.com/hashicorp/go-retryablehttp v0.7.0
	google.golang.org/api v0.63.0
)

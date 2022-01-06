module github.com/falldamagestudio/cloud-symbol-store/cli

go 1.13

require (
	github.com/falldamagestudio/cloud-symbol-store/upload-api v0.0.0
	github.com/hashicorp/go-retryablehttp v0.7.0
	golang.org/x/oauth2 v0.0.0-20211104180415-d3ed0bb246c8 // indirect
)

replace github.com/falldamagestudio/cloud-symbol-store/upload-api v0.0.0 => ../upload-api

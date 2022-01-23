module github.com/falldamagestudio/cloud-symbol-server/download-api/test

go 1.13

replace github.com/falldamagestudio/cloud-symbol-server/admin-api => ../../admin-api

require (
	github.com/falldamagestudio/cloud-symbol-server/admin-api v0.0.0-00010101000000-000000000000
	github.com/hashicorp/go-cleanhttp v0.5.1
	github.com/hashicorp/go-retryablehttp v0.7.0
)

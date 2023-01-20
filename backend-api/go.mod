module github.com/falldamagestudio/cloud-symbol-server/backend-api

go 1.16

require (
	cloud.google.com/go/cloudsqlconn v1.1.1
	cloud.google.com/go/firestore v1.9.0
	cloud.google.com/go/storage v1.29.0
	github.com/GoogleCloudPlatform/functions-framework-go v1.6.1
	github.com/friendsofgo/errors v0.9.2
	github.com/gorilla/mux v1.8.0
	github.com/jackc/pgconn v1.13.0
	github.com/jackc/pgerrcode v0.0.0-20220416144525-469b46aa5efa
	github.com/lestrrat-go/jwx v1.2.25
	github.com/rs/cors v1.8.3
	github.com/stretchr/testify v1.8.1
	github.com/volatiletech/null/v8 v8.1.2
	github.com/volatiletech/sqlboiler/v4 v4.14.0
	github.com/volatiletech/strmangle v0.0.4
	golang.org/x/oauth2 v0.4.0
	google.golang.org/api v0.108.0
	google.golang.org/grpc v1.52.0
)

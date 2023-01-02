module github.com/falldamagestudio/cloud-symbol-server/admin-api

go 1.16

require (
	cloud.google.com/go/cloudsqlconn v1.0.1
	cloud.google.com/go/firestore v1.6.1
	cloud.google.com/go/storage v1.27.0
	github.com/GoogleCloudPlatform/functions-framework-go v1.5.2
	github.com/friendsofgo/errors v0.9.2
	github.com/gorilla/mux v1.8.0
	github.com/jackc/pgconn v1.13.0
	github.com/jackc/pgerrcode v0.0.0-20220416144525-469b46aa5efa
	github.com/lestrrat-go/jwx v1.2.25
	github.com/rs/cors v1.8.3
	github.com/stretchr/testify v1.8.0
	github.com/volatiletech/null/v8 v8.1.2
	github.com/volatiletech/sqlboiler/v4 v4.13.0
	github.com/volatiletech/strmangle v0.0.4
	golang.org/x/oauth2 v0.0.0-20221014153046-6fdb5e3db783
	google.golang.org/api v0.101.0
	google.golang.org/grpc v1.50.1
)

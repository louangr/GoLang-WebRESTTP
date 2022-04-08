module projectRESTAPI

go 1.15

require github.com/gorilla/mux v1.8.0

// require github.com/boltdb/bolt v1.3.1

require internal/entities v1.0.0

replace internal/entities => ./internal/entities

require internal/web/rest v1.0.0

replace internal/web/rest => ./internal/web/rest

require internal/persistence v1.0.0

replace internal/persistence => ./internal/persistence

require (
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751 // indirect
	github.com/go-openapi/spec v0.20.5 // indirect
	github.com/go-openapi/swag v0.21.1 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/swaggo/http-swagger v1.2.5 // indirect
	github.com/swaggo/swag v1.8.1 // indirect
	github.com/urfave/cli/v2 v2.4.0 // indirect
	golang.org/x/net v0.0.0-20220407224826-aac1ed45d8e3 // indirect
	golang.org/x/sys v0.0.0-20220406163625-3f8b81556e12 // indirect
	golang.org/x/tools v0.1.10 // indirect
	internal/resources v1.0.0
)

replace internal/resources => ./internal/resources

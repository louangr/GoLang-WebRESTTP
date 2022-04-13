module web/rest

go 1.15

require github.com/gorilla/mux v1.8.0

require internal/entities v1.0.0

replace internal/entities => ../entities

require internal/persistence v1.0.0

replace internal/persistence => ../persistence

require internal/resources v1.0.0

replace internal/resources => ../resources
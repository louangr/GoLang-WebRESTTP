module persitence

go 1.15

require (
	github.com/boltdb/bolt v1.3.1
	github.com/stretchr/testify v1.7.1
)

require internal/entities v1.0.0

replace internal/entities => ../entities
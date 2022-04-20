SET portNumber=8000
SET dbFileName="myBolt"

DEL %dbFileName%.db
DEL %dbFileName%.db.lock
go clean -cache ./cmd/restserver/...
go build -o ./cmd/restserver -v ./cmd/restserver/...
.\cmd\restserver\restserver.exe %portNumber% %dbFileName%
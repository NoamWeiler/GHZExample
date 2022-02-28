# GHZExample

## Run server </br>
make run_server</br>
## Run GHZ </br>
make run_stream</br>

### cli command example:</br>

ghz --insecure --proto ./internal/proto_db/proto_db.proto --call GHZExample.Greeter.SayHello -d '{"Name":"Shula"}' 0.0.0.0:50052
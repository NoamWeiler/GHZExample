# GHZExample

## Run server - make run_server</br>
## Run GHZ - make run_stream</br>

### cli command example:</br>

ghz --insecure --proto ./internal/mutual_db/mutual_db.proto --call SensorServer.SensorStream.SensorMeasure -d '{"m":13,"serial":"ser345345"}' 0.0.0.0:50051
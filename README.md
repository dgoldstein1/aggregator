# aggregator
Aggregator service for product API

## Install

```bash
go mod vendor
```

## Build

```bash
go build -o aggregator
```

## Docker Build

```bash
docker build . -t dgoldstein1/aggregator
```

## Run


### Docker

```bash
docker-compose up -d
```

### Locally

```bash
export MONGO_URL="mongodb://localhost:27017"
./aggregator
```

## Insert Data

```bash
time mongo 127.0.0.1/products docker/mongo/insert_data.js
```

## Sample Requests

```bash
curl -s http://localhost:8080/products/13860428 | jq

curl -s http://localhost:8080/products/13860428 \
	-XPUT \
	-d \
	'{"price" : 10.50 }' | jq 

curl -s http://localhost:8080/products/13860428 | jq
```

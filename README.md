# aggregator
Aggregator service for product API

## Install

```bash
go mod vendor
```

## Run

```bash
TODO: build docker container

docker-compose up -d
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

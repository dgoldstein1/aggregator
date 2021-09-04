FROM golang:latest 
RUN mkdir -p /go/src/github.com/dgoldstein1/aggregator
ADD . /go/src/github.com/dgoldstein1/aggregator
WORKDIR /go/src/github.com/dgoldstein1/aggregator
RUN go build -o main . 

CMD ["/go/src/github.com/dgoldstein1/aggregator/main"]
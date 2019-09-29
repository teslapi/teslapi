FROM golang:1.13-alpine
WORKDIR /go/src/github.com/teslapi/teslapi
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN cd /go/src/github.com/teslapi/teslapi/cmd/teslapid/ && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ../../teslapid .
EXPOSE 8080
RUN chmod u+x /go/src/github.com/teslapi/teslapi/teslapid
ENTRYPOINT ["/go/src/github.com/teslapi/teslapi/teslapid"]

FROM golang:1.21.3
WORKDIR /app
COPY . .
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ./cmd/client
ENTRYPOINT ["./client"]
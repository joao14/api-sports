FROM golang:alpine
WORKDIR /go/src/myapp
COPY . .
RUN go build -o /go/bin/myapp src/main.go
FROM golang:latest
WORKDIR /go/src/github.com/victor-leee/portal-be
COPY . .
RUN go build -o main cmd/server/main.go
CMD ["./main"]
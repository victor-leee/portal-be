FROM golang:latest
ENV GOPROXY=https://proxy.golang.com.cn,direct
WORKDIR /go/src/github.com/victor-leee/portal-be
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o main cmd/server/main.go
CMD ["./main"]
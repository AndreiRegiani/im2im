FROM golang:1.20
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /app/main cmd/im2im.go
CMD ["/app/main"]

# Build stage
FROM golang:1.17-alpine3.14 AS builder
RUN mkdir /app
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./...

# Run stage
FROM alpine:3.14
COPY --from=builder /app .
EXPOSE 3000
CMD ["./main"]

FROM golang:1.17 AS builder
RUN mkdir /app
COPY go.mod /app
COPY go.sum /app
WORKDIR /app
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o mainapp ./cmd/main.go

FROM alpine:3.14.2
RUN apk add --no-cache ca-certificates && update-ca-certificates
WORKDIR /app
COPY --from=builder /app/mainapp .
EXPOSE 3000
CMD ["./mainapp"]

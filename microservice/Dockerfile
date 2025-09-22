FROM golang:1.24.3 AS builder

WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o fibserver ./cmd/api/main.go

FROM alpine:3.20
COPY --from=builder /app/fibserver /
RUN apk add --no-cache bash curl
EXPOSE 8080
ENTRYPOINT ["/fibserver"]
# ---- build stage ----
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Install git for go get
RUN apk add --no-cache git

# Copy go.mod and go.sum first for caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the code
COPY . .

# Build binary
RUN go build -o app ./main.go

# ---- run stage ----
FROM alpine:3.19

WORKDIR /app
COPY --from=builder /app/app .

# Expose backend port
EXPOSE 8080

CMD ["./app"]

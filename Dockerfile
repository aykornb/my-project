# Stage 1: Build
FROM golang:1.22-alpine AS builder
WORKDIR /app
# Only copy go.mod since go.sum does not exist yet
COPY go.mod ./
RUN go mod download
COPY . .
# Build the application located in cmd/api
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/api

# Stage 2: Runtime
FROM alpine:latest
WORKDIR /root/
# Copy only the binary from the builder
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]
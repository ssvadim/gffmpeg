# Stage 1: Build the Go binary
FROM golang:alpine AS builder

WORKDIR /app

# Copy only go.mod first (go.sum may not exist)
COPY go.mod ./

# Download dependencies (if any) and generate go.sum if needed
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -o echo-app

# Stage 2: Minimal container
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/echo-app .
ENV TEST="Default TEST value"
EXPOSE 8080
CMD ["./echo-app"]

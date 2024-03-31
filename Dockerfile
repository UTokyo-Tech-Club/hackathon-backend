# # FROM golang:1.22-alpine AS builder
# FROM golang:1.22-alpine

# # RUN apk --no-cache add ca-certificates

# WORKDIR /app

# COPY go.mod go.sum ./
# RUN go mod download

# COPY . .

# # RUN CGO_ENABLED=0 GOOS=linux go build -v -o myserver
# RUN go build -o myserver

# # FROM scratch

# RUN apk --no-cache add ca-certificates

# # COPY --from=builder /app/myserver /myserver

# # EXPOSE 8080

# CMD ["./myserver"]




# Builder stage
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o myserver

# Runner stage
FROM alpine:latest
WORKDIR /app
# Install CA certificates for HTTPS connections
RUN apk --no-cache add ca-certificates
# Copy the built binary from the builder stage
COPY --from=builder /app/myserver /app/myserver
CMD ["./myserver"]
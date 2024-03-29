# FROM golang:1.22-alpine AS builder
FROM golang:1.22-alpine

# RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# RUN CGO_ENABLED=0 GOOS=linux go build -v -o myserver
RUN go build -o myserver

# FROM scratch

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/myserver /myserver

# EXPOSE 8080

CMD ["./myserver"]
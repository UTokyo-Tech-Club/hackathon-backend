FROM golang:1.22-alpine AS builder

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -v -o myserver

# FROM scratch

COPY --from=builder /app/myserver /myserver

EXPOSE 8080

CMD ["./myserver"]
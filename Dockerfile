FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -v -o myserver

FROM scratch

COPY --from=builder /app/myserver /myserver

EXPOSE 8080

RUN apk add --no-cache ca-certificates

CMD ["./myserver"]
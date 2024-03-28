FROM golang:1.22 AS builder

RUN apt-get install -y ca-certificates openssl

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -v -o myserver

FROM scratch

COPY --from=builder /app/myserver /myserver

EXPOSE 8080

CMD ["./myserver"]
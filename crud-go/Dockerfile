FROM golang:1.16-alpine

WORKDIR /app

COPY main.go /app/
COPY go.mod /app/

RUN go run build -o crud

CMD ["./crud"]

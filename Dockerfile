FROM golang:1.23.5-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./main.go


FROM alpine:latest

WORKDIR /root

COPY --from=builder ./app .

CMD ["./main"]
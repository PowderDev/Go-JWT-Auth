FROM golang:1.20-alpine as builder
WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build -o auth ./cmd

RUN chmod +x auth


FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/auth .

CMD ["./auth"]


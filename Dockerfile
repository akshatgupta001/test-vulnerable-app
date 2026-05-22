FROM golang:1.18-alpine AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum* ./
RUN go mod download

COPY main.go .
RUN CGO_ENABLED=0 go build -o scs-demo main.go

FROM alpine:3.14

RUN apk add --no-cache ca-certificates

WORKDIR /root/

COPY --from=builder /app/scs-demo .

EXPOSE 8080

CMD ["./scs-demo"]

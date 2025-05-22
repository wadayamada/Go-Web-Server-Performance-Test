FROM golang:1.20 AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
ENV CGO_ENABLED=0
RUN go build -o server
FROM alpine:3.17
WORKDIR /app
COPY --from=builder /app/server /app/
EXPOSE 8081
CMD ["/app/server"]

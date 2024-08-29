FROM golang:1.23 AS builder

ARG CGO_ENABLED=0
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build

FROM scratch
COPY --from=builder /app/payment-service /payment-service
ENTRYPOINT ["/payment-service"]
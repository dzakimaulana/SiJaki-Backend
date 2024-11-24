FROM golang:1.21 AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./cmd ./cmd
COPY ./internal ./internal

RUN go build -o /app/bin/sijaki ./cmd/api/main.go

FROM gcr.io/distroless/static:nonroot

WORKDIR /app
COPY --from=builder /app/bin /app/bin

USER nonroot:nonroot

CMD ["/app/bin/sijaki"]

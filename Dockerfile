FROM golang:1.17 AS builder
RUN go version

WORKDIR /src
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /tmp/docgo .

FROM alpine:3.15

RUN apk add --no-cache tzdata

WORKDIR /app
COPY --from=builder /tmp/docgo .

ENTRYPOINT ["./docgo"]

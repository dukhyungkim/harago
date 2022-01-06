FROM golang:1.17 AS builder
RUN go version

WORKDIR /src
COPY . .

ARG GOPRIVATE
ENV GOPRIVATE=$GOPRIVATE

ARG GOPRIVATE_TOKEN
ENV GOPRIVATE_TOKEN=$GOPRIVATE_TOKEN

RUN git config --global url."https://x-access-token:${GOPRIVATE_TOKEN}@github.com".insteadOf "https://github.com"
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /tmp/docgo .

FROM alpine:3.15

RUN apk add --no-cache tzdata

WORKDIR /app
COPY --from=builder /tmp/docgo .

ENTRYPOINT ["./docgo"]

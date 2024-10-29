FROM golang:latest as builder
WORKDIR /app
COPY create /app
RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -o shortify-read /app/src/main.go

FROM alpine

ARG ENVIRONMENT
LABEL maintainer = "shortifyReadApp"

WORKDIR /app
COPY --from=builder /app/shortify-read /app
COPY --from=builder /app/config.${ENVIRONMENT}.yaml /app
COPY --from=builder /app/init.sql /app
ENTRYPOINT ./shortify-read
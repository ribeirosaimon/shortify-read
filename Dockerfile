FROM golang:latest as builder
COPY . /app
WORKDIR /app
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o shortify /app/src/main.go

FROM alpine
ARG ENVIRONMENT
LABEL maintainer = "shortify-readApp"
WORKDIR /app
COPY --from=builder /app/shortify-read /app
COPY --from=builder /app/config.${ENVIRONMENT}.yaml /app
COPY --from=builder /app/init.sql /app
ENTRYPOINT ./shortify
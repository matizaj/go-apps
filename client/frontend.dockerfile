FROM golang:1.22-alpine as builder
RUN mkdir /app
COPY . /app

WORKDIR /app
RUN CGO_ENABLED=0 go build -o frontendApp ./cmd/web
RUN chmod +x /app/frontendApp

FROM alpine:latest
RUN mkdir /app
COPY --from=builder /app/frontendApp /app
CMD ["/app/frontendApp"]
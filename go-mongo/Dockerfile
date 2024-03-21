FROM golang:1.22-alpine as builder
RUN mkdir /app
COPY . /app

WORKDIR /app
RUN CGO_ENABLED=0 go build -o App ./cmd/api
RUN chmod +x /app/App

FROM alpine:latest
RUN mkdir /app
COPY --from=builder /app/App /app
CMD ["/app/App"]
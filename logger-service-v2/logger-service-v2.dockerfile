FROM golang:1.22-alpine as builder
RUN mkdir /app
COPY . /app

WORKDIR /app
RUN CGO_ENABLED=0 go build -o loggerSvc ./cmd/api
RUN chmod +x /app/loggerSvc

FROM alpine:latest
RUN mkdir /app
COPY --from=builder /app/loggerSvc /app
CMD ["/app/loggerSvc"]
#FROM golang:1.22-alpine as builder
#RUN mkdir /app
#COPY . /app
#
#WORKDIR /app
#RUN CGO_ENABLED=0 go build -o listenerApp ./cmd/api
#RUN chmod +x /app/listenerApp
#
#FROM alpine:latest
#RUN mkdir /app
#COPY --from=builder /app/listenerApp /app
#CMD ["/app/listenerApp"]

FROM alpine:latest
RUN mkdir /app
COPY listenerApp /app
CMD ["/app/listenerApp"]



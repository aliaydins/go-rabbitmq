FROM golang:alpine
WORKDIR /build
COPY services.delivery services.delivery
WORKDIR /build/services.delivery
RUN go build -o main ./src/main.go
WORKDIR /dist
RUN cp /build/services.delivery/main .
RUN cp /build/services.delivery/src/app.env .
CMD ["/dist/main"]
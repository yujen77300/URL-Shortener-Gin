FROM golang:1.19-alpine3.16
WORKDIR /src
COPY . .
RUN go build -o main .



FROM alpine
WORKDIR /src
COPY --from=0 /src/main .
COPY --from=0 /src/views /src/views
COPY --from=0 /src/static /src/static
COPY config.env /src/config.env


CMD ["./main"]
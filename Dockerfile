FROM golang:latest AS BUILD

WORKDIR /build

COPY . .

RUN go mod download

RUN go build -o /build/main ./cmd/main.go

FROM alpine

WORKDIR /run

COPY --from=BUILD /build/config.json /build/main ./

EXPOSE 8080

CMD ["/run/main"]
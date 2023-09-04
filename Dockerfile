FROM golang:latest AS BUILD

WORKDIR /build

COPY . .

RUN go mod download

RUN go build -o /build/main ./cmd/main.go

FROM ubuntu

COPY --from=BUILD /build/config.json /build/main /run/

EXPOSE 8080

WORKDIR /run

ENTRYPOINT /run/main
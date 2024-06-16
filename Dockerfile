FROM golang:1.22 AS DEPS

WORKDIR /build

COPY . .

RUN go mod download

FROM DEPS AS BUILD

RUN go build -o /build/main ./cmd/main.go

FROM ubuntu

COPY --from=BUILD /build/ /run/

EXPOSE 5000

WORKDIR /run

RUN ls

ENTRYPOINT /run/main

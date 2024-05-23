FROM golang:1.22 AS BUILD

WORKDIR /build

COPY . .

#RUN go mod download
#
#RUN go build -o /build/main ./cmd/main.go

FROM ubuntu

COPY --from=BUILD /build/ /run/

EXPOSE 5000

WORKDIR /run

RUN ls

ENTRYPOINT /run/main

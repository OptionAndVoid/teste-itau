FROM golang:alpine AS build-stage

WORKDIR /app

RUN apk add --no-cache make

COPY go.mod go.sum Makefile ./

RUN go mod download

COPY cmd cmd
COPY internal internal
COPY pkg pkg

RUN make build

FROM alpine:latest AS release-stage

COPY --from=build-stage /app/bin/teste-itau /teste-itau

ENTRYPOINT ["/teste-itau"]

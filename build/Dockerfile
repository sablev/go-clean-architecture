FROM golang:1.15-alpine AS deps

RUN apk update && apk upgrade && apk add --no-cache --update git mc

WORKDIR /go/src/app

COPY go.mod .
COPY go.sum .
RUN go mod download


FROM deps AS builder

COPY . /tmp/app/
WORKDIR /tmp/app/
RUN go build -ldflags='-w -s' -o ./bin/api ./cmd/api/main.go


FROM alpine

RUN apk update && apk upgrade && apk add --no-cache --update ca-certificates mc

WORKDIR /
COPY --from=builder /tmp/app/bin ./app
COPY ./configs/ /app/config/
RUN chmod 755 /app/*

WORKDIR /app/
CMD ["./api"]
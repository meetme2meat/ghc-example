FROM golang:1.16 AS builder
RUN mkdir -p /apps/
WORKDIR /apps/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app .

FROM alpine:latest  
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN apk add bash
RUN mkdir -p /apps
WORKDIR /apps
COPY --from=builder /apps/app .
CMD ["./app"]
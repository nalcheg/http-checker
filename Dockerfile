FROM golang:1.13.4-alpine3.10 AS builder

RUN apk add --update ca-certificates

ADD . /build/
RUN cd /build && CGO_ENABLED=0 GOOS=linux go build -ldflags '-w -s' -a -installsuffix cgo -o /http-checker


FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /http-checker /http-checker

ENTRYPOINT ["/http-checker"]

FROM golang:1.15-alpine AS builder

ADD . /go/src/informator-69-bot

RUN cd /go/src/informator-69-bot && \
    go build -o /srv/informator-69-bot ./app

FROM alpine:3.14.0

WORKDIR /srv

COPY --from=builder /srv/informator-69-bot /srv

CMD ["/srv/informator-69-bot"]

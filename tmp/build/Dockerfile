FROM alpine:3.18

ADD tmp/_output/bin/smtp-operator /usr/local/bin/smtp-operator

RUN adduser -D smtp-operator
USER smtp-operator

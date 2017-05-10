FROM alpine:3.5

ADD . /daxxCore
RUN \
  apk add --update git go make gcc musl-dev linux-headers && \
  (cd daxxcore&& make geth)                           && \
  cp daxxCore/build/bin/geth /geth                     && \
  apk del git go make gcc musl-dev linux-headers          && \
  rm -rf /daxxcore&& rm -rf /var/cache/apk/*

EXPOSE 8545
EXPOSE 30303

ENTRYPOINT ["/geth"]

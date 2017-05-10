FROM alpine:3.5

RUN \
  apk add --update go git make gcc musl-dev linux-headers ca-certificates && \
  git clone --depth 1 --branch release/1.5 https://github.com/daxxcoin/daxxcore && \
  (cd daxxcore&& make geth) && \
  cp daxxCore/build/bin/geth /geth && \
  apk del go git make gcc musl-dev linux-headers && \
  rm -rf /daxxcore&& rm -rf /var/cache/apk/*

EXPOSE 8545
EXPOSE 30303

ENTRYPOINT ["/geth"]

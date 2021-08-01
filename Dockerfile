FROM alpine:3.6

RUN apk add --no-cache \
        ca-certificates \
        bash \
        libc6-compat \
    && rm -f /var/cache/apk/*

COPY bin/suave /usr/local/bin/suave

CMD ["/usr/local/bin/suave"]

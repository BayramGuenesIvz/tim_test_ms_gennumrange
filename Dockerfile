FROM alpine:latest

ADD cert/* /etc/ssl/certs/

# security updates, useful packages/tools
RUN apk update && \
    apk add \
        bash \
        curl \
        shadow \
        wget && \
    rm -rf /var/cache/apk/*

#RUN groupadd -g 3203 wdr-a2k8s
#RUN usermod -a -G wdr-a2k8s root


#COPY cmd/tim_test_ms_gennumrange/main app/main
COPY ./main app/main
COPY web app/web

WORKDIR /app

CMD ./main
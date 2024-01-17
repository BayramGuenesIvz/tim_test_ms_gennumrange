#FROM ghcr.io/westdeutscherrundfunkkoeln/am-docker-alpine-base:3.19-latest
#FROM ghcr.io/westdeutscherrundfunkkoeln/am-docker-oracle-ic-base:main-4
#COPY assets/timFileSys /timFileSys
#COPY cmd/tim_test_ms_gennumrange/main app/main
#COPY web app/web

#WORKDIR /app
#
#CMD ./main 



FROM alpine:latest
#FROM ghcr.io/westdeutscherrundfunkkoeln/am-docker-alpine-base:3.19-latest
#FROM ghcr.io/westdeutscherrundfunkkoeln/am-docker-oracle-ic-base:main-4
#COPY main app/main
ADD cert/* /etc/ssl/certs/

# security updates, useful packages/tools
RUN apk update && \
    apk add \
        bash \
        curl \
        shadow \
        wget && \
    rm -rf /var/cache/apk/*

RUN groupadd -g 3203 wdr-a2k8s
RUN usermod -a -G wdr-a2k8s root


#COPY cmd/tim_test_ms_gennumrange/main app/main
COPY ./main app/main
COPY web app/web

WORKDIR /app

CMD ./main
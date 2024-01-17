FROM oraclelinux:9-slim

ADD cert/* /etc/ssl/certs/

RUN microdnf install oracle-instantclient-release-el9
RUN microdnf install wget
RUN microdnf install findutils
RUN microdnf clean all

RUN groupadd -g 3203 wdr-a2k8s
RUN usermod -a -G wdr-a2k8s root


COPY ./main app/main
COPY web app/web

WORKDIR /app
USER root
CMD ./main
FROM ghcr.io/westdeutscherrundfunkkoeln/am-docker-alpine-base:3.19-latest
#FROM ghcr.io/westdeutscherrundfunkkoeln/am-docker-oracle-ic-base:main-4
#COPY assets/timFileSys /timFileSys
COPY cmd/tim_test_ms_gennumrange/main app/main
COPY web app/web

WORKDIR /app
#

CMD ./main 

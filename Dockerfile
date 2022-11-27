# multi stage-> build stage
FROM golang:1.17-alpine3.15 AS build-env
RUN apk add build-base

ADD . /src

RUN cd /src && go build -o ./bin/temp-file-transfer-app

# multi stage-> deploy stage
FROM alpine

WORKDIR /app

COPY --from=build-env /src/bin/temp-file-transfer-app /app/

RUN mkdir /app/uploads

EXPOSE 9090

ENTRYPOINT ./temp-file-transfer-app
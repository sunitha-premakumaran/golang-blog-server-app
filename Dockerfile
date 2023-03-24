FROM golang:1.20.2-alpine3.17 as base

ARG APP_NAME

ADD . /go/src/${APP_NAME}

ENV WORKDIR="/go/src/${APP_NAME}"

WORKDIR ${WORKDIR}


RUN go get ${APP_NAME}


RUN go install

ENV GO_PROJECT_BINARY="/go/bin/${APP_NAME}"


ENTRYPOINT [${GO_PROJECT_BINARY}]
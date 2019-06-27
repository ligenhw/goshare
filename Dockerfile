FROM golang

ENV APP_DIR goshare
ENV GOPROXY https://goproxy.io

ADD . $APP_DIR
WORKDIR $APP_DIR
RUN go install
ENTRYPOINT $GOPATH/bin/goshare

EXPOSE 8080

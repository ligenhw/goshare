FROM golang

ENV APP_DIR $GOPATH/src/github.com/ligenhw/goshare

ADD . $APP_DIR
WORKDIR $APP_DIR
RUN go get ./... && go install github.com/ligenhw/goshare
ENTRYPOINT $GOPATH/bin/goshare

EXPOSE 8080

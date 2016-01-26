FROM golang:1.4

ENV MYSQL_API_PORT=80

COPY ./daemon /go/src/github.com/ravaj-group/farmer/mysql/api/daemon
WORKDIR  /go/src/github.com/ravaj-group/farmer/mysql/api/daemon

RUN \
    go get -d -v && \
    go install -v

EXPOSE 80

CMD ["daemon"]

FROM golang

ADD . /go/src/github.com/PeterAkaJohn/SightSeeing/server/
RUN go install github.com/PeterAkaJohn/SightSeeing/server/
ENTRYPOINT /go/bin/src

EXPOSE 8443

FROM golang:1.6

RUN mkdir -p /go/src/github.com/byuoitav
ADD . /go/src/github.com/byuoitav/pjlink-service

WORKDIR /go/src/github.com/byuoitav/pjlink-service
RUN go get -d -v
RUN go install -v

CMD ["/go/bin/pjlink-service"]

EXPOSE 8005
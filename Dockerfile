FROM golang:1.8.1

ENV APPROOT ${GOPATH}/src/github.com/suusan2go/familog-api

WORKDIR ${APPROOT}

COPY . ${APPROOT}

RUN set -x \
  && go get -u github.com/golang/dep/... \
  && dep ensure \
  && go build -o bin/familog

EXPOSE 8080:8080

CMD bin/familog

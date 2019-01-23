FROM golang:alpine

RUN apk add --no-cache git mercurial

RUN git clone https://github.com/jcmturner/gokrb5 /go/src/gopkg.in/jcmturner/gokrb5.v5
RUN cd /go/src/gopkg.in/jcmturner/gokrb5.v5 && git checkout tags/v5.3.0
RUN git clone https://github.com/golang/crypto /go/src/golang.org/x/crypto

LABEL max-go.version="0.7" maintainer="Jeorch"
RUN go get github.com/alfredyang1986/blackmirror
RUN go get github.com/Jeorch/max-go

ADD deploy-config/ /go/bin/

RUN go install -v github.com/Jeorch/max-go

WORKDIR /go/bin

ENTRYPOINT ["max-go"]

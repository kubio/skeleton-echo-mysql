FROM golang:onbuild

WORKDIR /go/src/app

RUN go get -u gopkg.in/godo.v2/cmd/godo

RUN go-wrapper download
RUN go-wrapper install

EXPOSE 5000

FROM golang:latest

ENV GO111MODULE on

WORKDIR /go/src/github.com/yasukotelin/termcolor

COPY . /go/src/github.com/yasukotelin/termcolor

RUN go build

CMD ["/bin/bash"]

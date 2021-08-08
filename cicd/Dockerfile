FROM golang:1.16

RUN apt update
RUN apt upgrade -y
RUN apt install gcc-arm-linux-gnueabi -y

CMD /go/src/github.com/housekpr/core/cicd/build.sh
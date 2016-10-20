FROM golang:1.6-onbuild

ADD . /go/src/github.com/Zhanat87/stack-auth
CMD cd /go/src/github.com/Zhanat87/stack-auth; go get github.com/Zhanat87/stack-auth && go build -o /go/bin/stack-auth

EXPOSE 8082

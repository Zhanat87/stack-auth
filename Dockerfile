# golang image where workspace (GOPATH) configured at /go.
FROM golang:1.6-onbuild
# FROM golang
# Copy the local package files to the container’s workspace.
ADD . /go/src/github.com/Zhanat87/stack-auth
# Setting up working directory
WORKDIR /go/src/github.com/Zhanat87/stack-auth
# Get godeps for managing and restoring dependencies
RUN go get github.com/tools/godep
# Restore godep dependencies
# RUN godep restore
# Build the stack-auth command inside the container.
RUN go install github.com/Zhanat87/stack-auth
# Run the stack-auth command when the container starts.
ENTRYPOINT /go/bin/stack-auth
# Service listens on port 8082.
EXPOSE 8082

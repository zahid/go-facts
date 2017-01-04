# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/zahid/go-facts/
WORKDIR /go/src/github.com/zahid/go-facts/

RUN go get

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/go-facts

# Document that the service listens on port 8080.
EXPOSE 8000
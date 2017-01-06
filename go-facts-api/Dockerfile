# Use these to build then run this image
# docker build -t zahid/go-facts-api api/
# docker run --name go-facts-api -d zahid/go-facts-api
#
# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/zahid/go-facts-api
WORKDIR /go/src/github.com/zahid/go-facts-api

RUN go get

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/go-facts-api

# Document that the service listens on port 8080.
EXPOSE 8080
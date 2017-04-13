# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:alpine

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/masaruhoshi/go-palindrome

# Get dependency packages
RUN go-wrapper download
RUN go-wrapper install

# Build the app inside the container.
RUN go install github.com/masaruhoshi/go-palindrome

# Run the app by default when the container starts.
ENTRYPOINT /go/bin/gopal

# Document that the service listens on port 80.
EXPOSE 8080
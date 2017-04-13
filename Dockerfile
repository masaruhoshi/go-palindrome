# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:alpine

# Copy the remote source files to the container's workspace.
RUN mkdir -p /go/src

# Add workdir
WORKDIR /go/src

# Get dependency packages
## git doesn't come in alpine. This is a workaround to make it work with go-wrapper
RUN apk add --no-cache git mercurial supervisor && \
	# Clone source code repo
	git clone https://github.com/masaruhoshi/go-palindrome.git

# Move to source code path
WORKDIR go-palindrome

# Install dependencies
RUN go-wrapper download && \
	go-wrapper install && \
	# Get rid of git at the end
	apk del git mercurial && \
	# Build the app inside the container.
	go install go-palindrome && \
	# Clean up 
	rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

# Update path to include the service
ENV PATH="/go/bin:${PATH}"

# Document that the service listens on port 80.
EXPOSE 8080 27017

# Run supervisor 
ENTRYPOINT ["/usr/bin/supervisord", "-c", "/go/src/go-palindrome/supervisor.conf"]
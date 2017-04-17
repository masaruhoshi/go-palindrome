# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:alpine

# Directory for source files
RUN mkdir -p /go/src && \
	# This is for mongodb. Can be mapped to local mongo data if required
	mkdir -p /data/db

# Add workdir
WORKDIR /go/src

# Get dependency packages
## git, supervisor and mongo don't come in alpine. 
RUN echo http://dl-4.alpinelinux.org/alpine/edge/testing >> /etc/apk/repositories && \
	apk add --no-cache git mercurial supervisor mongodb && \
	# Clone source code repo
	git clone https://github.com/masaruhoshi/go-palindrome.git

# Move to source code path
WORKDIR go-palindrome

# Install dependencies
RUN go-wrapper download && \
	# Get rid of git at the end
	apk del git mercurial && \
	# Build the app inside the container.
	go install go-palindrome && \
	# Clean up 
	rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

# Document that the service listens on port 80.
EXPOSE 80 27017

# Run supervisor 
ENTRYPOINT ["/usr/bin/supervisord", "-c", "/go/src/go-palindrome/conf/supervisor.conf"]
FROM golang:1.10

# Get dependencies
RUN go get golang.org/x/oauth2
RUN go get cloud.google.com/go/compute/metadata

RUN go get github.com/gin-gonic/gin
run go get github.com/gin-gonic/contrib/static

# Add local volume
ADD . /go/src/github.com/schweigert/inscreveai

# Build web app
RUN go install github.com/schweigert/inscreveai/web

# Add public volume
ADD ./public ./public

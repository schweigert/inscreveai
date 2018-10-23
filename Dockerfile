FROM golang:1.10

# Get dependencies
RUN go get golang.org/x/oauth2
RUN go get cloud.google.com/go/compute/metadata

RUN go get github.com/gin-gonic/gin
RUN go get github.com/gin-gonic/contrib/static
RUN go get github.com/gorilla/context
RUN go get github.com/boj/redistore
RUN go get github.com/gorilla/sessions

# Add local volume
ADD . /go/src/github.com/schweigert/inscreveai

# Build web app
RUN go install github.com/schweigert/inscreveai/web

# Add public volume
ADD ./public ./public

# Add credentials
ADD ./google_creds.json ./google_creds.json

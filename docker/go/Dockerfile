FROM golang
ADD ./app/ /go/src/app
EXPOSE 8080
CMD ["/usr/local/go/bin/go", "run", "/go/src/app/server.go"]
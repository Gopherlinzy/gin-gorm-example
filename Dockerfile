FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/github.com/Gopherlinzy/go-gin-example
COPY . $GOPATH/github.com/Gopherlinzy/go-gin-example
RUN go build .

EXPOSE 8080
ENTRYPOINT ["./go-gin-example"]
FROM golang

WORKDIR /gopath/app
ENV GOPATH /gopath/app
ADD . /gopath/app/

RUN go get github.com/toukii/upload
RUN go build -o upload
RUN mkdir -p static

EXPOSE 80
CMD ["/gopath/app/upload"]

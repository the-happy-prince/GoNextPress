FROM golang:alpine3.19
RUN wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.56.2
WORKDIR /go/src/app
COPY . .
RUN go mod init
RUN go build  -o main
EXPOSE 80
RUN golangci-lint run
CMD ["./main"]
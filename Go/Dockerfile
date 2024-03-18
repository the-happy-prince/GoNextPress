FROM golang:1.16-alpine
WORKDIR /go/src/app
COPY . .
RUN go mod init
RUN go build  -o main
EXPOSE 4000
CMD ["./main"]

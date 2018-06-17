FROM golang:1.10.3-stretch
RUN go get github.com/gorilla/mux
RUN go get github.com/gocql/gocql
RUN mkdir /kamilProject
ADD . /kamilProject
WORKDIR /kamilProject
RUN go build -o main .
CMD ["/main"]


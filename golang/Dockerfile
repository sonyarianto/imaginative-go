FROM golang:latest

RUN go get -d -v github.com/go-sql-driver/mysql
RUN go install -v github.com/go-sql-driver/mysql

RUN go get -d -v github.com/mongodb/mongo-go-driver/mongo
RUN go install -v github.com/mongodb/mongo-go-driver/mongo

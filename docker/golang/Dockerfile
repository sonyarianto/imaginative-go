FROM golang:latest

RUN go get -d -v github.com/go-sql-driver/mysql
RUN go install -v github.com/go-sql-driver/mysql

RUN go get -d -v github.com/mongodb/mongo-go-driver/mongo
RUN go install -v github.com/mongodb/mongo-go-driver/mongo

RUN go get -d -v github.com/julienschmidt/httprouter
RUN go install -v github.com/julienschmidt/httprouter

RUN go get -d -v github.com/alecthomas/chroma/formatters/html
RUN go install -v github.com/alecthomas/chroma/formatters/html

RUN go get -d -v github.com/alecthomas/chroma/lexers
RUN go install -v github.com/alecthomas/chroma/lexers

RUN go get -d -v github.com/alecthomas/chroma/styles
RUN go install -v github.com/alecthomas/chroma/styles

RUN go get -d -v gopkg.in/russross/blackfriday.v2
RUN go install -v gopkg.in/russross/blackfriday.v2
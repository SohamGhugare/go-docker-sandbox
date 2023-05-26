FROM golang:1.20.4-alpine3.18
RUN mkdir /api
ADD . /api
WORKDIR /api
RUN go build -o main .
CMD ["/api/main"]

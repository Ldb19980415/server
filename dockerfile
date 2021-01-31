FROM golang:latest

WORKDIR /src/server

COPY . . 

EXPOSE 3009

CMD ["go" , "run server.go"]
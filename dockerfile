FROM golang:latest

WORKDIR /src/server

COPY . .

RUN go build -o main.exe server.go

WORKDIR /src/app

COPY /src/server/main.exe .

RUN rm -rf /src/server

EXPOSE 3005

ENTRYPOINT ["./main.exe"]
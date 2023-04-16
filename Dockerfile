FROM golang:1.20.1

WORKDIR /home

COPY . /home/

CMD [ "go", "run", "/home/app/ping-pong-server.go" ]
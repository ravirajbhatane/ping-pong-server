FROM golang:1.20.1

RUN mkdir -p /home/app

COPY ./app/*.go /home/app

CMD [ "go", "run", "/home/app/ping-pong-server.go" ]
# syntax=docker/dockerfile:1
FROM golang:1.20.1

# Set destination for COPY
WORKDIR /home

# Copy source code
COPY . .

# Download Go modules
RUN go mod download

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-ping-pong ./app/ping-pong-server.go

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/engine/reference/builder/#expose
EXPOSE 8090

# Run
CMD ["/docker-ping-pong"]
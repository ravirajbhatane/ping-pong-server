# ping-pong-server
GoLang PingPong server

# Build instructions
cd ../ping-pong-server
docker build -t ping-pong-server:1.0 .
docker run -p 8090:8090 ping-pong-server:1.0

# Testing
curl localhost:8090/ping -i
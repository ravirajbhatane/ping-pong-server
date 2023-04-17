# ping-pong-server
GoLang PingPong server

## Local
### Build
```
  go build  -o ping-pong app/ping-pong-server.go
``` 
### Run
```
./ping-pong
```
### Testing
- Ping
    ```
    curl --location 'http://localhost:8090/ping'
    ```
- Create user
    ```
    curl --location 'http://localhost:8090/user' \
    --header 'Content-Type: text/plain' \
    --data '{
        "Name":"alex",
        "Gender":"male",
        "Age":25
    }'
    ```
- Get user
    ```
    curl --location 'http://localhost:8090/user/643c8736a6e6d55fa6936749'
    ```
- Delete user
    ```
    curl --location --request DELETE 'http://localhost:8090/user/643c8736a6e6d55fa6936749'
    ```

## Docker
```
cd ../ping-pong-server
docker build -t ping-pong-server:1.0 .
docker run -p 8090:8090 ping-pong-server:1.0
```

### Testing
```
curl localhost:8090/ping -i

HTTP/1.1 200 OK
Date: Sun, 12 Mar 2023 23:24:52 GMT
Content-Length: 5
Content-Type: text/plain; charset=utf-8

pong
```

## Docker-Compose
### Build and Start
```
docker-compose up -d --build
```

### Stop
```
docker-compose down
```

## K8s
- Create K8s resources
```
kubectl apply -f pingpong.yaml

deployment.apps/pingpong-deployment created
service/pingpong-service created
```

- Get resource information
```
kubectl get all

NAME                                       READY   STATUS    RESTARTS   AGE
pod/pingpong-deployment-79bd9fdf6d-74c7z   1/1     Running   0          100s

NAME                       TYPE           CLUSTER-IP     EXTERNAL-IP   PORT(S)          AGE
service/kubernetes         ClusterIP      10.96.0.1      <none>        443/TCP          6d16h
service/pingpong-service   LoadBalancer   10.105.88.69   <pending>     8090:30000/TCP   100s

NAME                                  READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/pingpong-deployment   1/1     1            1           100s

NAME                                             DESIRED   CURRENT   READY   AGE
replicaset.apps/pingpong-deployment-79bd9fdf6d   1         1         1       100s
```

- Run minikube service
```
minikube service pingpong-service

|-----------|------------------|-------------|---------------------------|
| NAMESPACE |       NAME       | TARGET PORT |            URL            |
|-----------|------------------|-------------|---------------------------|
| default   | pingpong-service |        8090 | http://192.168.49.2:30000 |
|-----------|------------------|-------------|---------------------------|
🏃  Starting tunnel for service pingpong-service.
|-----------|------------------|-------------|------------------------|
| NAMESPACE |       NAME       | TARGET PORT |          URL           |
|-----------|------------------|-------------|------------------------|
| default   | pingpong-service |             | http://127.0.0.1:64471 |
|-----------|------------------|-------------|------------------------|
🎉  Opening service default/pingpong-service in default browser...
❗  Because you are using a Docker driver on darwin, the terminal needs to be open to run it.

```

## Testing
```
curl http://127.0.0.1:64471/ping -i

HTTP/1.1 200 OK
Date: Sun, 12 Mar 2023 23:15:41 GMT
Content-Length: 5
Content-Type: text/plain; charset=utf-8

pong
```
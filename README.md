## Project Architecture Explanation
- cmd/: Contains the main entry point (main.go) that boots up your HTTPS server.
- pkg/: Contains core business logic split into modules:
- server/: Handles setting up the HTTPS server, including TLS configuration and middleware.
- handlers/: Contains HTTP handlers for different endpoints (authentication, user management, etc.).
- mqtt/: Holds the MQTT client code that your services can use to communicate with the MQTT broker.
- configs/: Houses configuration files (YAML, JSON, etc.) to manage environment-specific settings.
- deployments/: Contains all deployment-related files:
- kubernetes/: Kubernetes manifests for deploying your HTTPS service and MQTT broker.
- docker-compose.yml: (Optional) Local multi-container orchestration configuration.
- Dockerfile: Defines how your HTTPS service container is built.
- `go.mod` and `go.sum`: Manage Go module dependencies.

## Docker

Build container image:

```
docker build -t @username/https-server:latest .
```

Launch container image mounting in the specified path the certificates for secure communication:

```
docker run -it -v local_path:/etc/https-server/certs @username/https-server:latest
```

Inspect container

```
docker exec -it container_id /bin/bash
```
## Kubernetes


### MQTT

- Apply MQTT deployment for communication between microservices

```
kubectl apply -f deployments/kubernetes/mqtt/mqtt-deployment.yaml
```

- Apply MQTT service

```
kubectl apply -f deployments/kubernetes/mqtt/mqtt-service.yaml
```

- Verify deployment and check service is running

```
kubectl get pods -l app=mqtt
kubectl get svc mqtt-service
```

### HTTPS Server

Once you build and push the https server docker image, you can start your cluster with a Docker driver by using minikube.

You must apply all the deplyments to create the pod and run the containers declared in each deployment.

- Apply Config Map

```
kubectl apply -f deployments/kubernetes/configmap.yaml
```

- Apply Secret

It should contains your certificate and key encoded in base64.
Please, update this fields according to your certificate and private key PEM files.

Once you finish the encoding, update both `cert.pem` and `key.pem` values. Then, you can apply the config as follows:

```
kubectl apply -f deployments/kubernetes/secret.yaml
```

- Apply HTTPS server config

In order to gather the deployments to the main go, you need to apply this congiuration file:

```
kubectl apply -f deployments/kubernetes/https-server-deployment.yaml
```

Make sure the pod is running

```
kubectl get pods -l app=https-server 
```

Expose deployment through a service

```
kubectl apply -f deployments/kubernetes/https-server-service.yaml
```

Check service has been deployed correctly

```
kubectl get svc https-server  
```

You should be able to see something like:

```
NAME           TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)    AGE
https-server   ClusterIP   10.111.17.234   <none>        8443/TCP   24h
```

meaning you can test connection, via for example port forwarding as

```
kubectl port-forward svc/https-service 8443:8443
curl -k https://localhost:8443/health
```

- Stop running deployment, services and minikube

```
kubectl delete deployment https-server
kubectl delete service https-server
minikube stop
```
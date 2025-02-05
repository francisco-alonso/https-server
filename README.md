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

## Documentation

Apply MQTT deplyment for communication between microservices

```
kubectl apply -f deployments/kubernetes/mqtt/mqtt-deployment.yaml
```

Apply MQTT service

```
kubectl apply -f deployments/kubernetes/mqtt/mqtt-service.yaml
```

Verify deployment and check service is running

```
kubectl get pods -l app=mqtt
kubectl get svc mqtt-service
```
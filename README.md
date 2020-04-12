# Example Service for Helm Charts

## Prerequisites

- [Docker](https://www.docker.com/)
- [Helm](https://helm.sh/)
- [Minikube](https://kubernetes.io/docs/setup/learning-environment/minikube/)
- [Go](https://golang.org/dl/)
- [Python](https://www.python.org/downloads/release/python-380/)
- [Make](https://www.gnu.org/software/make/)

## Build and Run in Local

- Run `make frontend` to install the python libraries
- Run `make run-frontend` to run the frontend
- Run `make backend` to build binary file for backend
- Run `make run-backend` to run the backend

## Build and Publish Docker Image

- Run `make publish-frontend username=<docker hub username> version=<tag-version>`.
- Run `make publish-backend username=<docker hub username> version=<tag-version>`.
- Replace `<username>` and `<tag-version>` with your own value.

## Test and Install Helm Chart

- Test chart for non-production.

```bash
make helm-test-nonprod
```

```yaml
---
# Source: stack/templates/service-backend.yaml
apiVersion: v1
kind: Service
metadata:
  name: backend
spec:
  selector:
    app: backend
  ports:
  - protocol: TCP
    port: 80
    targetPort: http
---
# Source: stack/templates/service-frontend.yaml
apiVersion: v1
kind: Service
metadata:
  name: frontend
spec:
  selector:
    app: frontend
  ports:
  - protocol: "TCP"
    port: 80
    targetPort: http
    nodePort: 30000
  type: NodePort
---
# Source: stack/templates/deployment-backend.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: backend
spec:
  selector:
    matchLabels:
      app: backend
  replicas: 2
  template:
    metadata:
      labels:
        app: backend
    spec:
      containers:
        - name: backend
          image: franzramadhan/backend:0.1.0
          ports:
            - name: http
              containerPort: 8888
---
# Source: stack/templates/deployment-frontend.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: frontend
spec:
  selector:
    matchLabels:
      app: frontend
  replicas: 2
  template:
    metadata:
      labels:
        app: frontend
    spec:
      containers:
        - name: frontend
          image: franzramadhan/frontend:0.1.0
          env:
            - name: API_URL
              value: "http://backend"
          ports:
            - name: http
              containerPort: 8080
```

- Test chart for production
  
  Before test new chart with same resources, make sure to uninstall the existing one using `helm uninstall chartname` e.g `helm uninstall chart-1586699279`

```bash
make helm-test-prod
```

```yaml
# Source: stack/templates/service-backend.yaml
apiVersion: v1
kind: Service
metadata:
  name: backend
spec:
  selector:
    app: backend
  ports:
  - protocol: TCP
    port: 80
    targetPort: http
---
# Source: stack/templates/service-frontend.yaml
apiVersion: v1
kind: Service
metadata:
  name: frontend
spec:
  selector:
    app: frontend
  ports:
  - protocol: "TCP"
    port: 80
    targetPort: http
    nodePort: 30000
  type: NodePort
---
# Source: stack/templates/deployment-backend.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: backend
spec:
  selector:
    matchLabels:
      app: backend
  replicas: 2
  template:
    metadata:
      labels:
        app: backend
    spec:
      containers:
        - name: backend
          image: franzramadhan/backend:0.1.0
          ports:
            - name: http
              containerPort: 8888
---
# Source: stack/templates/deployment-frontend.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: frontend
spec:
  selector:
    matchLabels:
      app: frontend
  replicas: 2
  template:
    metadata:
      labels:
        app: frontend
    spec:
      containers:
        - name: frontend
          image: franzramadhan/frontend:0.1.0
          env:
            - name: API_URL
              value: "http://backend"
          ports:
            - name: http
              containerPort: 8080
---
# Source: stack/templates/hpa-backend.yaml
apiVersion: autoscaling/v2beta2
kind: HorizontalPodAutoscaler
metadata:
  name: backend
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: backend
  minReplicas: 2
  maxReplicas: 10
  metrics:
        - type: Resource
          resource:
            name: cpu
            target:
              type: Utilization
              averageUtilization: 80
        - type: Resource
          resource:
            name: memory
            target:
              type: Utilization
              averageUtilization: 80
---
# Source: stack/templates/hpa-frontend.yaml
apiVersion: autoscaling/v2beta2
kind: HorizontalPodAutoscaler
metadata:
  name: frontend
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: frontend
  minReplicas: 2
  maxReplicas: 10
  metrics:
        - type: Resource
          resource:
            name: cpu
            target:
              type: Utilization
              averageUtilization: 80
        - type: Resource
          resource:
            name: memory
            target:
              type: Utilization
              averageUtilization: 80

```

- Run chart for non-production.

```bash
> $ make helm-run-prod                                                                                                                             [±master ●]
minikube start
😄  minikube v1.9.2 on Darwin 10.15.4
✨  Using the virtualbox driver based on existing profile
👍  Starting control plane node m01 in cluster minikube
🏃  Updating the running virtualbox "minikube" VM ...
🐳  Preparing Kubernetes v1.18.0 on Docker 19.03.8 ...
🌟  Enabling addons: default-storageclass, ingress, storage-provisioner
🏄  Done! kubectl is now configured to use "minikube"
helm lint ./chart
==> Linting ./chart

1 chart(s) linted, 0 chart(s) failed
helm install --generate-name --set isProduction=true ./chart
NAME: chart-1586699314
LAST DEPLOYED: Sun Apr 12 20:48:35 2020
NAMESPACE: default
STATUS: deployed
REVISION: 1
TEST SUITE: None

```

- Validate the list of running pods

```bash
> $ kubectl get pods
NAME                        READY   STATUS    RESTARTS   AGE
backend-59884b5db5-2jsjc    1/1     Running   0          7m31s
backend-59884b5db5-nsz9n    1/1     Running   0          7m31s
frontend-5968587999-g2lvn   1/1     Running   0          7m31s
frontend-5968587999-qlp7v   1/1     Running   0          7m31s
```

- Run chart for production

  Before run new chart with same resources, make sure to uninstall the existing one using `helm uninstall chartname` e.g `helm uninstall chart-1586699279`

```bash
> $ make helm-run-nonprod                                                                                                                          [±master ●]
minikube start
😄  minikube v1.9.2 on Darwin 10.15.4
✨  Using the virtualbox driver based on existing profile
👍  Starting control plane node m01 in cluster minikube
🏃  Updating the running virtualbox "minikube" VM ...
🐳  Preparing Kubernetes v1.18.0 on Docker 19.03.8 ...
🌟  Enabling addons: default-storageclass, ingress, storage-provisioner
🏄  Done! kubectl is now configured to use "minikube"
helm lint ./chart
==> Linting ./chart

1 chart(s) linted, 0 chart(s) failed
helm install --generate-name ./chart
NAME: chart-1586699279
LAST DEPLOYED: Sun Apr 12 20:48:00 2020
NAMESPACE: default
STATUS: deployed
REVISION: 1
TEST SUITE: None
```

- Validate the list of running pods

```bash
> $ kubectl get pods
NAME                        READY   STATUS        RESTARTS   AGE
backend-59884b5db5-9t2mv    1/1     Running       0          10s
backend-59884b5db5-fjpnn    1/1     Running       0          10s
frontend-5968587999-g2lvn   1/1     Terminating   0          13m
frontend-5968587999-kct5k   1/1     Running       0          10s
frontend-5968587999-p6bqz   1/1     Running       0          10s
frontend-5968587999-qlp7v   1/1     Terminating   0          13m
```

- Get URL of frontend

```bash
> $ minikube service list
|-------------|------------|--------------|-----------------------------|
|  NAMESPACE  |    NAME    | TARGET PORT  |             URL             |
|-------------|------------|--------------|-----------------------------|
| default     | backend    | No node port |
| default     | frontend   |           80 | http://192.168.99.102:30000 |
| default     | kubernetes | No node port |
| kube-system | kube-dns   | No node port |
|-------------|------------|--------------|-----------------------------|

> $ minikube service frontend --url
http://192.168.99.102:30000

```

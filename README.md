# Kubernetes Deployment using Go

This project is a Go application that reads a Kubernetes deployment YAML file from the local file system and deploys it to a Kubernetes cluster using the official Kubernetes Go client (`client-go`). The Kubernetes cluster is accessed via a kubeconfig file, which is typically located at `~/.kube/config`.

## Prerequisites

Before you begin, ensure you have the following:
- **Go** installed on your system (version 1.16 or higher).
- **Kubernetes cluster** access set up (e.g., Minikube, GKE, AKS, or EKS).
- A **kubeconfig** file that provides the credentials to connect to your Kubernetes cluster.
- **GOPATH** set up and required Go dependencies installed.

## Project Setup

### 1. Install Dependencies

Install the necessary Go packages to interact with the Kubernetes API and work with YAML files. Run the following commands in your terminal:

```bash
go get k8s.io/client-go@v0.23.0
go get k8s.io/api@v0.23.0
go get k8s.io/apimachinery@v0.23.0
```

### 2. Prepare the Deployment YAML

Ensure you have a Kubernetes deployment YAML file (`deployment.yaml`) in the project directory or at your preferred path. Here’s an example of a simple deployment YAML:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: my-deployment
  namespace: default
spec:
  replicas: 2
  selector:
    matchLabels:
      app: my-app
  template:
    metadata:
      labels:
        app: my-app
    spec:
      containers:
      - name: my-container
        image: nginx
        ports:
        - containerPort: 80
```

You can modify this file according to your needs (e.g., image, namespace, replicas).

### 3. Ensure Kubeconfig is Configured

Ensure your Kubernetes credentials are available in the default kubeconfig file (`~/.kube/config`) or update the code to point to your kubeconfig path.

### 4. Run the Application

To run the Go application, execute the following command in your terminal:

```bash
go run main.go
```

This will:
1. Read the `deployment.yaml` file.
2. Parse the YAML and convert it into a Kubernetes `Deployment` object.
3. Use your `kubeconfig` to authenticate with your Kubernetes cluster.
4. Deploy the defined deployment to your Kubernetes cluster.

## How It Works

- The Go code reads the Kubernetes deployment YAML file using `ioutil.ReadFile`.
- The YAML is unmarshalled into a Kubernetes `Deployment`.
- It then uses the Kubernetes client-go package to connect to the Kubernetes cluster via kubeconfig and deploys the parsed deployment object using `clientset.AppsV1().Deployments().Create()`.
- The successful creation of the deployment is logged to the console.



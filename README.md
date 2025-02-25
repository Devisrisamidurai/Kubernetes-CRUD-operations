
# Kubernetes CRUD Operations using Golang

## Overview
This project demonstrates how to perform CRUD (Create, Read, Update, Delete) operations on Kubernetes Pods using Golang.

## Prerequisites
- Go installed (`>=1.18`)
- Kubernetes cluster (Minikube, Kind, or any managed cluster)
- `kubectl` configured to interact with the cluster
- Client-go library installed

## Setup Instructions
1. Clone the repository:
   ```sh
   git clone https://github.com/Devisrisamidurai/Kubernetes-CRUD-operations.git
   cd Kubernetes-CRUD-operations
   ```

2. Install dependencies:
   ```sh
   go mod tidy
   ```

3. Ensure your Kubernetes cluster is running:
   ```sh
   kubectl get nodes
   ```

## Running CRUD Operations

### 1. List Pods (Read Operation)
To list all pods in the cluster, run:
   ```sh
   go run readPods.go
   ```
#### Expected Output:
   ```sh
   There are 5 pods in the cluster
   Name of 0th pod: demo-crud55wwk
   Name of 1th pod: demo-nginx
   Name of 2th pod: go-api-2mwpl
   Name of 3th pod: test-deploy-859f95ffcc-8p8t8
   Name of 4th pod: test-deploy-859f95ffcc-fcdld
   ```

### 2. Create a Pod
To create a new pod, run:
   ```sh
   go run createPod.go
   ```
#### Expected Output:
   ```sh
   Pod my-pod created successfully!
   ```
Verify the pod:
   ```sh
   kubectl get pods
   ```

### 3. Update a Pod
To update an existing pod:
   ```sh
   go run update_existing_Pod.go
   ```
#### Expected Output:
   ```sh
   Pod my-pod updated successfully!
   ```

### 4. Delete a Pod
To delete a pod:
   ```sh
   go run deletePod.go
   ```
#### Expected Output:
   ```sh
   Pod my-pod deleted successfully!
   ```
Verify deletion:
   ```sh
   kubectl get pods
   ```

## Conclusion
This project provides a simple way to interact with Kubernetes using Golang. You can extend it further to work with Deployments, Services, and ConfigMaps.

Happy Coding! 🚀


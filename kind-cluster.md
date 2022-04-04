#### 1. create a kubernetes cluster

Create a Kubernetes cluster consisting of the master node and some worker nodes. 

```
kind create cluster --name istio-cluster --config kind-config.yaml
```

#### 2.verify the kubernetes cluster.

```
kind get clusters
```

#### 3.connect kubectl to new cluster

```
kubectl cluster-info --context kind-istio-cluster
```

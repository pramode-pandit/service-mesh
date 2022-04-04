####1. Create a Kubernetes Cluster

Create a Kubernetes cluster consisting of the master node and some worker nodes. 

```
kind create cluster --name cassandra-kub-cluster --config kind-config.yaml
```

verify the Kubernetes cluster.

```
kind get clusters
```

connect kubectl to our new cassandra kubernetes cluster.

```
kubectl cluster-info --context kind-cassandra-kub-cluster
```

docker build . -t register --progress=plain

kind load docker-image register

// kubectl run register --image=register --image-pull-policy=Never

// kubectl expose pod register --target-port:8080

kubectl apply -f application.yaml
kubectl apply -f gateway.yaml

kubectl get pods 

kubectl -n istio-system port-forward svc/istio-ingressgateway  80:80

access as > http://myapp.servicemesh.io/register
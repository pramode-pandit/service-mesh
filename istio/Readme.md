#### Documentation
> https://istio.io/latest/docs/
> https://istio.io/latest/docs/setup/getting-started/

#### Step:1 Platform Setup
- Kind

```
#Windows
kind create cluster --name sevice-mesh --image kindest/node:v1.17.0@sha256:9512edae126da271b66b990b6fff768fbb7cd786c7d39e86bdf55906352fdf62

#Linux
kind create cluster --name service-mesh --kubeconfig ~/.kube/kind-vault --image kindest/node:v1.17.0@sha256:9512edae126da271b66b990b6fff768fbb7cd786c7d39e86bdf55906352fdf62
```

#### Step:2 Install
- Install with Istioctl
 
  - Download istioctl for your platform 
    > https://github.com/istio/istio/releases/tag/1.10.0
    
    Windows
    https://github.com/istio/istio/releases/download/1.10.0/istio-1.10.0-win.zip
  
  - Add the istioctl client to your path (Linux or macOS or Windows):
  
  - `istioctl version`
  
     no running Istio pods in "istio-system"
     1.10.0

  - `istioctl profile list`

     The simplest option is to install the default Istio configuration profile.
     > https://istio.io/latest/docs/setup/additional-setup/config-profiles/
     
  - Learn more 
    > https://istio.io/latest/docs/setup/install/istioctl/

  - Istio pre-lifght check to cluster compatbility
  
    `istioctl x precheck`

  - `istioctl install` will install default profile or use profile of your choice `istioctl install --set profile=demo -y`

  - `kubectl get all -n istio-system`

  - `kubectl get svc -n istio-system`
  
  - `kubectl label namespace default istio-injection=enabled` Add a namespace label to instruct Istio to automatically inject Envoy sidecar proxies when you deploy our application


#### Step:3 Deploy the sample application

- Bookinfo Application
  > https://istio.io/latest/docs/examples/bookinfo/
  
- `kubectl apply -f https://raw.githubusercontent.com/istio/istio/release-1.10/samples/bookinfo/platform/kube/bookinfo.yaml`

- `kubectl get services`

- `kubectl get pods`

- Verify everything is working correctly up to this point

  `kubectl exec "$(kubectl get pod -l app=ratings -o jsonpath='{.items[0].metadata.name}')" -c ratings -- curl -sS productpage:9080/productpage | grep -o "<title>.*</title>"<title>Simple Bookstore App</title>`

#### Step:4 Open the application to outside traffic

- Associate this application with the Istio gateway: 

  `kubectl apply -f https://raw.githubusercontent.com/istio/istio/release-1.10/samples/bookinfo/networking/bookinfo-gateway.yaml`

  `kubectl get gateway`

  `kubectl get virtualservice`

- Ensure that there are no issues with the configuration:

  `istioctl analyze`

#### Step:5 Determining the ingress IP and ports

 - `kubectl get svc istio-ingressgateway -n istio-system`

 -  Using kind cluster, so NO provion of using loadbalancer. We will port forward the ingress gateway pod to route the request from our local machine to the pod directly
 
 - `kubectl get pod  -n istio-system`
 
 - `kubectl -n istio-system port-forward svc/istio-ingressgateway  80:80`

 - Add  a fake domain servicemesh to /etc/hosts
 
   `127.0.0.1  servicemesh.io`

#### Step:6 Verify external access

- http://servicemesh.io/productpage

- http://servicemesh.io/api/v1/products

#### Step:7 Telemetry Addons

- This directory contains sample deployments of various addons that integrate with Istio. While these applications are not a part of Istio, they are essential to making the most of Istio's observability features.
  
  >https://github.com/istio/istio/tree/release-1.10/samples/addons

- Use the following instructions to deploy the Kiali dashboard, along with Prometheus, Grafana, and Jaeger.

  `kubectl apply -f release-1.10/samples/addons`

  `kubectl get pods -n istio-system`

- Access the Kiali dashboard

  `istioctl dashboard kiali`


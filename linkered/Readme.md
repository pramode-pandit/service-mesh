#### Documentation

https://linkerd.io/2.10/getting-started/

#### Step 1: Install the CLI

To install the CLI manually, run:

`curl -sL https://run.linkerd.io/install | sh`


From realese page
 
> https://github.com/linkerd/linkerd2/releases/

> https://github.com/linkerd/linkerd2/releases/download/edge-21.5.3/linkerd2-cli-edge-21.5.3-windows.exe

`linkerd version`

You should see the CLI version, and also Server version: unavailable. This is because you haven’t installed the control plane on your cluster. Don’t worry—we’ll fix that soon enough.

#### Step 2: Validate your Kubernetes cluster

linkerd check --pre

#### Step 3: Install the control plane onto your cluster

`linkerd install | kubectl apply -f -`

In this command, the linkerd install command generates a Kubernetes manifest with all the necessary control plane resources. Piping this manifest into kubectl apply then instructs Kubernetes to add those resources to your cluster.

`linkerd check`

Verify the control plane readiness

#### Step 4: Install add-on to the linkerd

`linkerd viz install | kubectl apply -f -` # on-cluster metrics stack

Extensions add non-critical but often useful functionality to Linkerd. For this guide, we need the viz extension, which will install Prometheus, dashboard, and metrics components onto the cluster

Optionally, at this point you can install other extensions. For example:

```
## optional
linkerd jaeger install | kubectl apply -f - # Jaeger collector and UI
linkerd multicluster install | kubectl apply -f - # multi-cluster components
```

`linkerd check`

Once you’ve installed the viz extension and any other extensions you’d like, we’ll validate everything again

#### Step 5: Explore Linkerd

`linkerd viz dashboard &`

#### Step 6: Install the demo app

kubectl apply -f https://run.linkerd.io/emojivoto.yml 

kubectl -n emojivoto port-forward svc/web-svc 8080:80

kubectl get -n emojivoto deploy -o yaml | linkerd inject - | kubectl apply -f -

linkerd -n emojivoto check --proxy

#### Step 7: Watch it run

linkerd -n emojivoto viz stat deploy
linkerd -n emojivoto viz top deploy
linkerd -n emojivoto viz tap deploy/web

    


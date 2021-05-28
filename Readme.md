#### What is Service mesh ?

A service mesh, is a way to control how different parts of an application share data with one another. 

service mesh is a dedicated infrastructure layer built right into an app for adding observability, security, and reliability features at the platform layer rather than the application layer.

Service mesh allows you to separate the business logic of the application from observability, and network and security policies. It allows you to **connect, secure, and monitor** your microservices.

- Connect: It enables intelligent routing to control the flow of traffic and API calls between services/endpoints. These also enable advanced deployment strategies such as blue/green, canaries or rolling upgrades, and more.

- Secure: It can enforce policies to allow or deny communication.It ca help with Authenticatio and authorosation of the requests before being served by the app. E.g. you can configure a policy to deny access to production services from a client service running in development environment.

- Monitor: Service Mesh often integrates out-of-the-box with monitoring and tracing tools (such as Prometheus and Jaeger in the case of Kubernetes) to allow you to discover and visualize dependencies between services, traffic flow, API latencies, and tracing.


#### What are Service Mesh Options for Kubernetes ?

- **Hashicorp Consul** : Consul Connect uses an agent installed on every node as a DaemonSet which communicates with the Envoy sidecar proxies that handles routing & forwarding of traffic.

- **Istio** : Istio has separated its data and control planes by using a sidecar loaded proxy which caches information so that it does not need to go back to the control plane for every call. 

- **Linkerd** : Linkerd is probably the second most popular service mesh on Kubernetes and, due to its rewrite in v2, its architecture mirrors Istio’s closely, with an initial focus on simplicity instead of flexibility. This fact, along with it being a Kubernetes-only solution, results in fewer moving pieces, which means that Linkerd has less complexity overall.



#### How can we comapre Istio, Linkerd and Consul ?

https://platform9.com/blog/kubernetes-service-mesh-a-comparison-of-istio-linkerd-and-consul/




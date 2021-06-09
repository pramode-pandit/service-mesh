docker build . -t register --progress=plain

kind load docker-image register

// kubectl run register --image=register --image-pull-policy=Never

// kubectl expose pod register --target-port:8080

kubectl apply -f application.yaml

kubectl apply -f gateway.yaml

kubectl get pods 

kubectl -n istio-system port-forward svc/istio-ingressgateway  80:80

Add a fake domain entry in /etc/hosts

`localhost myapp.servicemesh.io`

Access the app > http://myapp.servicemesh.io/

-----------------------------------

#### Application : vaccine program

**components**

auth   

- authenticate_user
- create_token
- validate_token
- receives GET /auth?user=1234567890&token=1234123412341234
- 
- receives GET /auth?mobile=1234567890
- stores data redis SET key family
- performs auth function()
- Success -> return token Failure -> return error             | 

register
- add user details of a family 
- receives POST /register?mobile=1234567890
- stores data redis to a hashmap key family:1234567890  
- performs regsiter function()
- Success -> return true -> return error 

profile
- show family user and vaccination detail
- receives GET /profile?id=1234567890
- sends hashmap key family:1234567890 

vaccinator
- update vaccine details 
- receives POST /vaccines?id=1234567890&&cowinid=09876
- updated hashmap key family:1234567890 

passport
- download a report
- receives GET /passport?id=1234567890&&cowinid=09876


**start database in container**

```
docker network create jabaap

docker volume create auth
docker volume create profile

docker run -d --rm --name auth --net redis -v auth:/data/  -p 6379:6379 redis:6.0-alpine redis-server 
docker run -d --rm --name profile --net redis -v profile:/data/  -p 6379:6379 redis:6.0-alpine redis-server 

```

**test component auth**

```
cd application\auth

go mod init example.com/auth

go run .

http://localhost:4000/auth?id=1234567890
```

**test component register**

```
cd application\auth

go mod init example.com/auth

go run .

http://localhost:4000/register?id=1234567890
```

**test component profile**

```
cd application\profile

go mod init example.com/profile

go run .

http://localhost:4000/profile?id=1234567890
```

**test component vaccinator**

```
cd application\vaccinator

go mod init example.com/vaccinator

go run .

http://localhost:4000/showMember?id=1234567890&member=2
http://localhost:4000/updateMember?id=1234567890&member=2
```


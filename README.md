# k8s-test-app

study for kubernetes

# Before start

- [Install kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/)
- [Install minikube](https://kubernetes.io/ja/docs/tasks/tools/install-minikube/)

# Run minikube

start

```
$ minikube start
```

check

```
$ minikube status
host: Running
kubelet: Running
apiserver: Running
kubectl: Correctly Configured: pointing to minikube-vm at 192.168.99.100
```

set minikube docker-env

```
$ eval $(minikube docker-env)
```

# Build docker

```
$ docker build -t simple-app:0.0.2 .
```

# Create deployment

```
$ kubectl create -f k8s/minikube/app-deployment.yml
$ kubectl create -f k8s/minikube/db-deployment.yml
```

check

```
$ kubectl get po
NAME                   READY   STATUS    RESTARTS   AGE
app-57c65b4fcc-kndxk   1/1     Running   0          5m56s
db-96cc996d4-wbtqk     1/1     Running   0          45m

$ kubectl get deployment
NAME   READY   UP-TO-DATE   AVAILABLE   AGE
app    1/1     1            1           8m9s
db     1/1     1            1           48m

```

# Create service

```
$ kubectl create -f k8s/minikube/app-service.yml
$ kubectl create -f k8s/minikube/db-service.yml
```

check

```
$ kubectl get service
NAME         TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)        AGE
app          NodePort    10.110.33.109    <none>        80:30765/TCP   70m
db           ClusterIP   10.106.253.125   <none>        3306/TCP       45m
kubernetes   ClusterIP   10.96.0.1        <none>        443/TCP        36h
```

# Set DB data

```
$ kubectl exec -it db-xxxxxxxxxx-xxxxx /bin/sh`
$ mysql -uroot -ppassword
$ CREATE DATABASE app;
$ USE app;
$ CREATE TABLE user (id int, name varchar(10));
$ INSERT INTO user (id, name) VALUES (1, "hoge");
$ INSERT INTO user (id, name) VALUES (2, "fuga");
```

# Access Server

get minikube service url

```
$ minikube service app --url
```

access app server (change ip:port what you get above)

```
$ curl http://192.168.99.100:30765/hello
1: hoge
2: fuga
```

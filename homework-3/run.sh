# prepare
k config use-context minikube

# create deployment
k create -f deployment.yml
k delete -f deployment.yml
k apply -f deployment.yml


curl -I 172.17.0.3/healthz

curl 127.0.0.1:31292/logreq
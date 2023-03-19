
docker build -t emanpicar/kubernetes-gkube-api:0.0.1 .
docker run --name gkube-api -d -p 8080:8080 emanpicar/kubernetes-gkube-api:0.0.1

kubectl apply -f yaml/deployment/gkube-api-deployment.yaml
kubectl port-forward deployments/gkube-api-deployment 8080:8080 -n playground-pub
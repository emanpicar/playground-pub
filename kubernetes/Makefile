gkubeApiVersion := 0.0.2

docker.build:
	docker build -t emanpicar/kubernetes-gkube-api:$(gkubeApiVersion) ./gkube-api/.

docker.push:
	docker push emanpicar/kubernetes-gkube-api:$(gkubeApiVersion)

docker.run:
	docker run --name gkube-api -d -p 8080:8080 emanpicar/kubernetes-gkube-api:$(gkubeApiVersion)
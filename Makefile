PROJECT_ROOT := $(abspath $(dir $(lastword $(MAKEFILE_LIST))))

# Todo:build all image
install-cluster:
	kind delete cluster --name douyin
	kind create cluster --config deployment/cluster/douyin-cluster.yaml

# Deploy MinIO Service
install-minio:
	-kubectl delete ns minio
	helm repo add minio https://charts.min.io/
	kubectl create ns minio
	helm install minio minio/minio -f deployment/minio/minio.yaml -n minio

# Deploy Etcd service, you must install it before use rpc
install-etcd:
	-kubectl delete ns etcd
	helm repo add etcd https://charts.bitnami.com/bitnami
	kubectl create ns etcd
	helm install etcd etcd/etcd --set replicaCount=2 -n etcd

# Forward MinIO's console service(web)
forward-minio-console:
	@echo Forwarding MinIO console service, visit it from http://localhost:9001
	kubectl port-forward -n minio svc/minio-console 9001:9001

# Demo pkg
buildx-demo:
	docker buildx build --platform=linux/arm64 -f ${PROJECT_ROOT}/cmd/demo/Dockerfile \
		--build-arg PROJECT_ROOT="${PROJECT_ROOT}" ${PROJECT_ROOT} \
		-t douyin/demo:nightly

install-demo:
	-kubectl delete ns demo
	kind load docker-image douyin/demo:nightly --name douyin
	-kubectl create ns demo
	kubectl apply -f deployment/demo/demo.yaml

forward-demo:
	kubectl port-forward -n demo service/demo 8000:80

# Userinfo-rpc/api Demo
build-userinfo-demo:
	docker build -f ${PROJECT_ROOT}/cmd/userinfo-demo/rpc/Dockerfile \
		--build-arg PROJECT_ROOT="${PROJECT_ROOT}" ${PROJECT_ROOT} \
		-t douyin/userinfo-demo-rpc:nightly
	docker build -f ${PROJECT_ROOT}/cmd/userinfo-demo/api/Dockerfile \
    	--build-arg PROJECT_ROOT="${PROJECT_ROOT}" ${PROJECT_ROOT} \
    	-t douyin/userinfo-demo-api:nightly

install-userinfo-demo:
	-kubectl delete ns userinfo-demo
	kind load docker-image douyin/userinfo-demo-api:nightly --name douyin
	kind load docker-image douyin/userinfo-demo-rpc:nightly --name douyin
	-kubectl create ns userinfo-demo
	kubectl apply -f deployment/userinfo-demo/userinfo-demo.yaml

forward-userinfo-demo:
	kubectl port-forward -n userinfo-demo svc/userinfo-demo 30001:8888
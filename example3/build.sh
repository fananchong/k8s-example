#!/bin/bash

set -ex

if [ ! -x "bin" ]; then
    mkdir "bin"
fi

docker run -it --rm -e GOBIN=/go/bin/ -v "$PWD"/bin:/go/bin/ -v "$PWD"/../:/go/src/ -w /go/src/example3 golang go install ./service_a.go ./k8s_api.go
docker run -it --rm -e GOBIN=/go/bin/ -v "$PWD"/bin:/go/bin/ -v "$PWD"/../:/go/src/ -w /go/src/example3 golang go install ./service_b.go ./k8s_api.go

docker build -t k8s-example3 .

docker tag k8s-example3:latest fananchong/k8s-example3:latest


set +ex

docker push fananchong/k8s-example3:latest

kubectl apply -f ./namespace.yaml
kubectl apply -f ./service_account.yaml
kubectl delete -f ./service_a.yaml
kubectl create -f ./service_a.yaml
kubectl delete -f ./service_b.yaml
kubectl create -f ./service_b.yaml


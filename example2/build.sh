#!/bin/bash

set -ex

if [ ! -x "bin" ]; then
    mkdir "bin"
fi

docker run -it --rm -e GOBIN=/go/bin/ -v "$PWD"/bin:/go/bin/ -v "$PWD"/../:/go/src/ -w /go/src/example2 golang go install ./service_a.go
docker run -it --rm -e GOBIN=/go/bin/ -v "$PWD"/bin:/go/bin/ -v "$PWD"/../:/go/src/ -w /go/src/example2 golang go install ./service_b.go

docker build -t k8s-example2 .

docker tag k8s-example2:latest fananchong/k8s-example2:latest

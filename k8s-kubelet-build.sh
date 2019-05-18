#!/bin/bash

rm -rf ./out/
docker rmi huanwei/k8s-kubelet-build
docker build -t huanwei/k8s-kubelet-build .
docker run -d --name=k8s-kubelet-build huanwei/k8s-kubelet-build sleep 10
mkdir -p ./out/
docker cp k8s-kubelet-build:/out/kubelet ./out/kubelet
docker rm -f k8s-kubelet-build
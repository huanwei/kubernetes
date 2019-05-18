# Copyright 2017 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# huanwei/k8s-kubelet-build
FROM golang
MAINTAINER Huan Wei<huan@harmonycloud.cn>

RUN mkdir -p $GOPATH/src/k8s.io/kubernetes
COPY . $GOPATH/src/k8s.io/kubernetes/
WORKDIR $GOPATH/src/k8s.io/kubernetes
RUN go build -v -i -o /out/kubelet ./cmd/kubelet
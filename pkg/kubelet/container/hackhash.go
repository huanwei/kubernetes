/*
Copyright 2014 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package container

import (
	"hash/fnv"
	"regexp"

	"k8s.io/api/core/v1"
	hashutil "k8s.io/kubernetes/pkg/util/hash"
)

const (
	podVersionLabel    = "k8s.harmonycloud.cn/pod-version"
	//podVersionLabel17  = "1.7"
	//podVersionLabel18  = "1.8"
	podVersionLabel19  = "1.9"
	//podVersionLabel110 = "1.10"
	//podVersionLabel111 = "1.11"
	//podVersionLabel112 = "1.12"
	//podVersionLabel113 = "1.13"
)

// Specific func to hack container hash
func hashContainerExt(container *v1.Container, hackFunc hashutil.HackFunc) uint64 {
	hash := fnv.New32a()
	hashutil.DeepHashObjectExt(hash, *container, hackFunc)
	return uint64(hash.Sum32())
}

// For container hash compatibility when upgrade cluster from old version.
// Use different hash method by pod version annotations
func HashContainerByPodVersion(pod *v1.Pod, container *v1.Container) uint64 {
	var containerHash uint64
	version, exists := pod.Annotations[podVersionLabel]
	/*if exists && version == podVersionLabel17 {
		containerHash = hashContainerExt(container, hackContainerGoStringTo17)
	} else if exists && version == podVersionLabel18 {
		containerHash = hashContainerExt(container, hackContainerGoStringTo18)
	} else if exists && version == podVersionLabel19 {
		containerHash = hashContainerExt(container, hackContainerGoStringTo19)
	} else if exists && version == podVersionLabel110 {
		containerHash = hashContainerExt(container, hackContainerGoStringTo110)
	} else {
		containerHash = HashContainer(container)
	}*/

	if exists && version == podVersionLabel19 {
		containerHash = hashContainerExt(container, hackContainerGoStringTo19)
	} else {
		containerHash = HashContainer(container)
	}
	return containerHash
}

// Convert container hash from 1.8 to 1.7
// xref: git diff v1.7.0 v1.8.0 staging/src/k8s.io/api/core/v1/types.go
func hackGoString18to17(s string) string {
	// MountPropagation:(*v1.MountPropagationMode)<nil>
	re := regexp.MustCompile(`\s*MountPropagation:(.*?)([\s}])`)
	s = re.ReplaceAllString(s, `${2}`)
	// AllowPrivilegeEscalation:(*bool)<nil>
	re = regexp.MustCompile(`\s*AllowPrivilegeEscalation:(.*?)([\s}])`)
	s = re.ReplaceAllString(s, `${2}`)
	return s
}

// Convert container hash from 1.9 to 1.8
// xref: git diff v1.8.0 v1.9.0 staging/src/k8s.io/api/core/v1/types.go
func hackGoString19to18(s string) string {
	// VolumeDevices:([]v1.VolumeDevice)<nil>
	re := regexp.MustCompile(`\s*VolumeDevices:(.*?)([\s}])`)
	s = re.ReplaceAllString(s, `${2}`)
	return s
}

// Convert container hash from 1.10 to 1.9
// xref: git diff v1.9.0 v1.10.0 staging/src/k8s.io/api/core/v1/types.go
func hackGoString110to19(s string) string {
	re := regexp.MustCompile(`\s*RunAsGroup:(.*?)([\s}])`)
	s = re.ReplaceAllString(s, `${2}`)
	return s
}

// Convert container hash from 1.11 to 1.10
// xref: git diff v1.10.0 v1.11.0 staging/src/k8s.io/api/core/v1/types.go
func hackGoString111to110(s string) string {
	return s
}

func hackGoString112to111(s string) string {
	re := regexp.MustCompile(`\s*ProcMount:(.*?)([\s}])`)
	s = re.ReplaceAllString(s, `${2}`)
	return s
}

func hackGoString113to112(s string) string {
	return s
}

func hackGoString114to113(s string) string {
	re := regexp.MustCompile(`\s*SubPathExpr:(.*?)([\s}])`)
	s = re.ReplaceAllString(s, `${2}`)
	return s
}

func hackContainerGoStringTo17(s string) string {
	s = hackGoString114to113(s)
	s = hackGoString113to112(s)
	s = hackGoString112to111(s)
	s = hackGoString111to110(s)
	s = hackGoString110to19(s)
	s = hackGoString19to18(s)
	s = hackGoString18to17(s)
	return s
}

func hackContainerGoStringTo18(s string) string {
	s = hackGoString114to113(s)
	s = hackGoString113to112(s)
	s = hackGoString112to111(s)
	s = hackGoString111to110(s)
	s = hackGoString110to19(s)
	s = hackGoString19to18(s)
	return s
}

func hackContainerGoStringTo19(s string) string {
	//s = hackGoString114to113(s)
	//s = hackGoString113to112(s)
	//s = hackGoString112to111(s)
	s = hackGoString111to110(s)
	s = hackGoString110to19(s)
	return s
}

func hackContainerGoStringTo110(s string) string {
	//s = hackGoString114to113(s)
	//s = hackGoString113to112(s)
	//s = hackGoString112to111(s)
	s = hackGoString111to110(s)
	return s
}
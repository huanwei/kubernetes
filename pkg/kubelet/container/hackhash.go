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

// Specific func to hack container hash
func hashContainerExt(container *v1.Container, hackFunc hashutil.HackFunc) uint64 {
	hash := fnv.New32a()
	hashutil.DeepHashObjectExt(hash, *container, hackFunc)
	return uint64(hash.Sum32())
}

// For container hash compatibility when upgrade cluster from 1.9
func HashContainerTo19(container *v1.Container) uint64 {
	return hashContainerExt(container, hackContainerGoStringTo19)
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


func hackContainerGoStringTo19(s string) string {
	s = hackGoString112to111(s)
	s = hackGoString111to110(s)
	s = hackGoString110to19(s)
	return s
}
/*
Copyright 2022.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// EtcdBenchmarkSpec defines the desired state of EtcdBenchmark
type EtcdBenchmarkSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of EtcdBenchmark. Edit etcdbenchmark_types.go to remove/update
	Backend  EtcdBenchmarkBackend `json:"backend"`
	Replicas int32                `json:"replicas"`
}

type EtcdBenchmarkBackend struct {
	Type           string                   `json:"type"`
	PodBackend     *EtcdBenchmarkPodBackend `json:"podBackend,omitempty"`
	CounterBackend *EtcdCounterBackend      `json:"counterBackend,omitempty"`
}

type EtcdBenchmarkPodBackend struct {
	Image   string   `json:"image"`
	Command []string `json:"command,omitempty"`
}

type EtcdCounterBackend struct {
}

// EtcdBenchmarkStatus defines the observed state of EtcdBenchmark
type EtcdBenchmarkStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// EtcdBenchmark is the Schema for the etcdbenchmarks API
type EtcdBenchmark struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EtcdBenchmarkSpec   `json:"spec,omitempty"`
	Status EtcdBenchmarkStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// EtcdBenchmarkList contains a list of EtcdBenchmark
type EtcdBenchmarkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EtcdBenchmark `json:"items"`
}

func init() {
	SchemeBuilder.Register(&EtcdBenchmark{}, &EtcdBenchmarkList{})
}

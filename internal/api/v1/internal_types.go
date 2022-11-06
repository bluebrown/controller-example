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

// internal implements the fielder interface
// so that we can pass it to the shared logic
func (i Internal) GetFoo() string {
	return i.Spec.Foo
}

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// InternalSpec defines the desired state of Internal
type InternalSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of Internal. Edit internal_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// InternalStatus defines the observed state of Internal
type InternalStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Internal is the Schema for the internals API
type Internal struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   InternalSpec   `json:"spec,omitempty"`
	Status InternalStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// InternalList contains a list of Internal
type InternalList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Internal `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Internal{}, &InternalList{})
}

/*
Copyright 2023.

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

// GraphicsCardSpec defines the desired state of GraphicsCard
type GraphicsCardSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Name of the model
	// +kubebuilder:validation:Enum="3060";"3060TI";"3070";"3070TI"
	Model string `json:"model,omitempty"`

	// Something I pickup online
	// +kubebuilder:validation:Minimum=1000
	// +kubebuilder:validation:Maximum=8000
	CudaCores int32 `json:"cuda_cores,omitempty"`

	// Boost Clock in GHz
	// +kubebuilder:validation:MaxLength=4
	// +kubebuilder:validation:MinLength=1
	BoostClock string `json:"boost_clock,omitempty"`

	// Memory in GB (e.g 8GB, 12GB)
	// +kubebuilder:validation:Required
	MemorySize string `json:"memory_size,omitempty"`

	// Memory type
	// +kubebuilder:validation:Enum=GDDR6;GDDR6X
	// +kubebuilder:validation:Optional
	MemoryType string `json:"memory_type,omitempty"`
}

// GraphicsCardStatus defines the observed state of GraphicsCard
type GraphicsCardStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// PodName of the active GraphicsCard node.
	// +optional
	Active string `json:"active"`

	// PodNames of the standby GraphicsCard nodes.
	// +optional
	Standby []string `json:"standby"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// GraphicsCard is the Schema for the graphicscards API
type GraphicsCard struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GraphicsCardSpec   `json:"spec,omitempty"`
	Status GraphicsCardStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// GraphicsCardList contains a list of GraphicsCard
type GraphicsCardList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GraphicsCard `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GraphicsCard{}, &GraphicsCardList{})
}

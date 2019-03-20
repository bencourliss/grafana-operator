/*
Copyright 2019 Jagadish Nagarajaiah.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// GrafanaSpec defines the desired state of Grafana
type GrafanaSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// +optional
	Replicas *int32 `json:"replicas,omitempty"`

	// +optional
	Image string `json:"image,omitempty"`

	// +optional
	GrafanaAdminUser string `json:"grafana_admin_user,omitempty"`

	// +optional
	GrafanaAdminPassword string `json:"grafana_admin_password,omitempty"`

	//+optional
	PrometheusUrl string `json:"prometheus_url,omitempty"`
}

// GrafanaStatus defines the observed state of Grafana
type GrafanaStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// ActiveReplicaCount is the number of grafana replicas
	ActiveGrafanaCount int32 `json:"active_grafana_count"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Grafana is the Schema for the grafanas API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
type Grafana struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GrafanaSpec   `json:"spec,omitempty"`
	Status GrafanaStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GrafanaList contains a list of Grafana
type GrafanaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Grafana `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Grafana{}, &GrafanaList{})
}

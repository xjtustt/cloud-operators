/*
 * Copyright 2020 IBM Corporation
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ServiceSpec defines the desired state of Service
type ServiceSpec struct {
	// ServiceClass is the name of the service from the IBM Cloud Catalog
	ServiceClass string `json:"serviceClass"`
	// Plan for the service from the IBM Cloud Catalog
	Plan string `json:"plan"`
	// ServiceClassType is set to CF if the service is CloundFoundry
	// +optional
	ServiceClassType string `json:"serviceClassType,omitempty"`
	// ExternalName is the name for the service as it appears on IBM Cloud
	// +optional
	ExternalName string `json:"externalName,omitempty"`
	// Parameters pass configuration to the service during creation
	// +optional
	Parameters []Param `json:"parameters,omitempty"`
	// +optional
	Tags []string `json:"tags,omitempty"`
	// +optional
	Context ResourceContext `json:"context,omitempty"`
}

// ServiceStatus defines the observed state of Service
type ServiceStatus struct {
	// State is a short name for the current status
	State string `json:"state,omitempty"`
	// Message is a detailed message on current status
	Message string `json:"message,omitempty"`

	Generation int64 `json:"generation,omitempty"`

	// ServiceClass is the name of the service from the IBM Cloud Catalog
	ServiceClass string `json:"serviceClass"`
	// ServiceClassType is set to CF if the service is CloundFoundry
	ServiceClassType string `json:"serviceClassType"`
	// Plan for the service from the IBM Cloud Catalog
	Plan string `json:"plan"`
	// InstanceID is the instance ID for the service
	// +optional
	InstanceID string `json:"instanceId,omitempty"`
	// ExternalName is the name for the service as it appears on IBM Cloud
	// +optional
	ExternalName string `json:"externalName,omitempty"`
	// +optional
	Context ResourceContext `json:"context,omitempty"`
	// Parameters pass configuration to the service during creation
	// +optional
	Parameters []Param `json:"parameters,omitempty"`
	// +optional
	Tags []string `json:"tags,omitempty"`
	// DashboardURL is the dashboard URL for the service
	// +optional
	DashboardURL string `json:"dashboardURL,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:printcolumn:name="Status",type="string",JSONPath=".status.state"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Service is an instance of a Service resource on IBM Cloud.
type Service struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ServiceSpec   `json:"spec,omitempty"`
	Status ServiceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceList contains a list of Service
type ServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Service `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Service{}, &ServiceList{})
}

type ServiceContext struct {
	// +optional
	Org string `json:"org,omitempty"`
	// +optional
	Space string `json:"space,omitempty"`
	// +optional
	Region string `json:"region,omitempty"`
	// +optional
	ResourceGroup string `json:"resourcegroup,omitempty"`
	// +optional
	ResourceGroupID string `json:"resourcegroupid,omitempty"`
	// +optional
	ResourceLocation string `json:"resourcelocation,omitempty"`
	// +optional
	User string `json:"user,omitempty"`
}

// ResourceContext defines the CloudFoundry context and resource group
type ResourceContext struct {
	// +optional
	Org string `json:"org,omitempty"`
	// +optional
	Space string `json:"space,omitempty"`
	// +optional
	Region string `json:"region,omitempty"`
	// +optional
	ResourceGroupID string `json:"resourcegroupid,omitempty"`
	// +optional
	ResourceLocation string `json:"resourcelocation,omitempty"`
	// +optional
	User string `json:"user,omitempty"`
}

func (s *Service) GetState() string {
	return s.Status.State
}

func (s *Service) GetMessage() string {
	return s.Status.Message
}

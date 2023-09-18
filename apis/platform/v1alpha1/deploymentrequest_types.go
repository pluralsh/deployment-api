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

package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

func init() {
	SchemeBuilder.Register(&DeploymentRequest{}, &DeploymentRequestList{})
}

type DeploymentRequestSpec struct {
	// DeploymentClassName name of the DeploymentClass
	DeploymentClassName string `json:"deploymentClassName,omitempty"`

	// Name of a deployment object.
	// If unspecified, then a new Deployment will be dynamically provisioned
	// +optional
	ExistingDeploymentName string `json:"existingBucketName,omitempty"`
}

type DeploymentRequestStatus struct {
	// Ready is true when the provider resource is ready.
	// +optional
	Ready bool `json:"ready"`

	// DeploymentName is the name of the provisioned Deployment in response to this DeploymentRequest.
	// +optional
	DeploymentName string `json:"deploymentName,omitempty"`
}

// DeploymentRequest represents single request for Deployment.
// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.ready",description="Deployment ready status"
// +kubebuilder:printcolumn:name="Deployment",type="string",JSONPath=".status.deploymentName",description="Deployment name"
type DeploymentRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DeploymentRequestSpec   `json:"spec,omitempty"`
	Status DeploymentRequestStatus `json:"status,omitempty"`
}

// DeploymentRequestList is a list of DeploymentRequest.
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type DeploymentRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DeploymentRequest `json:"items"`
}

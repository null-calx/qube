/*
Copyright 2025.

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

// TorrentSpec defines the desired state of Torrent
type TorrentSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	// The following markers will use OpenAPI v3 schema to validate the value
	// More info: https://book.kubebuilder.io/reference/markers/crd-validation.html

	// foo is an example field of Torrent. Edit torrent_types.go to remove/update
	// +optional
	TorrentURL                *string `json:"torrent_url"`
	Upload                    *bool   `json:"upload"`
	PersistentVolumeClaimName *string `json:"persistent_volume_claim_name"`
}

// TorrentStatus defines the observed state of Torrent.
type TorrentStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Peers *int64 `json:"peers"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Torrent is the Schema for the torrents API
type Torrent struct {
	metav1.TypeMeta `json:",inline"`

	// metadata is a standard object metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty,omitzero"`

	// spec defines the desired state of Torrent
	// +required
	Spec TorrentSpec `json:"spec"`

	// status defines the observed state of Torrent
	// +optional
	Status TorrentStatus `json:"status,omitempty,omitzero"`
}

// +kubebuilder:object:root=true

// TorrentList contains a list of Torrent
type TorrentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Torrent `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Torrent{}, &TorrentList{})
}

// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ReservedIPSpec defines the desired state of ReservedIP
type ReservedIPSpec struct {
	// +kubebuilder:default=4
	// +kubebuilder:validation:Enum=4;6
	// +kubebuilder:validation:Optional
	IPVersion *int64 `json:"ipVersion,omitempty"`

	// +kubebuilder:validation:Required
	IPs []string `json:"ips"`
}

// +kubebuilder:resource:categories={spiderpool},path="spiderreservedips",scope="Cluster",shortName={sr},singular="spiderreservedip"
// +kubebuilder:printcolumn:JSONPath=".spec.ipVersion",description="ipVersion",name="VERSION",type=string
// +kubebuilder:object:root=true

// ReservedIP is the Schema for the reservedips API
type ReservedIP struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ReservedIPSpec `json:"spec,omitempty"`
}

// +kubebuilder:object:root=true

// ReservedIPList contains a list of ReservedIP
type ReservedIPList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []ReservedIP `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ReservedIP{}, &ReservedIPList{})
}

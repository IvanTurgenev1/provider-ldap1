/*
Copyright 2022 The Crossplane Authors.

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
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// GroupldapParameters are the configurable fields of a Groupldap.
type GroupldapParameters struct {
	ConfigurableField string `json:"configurableField"`
}

// GroupldapObservation are the observable fields of a Groupldap.
type GroupldapObservation struct {
	ObservableField string `json:"observableField,omitempty"`
}

// A GroupldapSpec defines the desired state of a Groupldap.
type GroupldapSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       GroupldapParameters `json:"forProvider"`
}

// A GroupldapStatus represents the observed state of a Groupldap.
type GroupldapStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          GroupldapObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A Groupldap is an example API type.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ldap}
type Groupldap struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GroupldapSpec   `json:"spec"`
	Status GroupldapStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GroupldapList contains a list of Groupldap
type GroupldapList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Groupldap `json:"items"`
}

// Groupldap type metadata.
var (
	GroupldapKind             = reflect.TypeOf(Groupldap{}).Name()
	GroupldapGroupKind        = schema.GroupKind{Group: Group, Kind: GroupldapKind}.String()
	GroupldapKindAPIVersion   = GroupldapKind + "." + SchemeGroupVersion.String()
	GroupldapGroupVersionKind = SchemeGroupVersion.WithKind(GroupldapKind)
)

func init() {
	SchemeBuilder.Register(&Groupldap{}, &GroupldapList{})
}

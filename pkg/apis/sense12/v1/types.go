package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AppServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []AppService `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AppService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              AppServiceSpec   `json:"spec"`
	Status            AppServiceStatus `json:"status,omitempty"`
}

type ContainerPort struct {
	Name string `json:"name"`
	Port int32  `json:"containerPort"`
}

type AppServiceSpec struct {
	// Name of the application
	Name string `json:"name"`
	// Docker image name including tag
	Image string `json:"image"`
	// Main port that the App is exposing
	Ports []ContainerPort `json:"ports"`
}

type AppServiceStatus struct {
	// Latest stable image
	StableImage string `json:"stableImage"`
	// When upgrade is in progress a new image will be launced in Beta mode
	UpgradeInProgress bool `json:"upgradeInProgress"`
}

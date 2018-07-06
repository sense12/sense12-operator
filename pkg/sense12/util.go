package sense12

import (
	api "github.com/sense12/sense12-operator/pkg/apis/sense12/v1"
	corev1 "k8s.io/api/core/v1"
)

// labelsForVault returns the labels for selecting the resources
// belonging to the given appService name.
func specLabels(name string) map[string]string {
	return map[string]string{"app": "sense12", "app_service": name}
}

func specContainerPorts(p []api.ContainerPort) []corev1.ContainerPort {
	var out []corev1.ContainerPort
	for _, port := range p {
		out = append(out, corev1.ContainerPort{
			Name:          port.Name,
			ContainerPort: port.Port,
		})
	}
	return out
}

func specServicePorts(p []api.ContainerPort) []corev1.ServicePort {
	var out []corev1.ServicePort
	for _, port := range p {
		out = append(out, corev1.ServicePort{
			Name: port.Name,
			Port: port.Port,
		})
	}
	return out
}

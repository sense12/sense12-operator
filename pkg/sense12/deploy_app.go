package sense12

import (
	"github.com/operator-framework/operator-sdk/pkg/sdk/action"
	api "github.com/sense12/sense12-operator/pkg/apis/sense12/v1"
	"github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const (
	servicePortName = "http"
)

func createDeployment(cr *api.AppService) error {
	pod := &corev1.Pod{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Pod",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Spec.Name,
			Namespace: cr.Namespace,
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(cr, schema.GroupVersionKind{
					Group:   api.SchemeGroupVersion.Group,
					Version: api.SchemeGroupVersion.Version,
					Kind:    api.AppServiceKind,
				}),
			},
			Labels: specLabels(cr.GetName()),
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:    "app",
					Image:   cr.Spec.Image,
					Command: []string{"sleep", "3600"},
					Ports:   specContainerPorts(cr.Spec.Ports),
				},
			},
		},
	}

	err := action.Create(pod)
	if err != nil && !errors.IsAlreadyExists(err) {
		logrus.Errorf("Failed to create AppService pod : %v", err)
		return err
	}

	svc := &corev1.Service{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Spec.Name,
			Namespace: cr.Namespace,
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(cr, schema.GroupVersionKind{
					Group:   api.SchemeGroupVersion.Group,
					Version: api.SchemeGroupVersion.Version,
					Kind:    "AppService",
				}),
			},
			Labels: specLabels(cr.GetName()),
		},
		Spec: corev1.ServiceSpec{
			Selector: specLabels(cr.GetName()),
			Ports:    specServicePorts(cr.Spec.Ports),
		},
	}
	return action.Create(svc)
}

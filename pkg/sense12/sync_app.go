package sense12

import (
	"fmt"

	api "github.com/sense12/sense12-operator/pkg/apis/sense12/v1"
	"github.com/sirupsen/logrus"

	"github.com/operator-framework/operator-sdk/pkg/sdk/action"
	"github.com/operator-framework/operator-sdk/pkg/sdk/query"
	// corev1 "k8s.io/api/core/v1"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// upgradeAppImage ensures that the pods running upgrade when the AppService is changed
func upgradeAppImage(cr *api.AppService) error {
	p := &corev1.Pod{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Pod",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.GetName(),
			Namespace: cr.GetNamespace(),
		},
	}
	err := query.Get(p)
	if err != nil {
		return fmt.Errorf("failed to get deployment (%s): %v", p.Name, err)
	}

	container := p.Spec.Containers[0]
	if container.Image != cr.Spec.Image {
		// create a new Deployment with new Image
		container.Image = cr.Spec.Image
		// todo: update State of CR

		logrus.Info("Updating Pod with difference in Image")
		err = action.Update(p)
		if err != nil {
			return fmt.Errorf("failed to update size of deployment (%s): %v", p.Name, err)
		}
	} else {
		logrus.WithFields(logrus.Fields{
			"container": container,
			"spec":      cr.Spec,
		}).Info("No ifference in Image")
	}
	return nil
}

package sense12

import (
	api "github.com/sense12/sense12-operator/pkg/apis/sense12/v1"
	"github.com/sirupsen/logrus"
)

func Reconcile(appService *api.AppService) (err error) {

	logrus.Info("Create new pod")
	err = createDeployment(appService)
	if err != nil {
		return err
	}

	err = upgradeAppImage(appService)
	if err != nil {
		return err
	}
	logrus.Info("Finished Reconcile without errors")
	return nil
}

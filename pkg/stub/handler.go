package stub

import (
	"github.com/sense12/sense12-operator/pkg/sense12"

	"github.com/operator-framework/operator-sdk/pkg/sdk/handler"
	"github.com/operator-framework/operator-sdk/pkg/sdk/types"
	api "github.com/sense12/sense12-operator/pkg/apis/sense12/v1"
	"github.com/sirupsen/logrus"
)

func NewHandler() handler.Handler {
	return &Handler{}
}

type Handler struct {
	// Fill me
}

func (h *Handler) Handle(ctx types.Context, event types.Event) error {
	logrus.WithFields(logrus.Fields{
		"event": event,
	}).Info("Handle Event")

	switch o := event.Object.(type) {
	case *api.AppService:
		app := o

		// Ignore delete event
		if event.Deleted {
			logrus.Info("Ignored event.Deleted")
			return nil
		}

		return sense12.Reconcile(app)
	}
	return nil
}

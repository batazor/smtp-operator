package stub

import (
	api "github.com/batazor/smtp-operator/pkg/apis/smtp/v1"

	"github.com/operator-framework/operator-sdk/pkg/sdk/handler"
	"github.com/operator-framework/operator-sdk/pkg/sdk/types"
	"github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/api/errors"
	"github.com/batazor/smtp-operator/pkg/smtp"
)

func NewHandler() handler.Handler {
	return &Handler{}
}

type Handler struct {
	// Fill me
}

func (h *Handler) Handle(ctx types.Context, event types.Event) error {
	switch o := event.Object.(type) {
	case *api.Template:
		err := newTemplate(o)
		if err != nil && !errors.IsAlreadyExists(err) {
			logrus.Errorf("Failed to create busybox pod : %v", err)
			return err
		}
	}
	return nil
}

func newTemplate(cr *api.Template) error {
	err := smtp.NewTemplate(cr)
	if err != nil {
		return err
	}

	return nil
}

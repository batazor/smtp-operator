package smtp

import (
	"github.com/batazor/smtp-operator/pkg/apis/v1"
	"io/ioutil"
	"github.com/sirupsen/logrus"
)

func NewTemplate(cr *v1.SMTPTemplate) error {
	template := []byte(cr.Spec.Template)
	name := cr.ObjectMeta.Name

	err := ioutil.WriteFile("/tmp/" + name, template, 0644)
	if err != nil {
		return err
	}

	logrus.Info("Success create template: %v", name)

	return nil
}
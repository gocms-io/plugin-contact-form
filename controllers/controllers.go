package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gocms-io/plugin-contact-form/controllers/contact-form"
	"github.com/gocms-io/plugin-contact-form/services"
)

type ControllersGroup struct {
	apiControllers *ApiControllers
}

type ApiControllers struct {
	ContactFormController *contact_form_ctrl.ContactFormController
}

func DefaultControllerGroup(r *gin.Engine, sg *services.ServicesGroup) *ControllersGroup {
	controllersGroup := &ControllersGroup{
		apiControllers: &ApiControllers{
			ContactFormController: contact_form_ctrl.DefaultContactFormController(r, sg),
		},
	}

	return controllersGroup
}

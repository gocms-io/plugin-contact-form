package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gocms-io/plugin-contact-form/controllers/api/contact-form"
	"github.com/gocms-io/plugin-contact-form/controllers/content"
	"github.com/gocms-io/plugin-contact-form/services"
)

type ControllersGroup struct {
	apiControllers     *ApiControllers
	contentControllers *ContentControllers
}

type ApiControllers struct {
	ContactFormController *contact_form_ctrl.ContactFormController
}

type ContentControllers struct {
	DocumentationController *content_ctrl.DocumentationController
}

func DefaultControllerGroup(r *gin.Engine, sg *services.ServicesGroup) *ControllersGroup {
	controllersGroup := &ControllersGroup{
		apiControllers: &ApiControllers{
			ContactFormController: contact_form_ctrl.DefaultContactFormController(r, sg),
		},
		contentControllers: &ContentControllers{
			DocumentationController: content_ctrl.DefaultDocumentationController(r),
		},
	}

	return controllersGroup
}

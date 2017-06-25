package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gocms-io/plugin-contact-form/controllers/example"
	"github.com/gocms-io/plugin-contact-form/services"
)

type ControllersGroup struct {
	apiControllers *ApiControllers
}

type ApiControllers struct {
	eventsController *events_ctrl.ExampleController
}

func DefaultControllerGroup(r *gin.Engine, sg *services.ServicesGroup) *ControllersGroup {
	controllersGroup := &ControllersGroup{
		apiControllers: &ApiControllers{
			eventsController: events_ctrl.DefaultExampleController(r, sg),
		},
	}

	return controllersGroup
}

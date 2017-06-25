package events_ctrl

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/gocms-io/plugin-contact-form/services"
)

type ExampleController struct {
	ServiceGroup *services.ServicesGroup
	Routes       *gin.RouterGroup
}

func DefaultExampleController(r *gin.Engine, serviceGroup *services.ServicesGroup) *ExampleController {
	ec := &ExampleController{
		Routes:       r.Group("/api"),
		ServiceGroup: serviceGroup,
	}

	ec.Default()
	return ec
}

func (ec *ExampleController) Default() {
	ec.Routes.GET("/example/:id", ec.getExample)
}

/**
* @api {get} /example Example
* @apiName GetExample
* @apiGroup Example
* @apiUse ExampleDisplay
* @apiDescription Get an example count by the exampleId.
 */
func (ec *ExampleController) getExample(c *gin.Context) {

	c.Status(http.StatusOK)
}

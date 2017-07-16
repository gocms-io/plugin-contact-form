package contact_form_ctrl

import (
	"github.com/gin-gonic/gin"
	"github.com/gocms-io/plugin-contact-form/models"
	"github.com/gocms-io/plugin-contact-form/services"
	"github.com/gocms-io/plugin-contact-form/utility/errors"
	"net/http"
)

type ContactFormController struct {
	ServiceGroup *services.ServicesGroup
	Routes       *gin.RouterGroup
}

func DefaultContactFormController(r *gin.Engine, serviceGroup *services.ServicesGroup) *ContactFormController {
	ec := &ContactFormController{
		Routes:       r.Group("/api"),
		ServiceGroup: serviceGroup,
	}

	ec.Default()
	return ec
}

func (ec *ContactFormController) Default() {
	ec.Routes.POST("/contact-form", ec.submitContactForm)
}

/**
* @api {post} /contact-form Submit Contact Form
* @apiName SubmitContactForm
* @apiGroup ContactForm
* @apiUse ContactFormInput
* @apiDescription Send the contents of the contact form to the website admin
 */
func (cfc *ContactFormController) submitContactForm(c *gin.Context) {

	// get contact form data into model
	var cfi models.ContactFormInput
	err := c.BindJSON(&cfi)
	if err != nil {
		errors.Response(c, http.StatusBadRequest, err.Error(), err)
		return
	}

	// create domain model
	//contactForm := models.ContactForm{
	//	Email:      cfi.Email,
	//	Name:       cfi.Name,
	//	Message:    cfi.Message,
	//	Additional: cfi.Additional,
	//	Created:    time.Now(),
	//}
	//
	//err = cfc.ServiceGroup.ContactFormService.ValidateAndSubmit(&contactForm)
	//if err != nil {
	//	errors.Response(c, http.StatusBadRequest, errors.ApiError_Server, err)
	//	return
	//}

	//c.Status(http.StatusOK)
	c.JSON(http.StatusUnauthorized, gin.H{"message": "there was an error"})
}

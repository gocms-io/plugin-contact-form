package content_ctrl

import (
	"github.com/gin-gonic/gin"
)

type StaticContentController struct {
	r *gin.Engine
}

func DefaultStaticContentController(r *gin.Engine) *StaticContentController {
	dc := &StaticContentController{
		r: r,
	}

	dc.Default()
	return dc
}

func (dc *StaticContentController) Default() {
	dc.r.Static("/content", "./content/")
}

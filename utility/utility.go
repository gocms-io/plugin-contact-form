package utility

import (
	"github.com/gin-gonic/gin"
)

var MySQL_TIME_STAMP_FORMAT = "2006-01-02 15:04:05"

//var EVENT_SEARCH_TIME_STAMP_FORMAT = "01-02-2006-15-04-05"
var EVENT_SEARCH_TIME_STAMP_FORMAT = "01-02-2006"

func GetUserIdFromHeader(c *gin.Context) string {
	return c.Request.Header.Get("GOCMS-AUTH-USER-ID")
}

func GetUserNameFromHeader(c *gin.Context) string {
	return c.Request.Header.Get("GOCMS-AUTH-NAME")
}

func GetUserEmailFromHeader(c *gin.Context) string {
	return c.Request.Header.Get("GOCMS-AUTH-EMAIL")
}

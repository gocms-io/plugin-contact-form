package models

import "time"

type ContactForm struct {
	Email      string
	Name       string
	Message    string
	Additional map[string]interface{}
	Created    time.Time
}

/**
* @apiDefine ContactFormInput
* @apiParam (Request) {string} email *required
* @apiParam (Request) {string} name *required
* @apiParam (Request) {string} message
* @apiParam (Request) {string} other Json object of key value pairs containing additional information
 */
type ContactFormInput struct {
	Email      string                 `json:"email" binding:"required"`
	Name       string                 `json:"name" binding:"required"`
	Message    string                 `json:"message"`
	Additional map[string]interface{} `json:"additional,omitempty"`
}

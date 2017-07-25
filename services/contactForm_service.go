package services

import (
	"fmt"
	"github.com/gocms-io/plugin-contact-form/context"
	"github.com/gocms-io/plugin-contact-form/models"
)

type IContactFormService interface {
	ValidateAndSubmit(contactForm *models.ContactForm) error
}

type ContactFormService struct {
	mailService *MailService
}

func DefaultContactFormService(mailService *MailService) *ContactFormService {

	contactFormService := &ContactFormService{
		mailService: mailService,
	}

	return contactFormService

}

func (ms *ContactFormService) ValidateAndSubmit(contactForm *models.ContactForm) error {

	//// recipient /////
	// text body
	htmlBody, textBody := ms.generateEmailToContactFormRecipient(contactForm)

	// create email for recipient
	mail := Mail{
		To:       context.Config.SMTPRecipient,
		Subject:  "Contact Form Submission",
		Body:     textBody,
		BodyHTML: htmlBody,
		From:     fmt.Sprintf("%v <%v>", contactForm.Name, context.Config.SMTPFromAddress),
		ReplyTo:  fmt.Sprintf("%v <%v>", contactForm.Name, contactForm.Email),
	}

	// send email for recipient
	err := ms.mailService.Send(&mail)
	if err != nil {
		fmt.Printf("Contact Form Service - Error sending recipient mail: %v\n", err.Error())
		return err
	}

	//// sender confirmation /////
	// text body
	htmlBodyS, textBodyS := ms.generateEmailToContactFormConfirmationToSender(contactForm)

	// create email for recipient
	mailS := Mail{
		To:       contactForm.Email,
		Subject:  "Contact Form - Confirmation",
		Body:     textBodyS,
		BodyHTML: htmlBodyS,
	}

	// send email for recipient
	err = ms.mailService.Send(&mailS)
	if err != nil {
		fmt.Printf("Contact Form Service - Error sending confirmation email mail: %v\n", err.Error())
		return err
	}

	return nil
}

func (ms *ContactFormService) generateEmailToContactFormRecipient(contactForm *models.ContactForm) (html string, text string) {
	// create body
	textMessage := fmt.Sprintf("From: %v\nEmail: %v\nMessage:\n%v\n", contactForm.Name, contactForm.Email, contactForm.Message)
	htmlMessage := fmt.Sprintf("<h1>Contact Form Submission</h1><h3><b>From:</b> %v</h3><h3><b>Email:</b> %v</h3><h3><b>Message:</b></h3><p>%v</p>", contactForm.Name, contactForm.Email, contactForm.Message)

	// add additional information
	if len(contactForm.Additional) > 0 {
		textMessage += "\n\nAdditional Info:\n"
		htmlMessage += "<br/><h3><b>Additional Info:</b></h3>"
	}
	for key, value := range contactForm.Additional {
		textMessage += fmt.Sprintf("%v: %v\n", key, value)
		htmlMessage += fmt.Sprintf("<p><b>%v:</b> %v</p>", key, value)
	}

	return htmlMessage, textMessage
}

func (ms *ContactFormService) generateEmailToContactFormConfirmationToSender(contactForm *models.ContactForm) (html string, text string) {
	// create body
	textMessage := fmt.Sprintf("Contact Form Confirmation\n\n%v, Thank you for your email, we will review it shortly.", contactForm.Name)
	htmlMessage := fmt.Sprintf("<h1>Contact Form Confirmation</h1><p><b>%v</b>,<br/> Thank you for your email, we will review it shorlty.</p>", contactForm.Name)

	return htmlMessage, textMessage
}

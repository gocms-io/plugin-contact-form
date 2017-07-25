package services

import (
	"fmt"
	"github.com/gocms-io/plugin-contact-form/context"
	"gopkg.in/gomail.v2"
	"io"
	"log"
	"path/filepath"
	"text/template"
	"time"
)

type IMailService interface {
	Send(*Mail) error
}

type MailService struct {
	Dialer          *gomail.Dialer
	From            string
	DefaultTemplate *template.Template
}

type Mail struct {
	To       string
	Subject  string
	Body     string
	BodyHTML string
	From     string
	ReplyTo string
}

func DefaultMailService() *MailService {
	defaultTemplatePath := filepath.Join("../../themes", context.Config.ActiveTheme, "theme_email.tmpl")
	defaultTemplate := template.Must(template.ParseGlob(defaultTemplatePath))

	mailService := &MailService{
		Dialer:          gomail.NewDialer(context.Config.SMTPServer, int(context.Config.SMTPPort), context.Config.SMTPUser, context.Config.SMTPPassword),
		From:     fmt.Sprintf("%v <%v>", context.Config.SMTPFromName, context.Config.SMTPFromAddress),
		DefaultTemplate: defaultTemplate,
	}

	return mailService

}

func (ms *MailService) Send(mail *Mail) error {

	if mail.BodyHTML == "" {
		mail.BodyHTML = mail.Body
	}

	htmlData := map[string]string{
		"subject":     mail.Subject,
		"message":     mail.BodyHTML,
		"year":        time.Now().Format("2006"),
		"headerImage": fmt.Sprintf("http://static.gocms.io/default_assets/default_email_img.jpg"),
	}

	// set from
	from := ms.From
	if mail.From != "" {
		from = mail.From
	}

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", mail.To)
	m.SetHeader("Subject", mail.Subject)
	m.SetBody("text/plain", mail.Body)
	m.AddAlternativeWriter("text/html", func(w io.Writer) error {
		err := ms.DefaultTemplate.Execute(w, htmlData)
		if err != nil {
			fmt.Printf("Error adding alt writter to html email: %v\n", err.Error())
		}
		return err
	})

	// set reply to if needed
	if mail.ReplyTo != "" {
		m.SetHeader("Reply-to", mail.ReplyTo)
	}

	// Send the email
	if !context.Config.SMTPSimulate {
		err := ms.Dialer.DialAndSend(m)
		if err != nil {
			log.Print("Error sending mail: " + err.Error())
			return err
		}
	} else {
		log.Print("Email simulated: " + mail.Body)
	}

	return nil
}

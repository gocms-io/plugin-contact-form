package context

import (
	"github.com/gocms-io/plugin-contact-form/models"
	"log"
)

var Config *RuntimeConfig

type RuntimeConfig struct {
	// DB (GET FROM ENV)
	DbName     string
	DbUser     string
	DbPassword string
	DbServer   string

	// Debug
	Debug bool

	// App
	SettingsRefreshRate int64

	// Theme
	ActiveTheme string

	// SMTP
	SMTPRecipient   string
	SMTPServer      string
	SMTPPort        int64
	SMTPUser        string
	SMTPPassword    string
	SMTPFromAddress string
	SMTPFromName    string
	SMTPSimulate    bool
}

func (c *RuntimeConfig) ApplySettingsToConfig(settings map[string]models.Setting) {

	log.Println("Refreshed Contact Form Plugin Settings")

	// Debug
	c.Debug = GetBoolOrFail("DEBUG", settings)

	// App
	c.SettingsRefreshRate = GetIntOrFail("SETTINGS_REFRESH_RATE", settings)

	// Theme
	c.ActiveTheme = GetStringOrFail("ACTIVE_THEME", settings)

	// SMTP
	c.SMTPRecipient = GetStringOrFail("CONTACT_FORM_RECIPIENT", settings)
	c.SMTPServer = GetStringOrFail("CONTACT_FORM_SMTP_SERVER", settings)
	c.SMTPPort = GetIntOrFail("CONTACT_FORM_SMTP_PORT", settings)
	c.SMTPUser = GetStringOrFail("CONTACT_FORM_SMTP_USER", settings)
	c.SMTPPassword = GetStringOrFail("CONTACT_FORM_SMTP_PASSWORD", settings)
	c.SMTPFromAddress = GetStringOrFail("CONTACT_FORM_SMTP_FROM_ADDRESS", settings)
	c.SMTPFromName = GetStringOrFail("CONTACT_FORM_SMTP_FROM_NAME", settings)

}

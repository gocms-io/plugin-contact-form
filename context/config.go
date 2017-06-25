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
	Debug         bool
	DebugSecurity bool

	// App
	SettingsRefreshRate int64

	// SMTP
	SMTPServer      string
	SMTPPort        int64
	SMTPUser        string
	SMTPPassword    string
	SMTPFromAddress string
	SMTPSimulate    bool
}

func (c *RuntimeConfig) ApplySettingsToConfig(settings map[string]models.Setting) {

	log.Println("Refreshed Contact Form Plugin Settings")

	// Debug
	c.Debug = GetBoolOrFail("DEBUG", settings)

	// App
	c.SettingsRefreshRate = GetIntOrFail("SETTINGS_REFRESH_RATE", settings)

	// SMTP
	c.SMTPServer = GetStringOrFail("SMTP_SERVER", settings)
	c.SMTPPort = GetIntOrFail("SMTP_PORT", settings)
	c.SMTPUser = GetStringOrFail("SMTP_USER", settings)
	c.SMTPPassword = GetStringOrFail("SMTP_PASSWORD", settings)
	c.SMTPFromAddress = GetStringOrFail("SMTP_FROM_ADDRESS", settings)
	c.SMTPSimulate = GetBoolOrFail("SMTP_SIMULATE", settings)

}

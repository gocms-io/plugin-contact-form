package services

import (
	"github.com/gocms-io/plugin-contact-form/context"
	"github.com/gocms-io/plugin-contact-form/repositories"
	"time"
)

type ServicesGroup struct {
	SettingsService ISettingsService
	MailService     IMailService
	ContactFormService IContactFormService
}

func DefaultServicesGroup(rg *repositories.RepositoriesGroup) *ServicesGroup {

	// setup settings
	settingsService := DefaultSettingsService(rg)
	settingsService.RegisterRefreshCallback(context.Config.ApplySettingsToConfig)

	// refresh settings every x minutes
	refreshSettings := time.Duration(context.Config.SettingsRefreshRate) * time.Minute
	context.Schedule.AddTicker(refreshSettings, func() {
		settingsService.RefreshSettingsCache()
	})

	mailService := DefaultMailService()

	sg := &ServicesGroup{
		SettingsService: settingsService,
		ContactFormService: DefaultContactFormService(mailService),
		MailService:     mailService,
	}
	return sg
}

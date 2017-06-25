package services

import (
	"time"
	"github.com/gocms-io/plugin-contact-form/repositories"
	"github.com/gocms-io/plugin-contact-form/context"
)

type ServicesGroup struct {
	SettingsService ISettingsService
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

	sg := &ServicesGroup{
		SettingsService: settingsService,
	}
	return sg
}

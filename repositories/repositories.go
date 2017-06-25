package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/gocms-io/plugin-contact-form/database"
)

type RepositoriesGroup struct {
	SettingsRepository  ISettingsRepository
	dbx                 *sqlx.DB
}

func DefaultRepositoriesGroup(database *database.Database) *RepositoriesGroup {

	// setup repositories
	rg := RepositoriesGroup{
		dbx:                database.SQL.Dbx,
		SettingsRepository: DefaultSettingsRepository(database.SQL.Dbx),
	}
	return &rg
}

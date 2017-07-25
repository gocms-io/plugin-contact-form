package migrations

import "github.com/rubenv/sql-migrate"

func CreateInitial() *migrate.Migration {
	create_initial := migrate.Migration{
		Id: "1",
		Up: []string{`
			INSERT INTO gocms_settings (name, value, description) VALUES('CONTACT_FORM_RECIPIENT', 'RECIPIENT EMAIL FOR CONTACT FORM HERE', 'Email address to receive contact form submissions. ex. info@gocms.io');
			`, `
			INSERT INTO gocms_settings (name, value, description) VALUES('CONTACT_FORM_REPLY_ADDRESS', 'REPLY ADDRESS', 'Replay address for confirmation ex. noreply@gocms.io');
			`, `
			INSERT INTO gocms_settings (name, value, description) VALUES('CONTACT_FORM_SMTP_SERVER', 'SMTP SERVER HERE', 'SMTP server domain name or ip.');
			`, `
			INSERT INTO gocms_settings (name, value, description) VALUES('CONTACT_FORM_SMTP_PORT', '465', 'Port to send smtp mail to.');
			`, `
			INSERT INTO gocms_settings (name, value, description) VALUES('CONTACT_FORM_SMTP_USER', 'USER_HERE', 'Username for SMTP authentication.');
			`, `
			INSERT INTO gocms_settings (name, value, description) VALUES('CONTACT_FORM_SMTP_PASSWORD', 'PASSWORD_HERE', 'Password from SMTP authentication');
			`, `
			INSERT INTO gocms_settings (name, value, description) VALUES('CONTACT_FORM_SMTP_FROM_ADDRESS', 'email@address.com', 'FROM Address and email address for outgoing email. ');
			`, `
			INSERT INTO gocms_settings (name, value, description) VALUES('CONTACT_FORM_SMTP_FROM_NAME', 'FROM NAME HERE', 'FROM Name and email address for outgoing email. ');
			`,
		},
		Down: []string{
			"DELETE FROM gocms_settings WHERE name=CONTACT_FORM_RECIPIENT",
			"DELETE FROM gocms_settings WHERE name=CONTACT_FORM_REPLY_ADDRESS",
			"DELETE FROM gocms_settings WHERE name=CONTACT_FORM_SMTP_SERVER",
			"DELETE FROM gocms_settings WHERE name=CONTACT_FORM_SMTP_PORT",
			"DELETE FROM gocms_settings WHERE name=CONTACT_FORM_SMTP_USER",
			"DELETE FROM gocms_settings WHERE name=CONTACT_FORM_SMTP_PASSWORD",
			"DELETE FROM gocms_settings WHERE name=CONTACT_FORM_SMTP_FROM_ADDRESS",
			"DELETE FROM gocms_settings WHERE name=CONTACT_FORM_SMTP_FROM_NAME",
		},
	}

	return &create_initial
}

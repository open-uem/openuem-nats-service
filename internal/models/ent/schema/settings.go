package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Settings holds the schema definition for the Settings entity.
type Settings struct {
	ent.Schema
}

// Fields of the Settings.
func (Settings) Fields() []ent.Field {
	return []ent.Field{
		field.String("language").Optional(),
		field.String("organization").Optional(),
		field.String("postal_address").Optional(),
		field.String("postal_code").Optional(),
		field.String("locality").Optional(),
		field.String("province").Optional(),
		field.String("state").Optional(),
		field.String("country").Optional().Default("ES"),
		field.String("smtp_server").Optional(),
		field.Int("smtp_port").Optional().Default(587),
		field.String("smtp_user").Optional(),
		field.String("smtp_password").Optional(),
		field.String("smtp_auth").Optional().Default("LOGIN"),
		field.Bool("smtp_tls").Optional().Default(false),
		field.Bool("smtp_starttls").Optional().Default(true),
		field.String("nats_server").Optional(),
		field.String("nats_port").Optional(),
		field.String("message_from").Optional(),
		field.String("max_upload_size").Optional().Default("512M"),
		field.Int("user_cert_years_valid").Optional().Default(1),
		field.Int("nats_request_timeout_seconds").Optional().Default(20),
		field.Int("refresh_time_in_minutes").Optional().Default(5),
		field.Int("session_lifetime_in_minutes").Optional().Default(1440),
		field.String("update_channel").Optional().Default("stable"),
		field.Time("created").Optional().Default(time.Now),
		field.Time("modified").Optional().Default(time.Now).UpdateDefault(time.Now),
		field.Int("agent_report_frequence_in_minutes").Optional().Default(60),
	}
}

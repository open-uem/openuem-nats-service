package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Component holds the schema definition for the Server entity.
type Component struct {
	ent.Schema
}

// Fields of the Component.
func (Component) Fields() []ent.Field {
	return []ent.Field{
		field.String("hostname"),
		field.String("arch"),
		field.String("os"),
		field.Enum("component").Values("ocsp", "nats", "cert-manager", "agent-worker", "notification-worker", "cert-manager-worker", "console"),
		field.String("version"),
		field.Enum("channel").Values("stable", "testing", "devel"),
		field.Enum("update_status").Values("Success", "Error", "Pending").Optional(),
		field.String("update_message").Optional(),
		field.Time("update_when").Optional(),
	}
}

// Indexes of the Component.
func (Component) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("hostname", "arch", "os", "component", "version", "channel").Unique(),
	}
}

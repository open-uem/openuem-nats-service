package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Certificate holds the schema definition for the Certificate entity.
type Certificate struct {
	ent.Schema
}

// Fields of the Certificate.
func (Certificate) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique().StorageKey("serial"),
		field.Enum("type").Values("console", "worker", "agent", "user", "ocsp", "nats", "proxy", "sftp", "updater"),
		field.String("description").Optional(),
		field.Time("expiry").Optional(),
		field.String("uid").Optional(),
	}
}

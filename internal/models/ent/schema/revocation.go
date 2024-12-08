package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Revocation holds the schema definition for the Revocation entity.
type Revocation struct {
	ent.Schema
}

// Fields of the Revocation.
func (Revocation) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique().StorageKey("serial"),
		field.Int("reason").Default(0).Optional(),
		field.String("info").Optional(),
		field.Time("expiry").Optional(),
		field.Time("revoked").Default(time.Now),
	}
}

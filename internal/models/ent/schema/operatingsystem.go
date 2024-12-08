package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// OperatingSystem holds the schema definition for the OperatingSystem entity.
type OperatingSystem struct {
	ent.Schema
}

// Fields of the OperatingSystem.
func (OperatingSystem) Fields() []ent.Field {
	return []ent.Field{
		field.String("type").Optional(),
		field.String("version"),
		field.String("description"),
		field.String("edition").Optional(),
		field.Time("install_date").Optional(),
		field.String("arch").Optional(),
		field.String("username"),
		field.Time("last_bootup_time").Optional(),
	}
}

// Edges of the OperatingSystem.
func (OperatingSystem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", Agent.Type).Unique().Ref("operatingsystem").Required(),
	}
}

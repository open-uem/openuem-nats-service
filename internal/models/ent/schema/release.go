package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Release holds the schema definition for the Release entity.
type Release struct {
	ent.Schema
}

func (Release) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("release_type").Values("agent", "updater", "messenger").Optional(),
		field.String("version").Optional(),
		field.String("channel").Optional(),
		field.String("summary").Optional(),
		field.String("release_notes").Optional(),
		field.String("file_url").Optional(),
		field.String("checksum").Optional(),
		field.Bool("is_critical").Optional(),
		field.Time("release_date").Optional(),
		field.String("os").Optional(),
		field.String("arch").Optional(),
	}
}

// Edges of the Release.
func (Release) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("agents", Agent.Type),
	}
}

// Indexes of the Release.
func (Release) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("release_type", "version", "channel", "os", "arch").Unique(),
	}
}

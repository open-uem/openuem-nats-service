package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// SystemUpdate holds the schema definition for the SystemUpdate entity.
type SystemUpdate struct {
	ent.Schema
}

// Fields of the SystemUpdate.
func (SystemUpdate) Fields() []ent.Field {
	return []ent.Field{
		field.String("status"),
		field.Time("last_install"),
		field.Time("last_search"),
		field.Bool("pending_updates"),
	}
}

// Edges of the SystemUpdate.
func (SystemUpdate) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", Agent.Type).Unique().Ref("systemupdate").Required(),
	}
}

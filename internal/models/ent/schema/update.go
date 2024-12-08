package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Update holds the schema definition for the Update entity.
type Update struct {
	ent.Schema
}

// Fields of the Update.
func (Update) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
		field.Time("date"),
		field.String("support_url").Optional(),
	}
}

// Edges of the Update.
func (Update) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", Agent.Type).Unique().Ref("updates").Required(),
	}
}

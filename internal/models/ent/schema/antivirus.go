package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Antivirus holds the schema definition for the Antivirus entity.
type Antivirus struct {
	ent.Schema
}

// Fields of the Antivirus.
func (Antivirus) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Bool("is_active"),
		field.Bool("is_updated"),
	}
}

// Edges of the Antivirus.
func (Antivirus) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", Agent.Type).Unique().Ref("antivirus").Required(),
	}
}

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Computer holds the schema definition for the Computer entity.
type Computer struct {
	ent.Schema
}

// Fields of the Computer.
func (Computer) Fields() []ent.Field {
	return []ent.Field{
		field.String("manufacturer").Optional(),
		field.String("model").Optional(),
		field.String("serial").Optional(),
		field.Uint64("memory").Optional(),
		field.String("processor").Optional(),
		field.Int64("processor_cores").Optional(),
		field.String("processor_arch").Optional(),
	}
}

// Edges of the Computer.
func (Computer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", Agent.Type).Unique().Ref("computer").Required(),
	}
}

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Monitor holds the schema definition for the Monitor entity.
type Monitor struct {
	ent.Schema
}

// Fields of the Monitor.
func (Monitor) Fields() []ent.Field {
	return []ent.Field{
		field.String("manufacturer").Optional(),
		field.String("model").Optional(),
		field.String("serial").Optional(),
	}
}

// Edges of the Monitor.
func (Monitor) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", Agent.Type).Unique().Ref("monitors").Required(),
	}
}

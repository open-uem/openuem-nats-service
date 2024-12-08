package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Printer holds the schema definition for the Printer entity.
type Printer struct {
	ent.Schema
}

// Fields of the Printer.
func (Printer) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("port").Optional(),
		field.Bool("is_default").Optional(),
		field.Bool("is_network").Optional(),
	}
}

// Edges of the Printer.
func (Printer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", Agent.Type).Unique().Ref("printers").Required(),
	}
}

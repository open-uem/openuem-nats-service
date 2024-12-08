package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Metadata holds the schema definition for the Metadata entity.
type Metadata struct {
	ent.Schema
}

// Fields of the Metadata.
func (Metadata) Fields() []ent.Field {
	return []ent.Field{
		field.String("value"),
	}
}

// Edges of the Metadata.
func (Metadata) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", Agent.Type).Ref("metadata").Unique().Required(),
		edge.From("org", OrgMetadata.Type).Ref("metadata").Unique().Required(),
	}
}

// Indexes of the Metadata.
func (Metadata) Indexes() []ent.Index {
	return []ent.Index{
		index.Edges("org", "owner").Unique(),
	}
}

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// OrgMetadata holds the schema definition for the OrgMetadata entity.
type OrgMetadata struct {
	ent.Schema
}

// Fields of the OrgMetadata.
func (OrgMetadata) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Unique(),
		field.String("description").Optional(),
	}
}

// Edges of the OrgMetadata.
func (OrgMetadata) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("metadata", Metadata.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
	}
}

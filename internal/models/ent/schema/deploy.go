package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Deployment holds the schema definition for the Deployment entity.
type Deployment struct {
	ent.Schema
}

// Fields of the Deployment.
func (Deployment) Fields() []ent.Field {
	return []ent.Field{
		field.String("package_id"),
		field.String("name"),
		field.String("version").Optional(),
		field.Time("installed").Optional().Default(time.Now),
		field.Time("updated").Optional().Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Deployment.
func (Deployment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", Agent.Type).Unique().Ref("deployments").Required(),
	}
}

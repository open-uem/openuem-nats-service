package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Sessions holds the schema definition for the Sessions entity.
type Sessions struct {
	ent.Schema
}

// Fields of the Sessions.
func (Sessions) Fields() []ent.Field {
	return []ent.Field{
		field.Text("id").NotEmpty().Unique().StorageKey("token"),
		field.Bytes("data").NotEmpty(),
		field.Time("expiry"),
	}
}

func (Sessions) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Unique().Ref("sessions"),
	}
}

func (Sessions) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("expiry").StorageKey("sessions_expiry_idx"),
	}
}

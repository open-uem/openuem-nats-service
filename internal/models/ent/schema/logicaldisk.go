package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Logical Disk holds the schema definition for the Logical Disk entity.
type LogicalDisk struct {
	ent.Schema
}

// Fields of the Logical Disk.
func (LogicalDisk) Fields() []ent.Field {
	return []ent.Field{
		field.String("label"),
		field.String("filesystem").Optional(),
		field.Int8("usage").Default(0),
		field.String("size_in_units").Optional(),
		field.String("remaining_space_in_units").Optional(),
		field.String("volume_name").Optional(),
		field.String("bitlocker_status").Optional(),
	}
}

// Edges of the Logical Disk.
func (LogicalDisk) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", Agent.Type).Unique().Ref("logicaldisks").Required(),
	}
}

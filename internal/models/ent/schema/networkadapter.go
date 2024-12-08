package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// NetworkAdapter holds the schema definition for the NetworkAdapter entity.
type NetworkAdapter struct {
	ent.Schema
}

// Fields of the NetworkAdapter.
func (NetworkAdapter) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("mac_address"),
		field.String("addresses"),
		field.String("subnet").Optional(),
		field.String("default_gateway").Optional(),
		field.String("dns_servers").Optional(),
		field.String("dns_domain").Optional(),
		field.Bool("dhcp_enabled").Optional(),
		field.Time("dhcp_lease_obtained").Optional(),
		field.Time("dhcp_lease_expired").Optional(),
		field.String("speed"),
	}
}

// Edges of the NetworkAdapter.
func (NetworkAdapter) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", Agent.Type).Unique().Ref("networkadapters").Required(),
	}
}

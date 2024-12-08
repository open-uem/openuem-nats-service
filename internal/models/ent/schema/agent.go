package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Agent holds the schema definition for the Agent entity.
type Agent struct {
	ent.Schema
}

// Fields of the Agent.
func (Agent) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").NotEmpty().Unique().StorageKey("oid"),
		field.String("os").NotEmpty(),
		field.String("hostname").NotEmpty(),
		field.String("ip").Default(""),
		field.String("mac").Default(""),
		field.Time("first_contact").Optional(),
		field.Time("last_contact").Optional(),
		field.String("vnc").Optional().Default(""),
		field.Text("notes").Optional(),
		field.String("update_task_status").Optional().Default(""),
		field.String("update_task_description").Optional().Default(""),
		field.String("update_task_result").Optional().Default(""),
		field.Time("update_task_execution").Optional(),
		field.String("update_task_version").Optional().Default(""),
		field.String("vnc_proxy_port").Optional().Default(""),
		field.String("sftp_port").Optional().Default(""),
		field.Enum("status").Values("WaitingForAdmission", "Enabled", "Disabled").Optional().Default("WaitingForAdmission"),
		field.Bool("certificate_ready").Optional().Default(false),
		field.Bool("restart_required").Optional().Default(false),
	}
}

// Edges of the Agent.
func (Agent) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("computer", Computer.Type).Unique().Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("operatingsystem", OperatingSystem.Type).Unique().Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("systemupdate", SystemUpdate.Type).Unique().Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("antivirus", Antivirus.Type).Unique().Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("logicaldisks", LogicalDisk.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("apps", App.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("monitors", Monitor.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("shares", Share.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("printers", Printer.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("networkadapters", NetworkAdapter.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("deployments", Deployment.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("updates", Update.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("tags", Tag.Type),
		edge.To("metadata", Metadata.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.From("release", Release.Type).Ref("agents").Unique(),
	}
}

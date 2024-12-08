// Code generated by ent, DO NOT EDIT.

package openuem_ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/doncicuto/openuem_ent/agent"
	"github.com/doncicuto/openuem_ent/update"
)

// Update is the model entity for the Update schema.
type Update struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Date holds the value of the "date" field.
	Date time.Time `json:"date,omitempty"`
	// SupportURL holds the value of the "support_url" field.
	SupportURL string `json:"support_url,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UpdateQuery when eager-loading is set.
	Edges         UpdateEdges `json:"edges"`
	agent_updates *string
	selectValues  sql.SelectValues
}

// UpdateEdges holds the relations/edges for other nodes in the graph.
type UpdateEdges struct {
	// Owner holds the value of the owner edge.
	Owner *Agent `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UpdateEdges) OwnerOrErr() (*Agent, error) {
	if e.Owner != nil {
		return e.Owner, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: agent.Label}
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Update) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case update.FieldID:
			values[i] = new(sql.NullInt64)
		case update.FieldTitle, update.FieldSupportURL:
			values[i] = new(sql.NullString)
		case update.FieldDate:
			values[i] = new(sql.NullTime)
		case update.ForeignKeys[0]: // agent_updates
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Update fields.
func (u *Update) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case update.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int(value.Int64)
		case update.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				u.Title = value.String
			}
		case update.FieldDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field date", values[i])
			} else if value.Valid {
				u.Date = value.Time
			}
		case update.FieldSupportURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field support_url", values[i])
			} else if value.Valid {
				u.SupportURL = value.String
			}
		case update.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field agent_updates", values[i])
			} else if value.Valid {
				u.agent_updates = new(string)
				*u.agent_updates = value.String
			}
		default:
			u.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Update.
// This includes values selected through modifiers, order, etc.
func (u *Update) Value(name string) (ent.Value, error) {
	return u.selectValues.Get(name)
}

// QueryOwner queries the "owner" edge of the Update entity.
func (u *Update) QueryOwner() *AgentQuery {
	return NewUpdateClient(u.config).QueryOwner(u)
}

// Update returns a builder for updating this Update.
// Note that you need to call Update.Unwrap() before calling this method if this Update
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *Update) Update() *UpdateUpdateOne {
	return NewUpdateClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the Update entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *Update) Unwrap() *Update {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("openuem_ent: Update is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *Update) String() string {
	var builder strings.Builder
	builder.WriteString("Update(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("title=")
	builder.WriteString(u.Title)
	builder.WriteString(", ")
	builder.WriteString("date=")
	builder.WriteString(u.Date.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("support_url=")
	builder.WriteString(u.SupportURL)
	builder.WriteByte(')')
	return builder.String()
}

// Updates is a parsable slice of Update.
type Updates []*Update

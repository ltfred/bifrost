// Code generated by entc, DO NOT EDIT.

package entmodels

import (
	"bifrost/gen/entmodels/district"
	"fmt"
	"strings"

	"github.com/facebook/ent/dialect/sql"
)

// District is the model entity for the District schema.
type District struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// Pid holds the value of the "pid" field.
	Pid uint64 `json:"pid,omitempty"`
	// DistrictName holds the value of the "district_name" field.
	DistrictName string `json:"district_name,omitempty"`
	// ShorterName holds the value of the "shorter_name" field.
	ShorterName string `json:"shorter_name,omitempty"`
	// CarCode holds the value of the "car_code" field.
	CarCode string `json:"car_code,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*District) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullInt64{},  // pid
		&sql.NullString{}, // district_name
		&sql.NullString{}, // shorter_name
		&sql.NullString{}, // car_code
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the District fields.
func (d *District) assignValues(values ...interface{}) error {
	if m, n := len(values), len(district.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	d.ID = uint64(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field pid", values[0])
	} else if value.Valid {
		d.Pid = uint64(value.Int64)
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field district_name", values[1])
	} else if value.Valid {
		d.DistrictName = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field shorter_name", values[2])
	} else if value.Valid {
		d.ShorterName = value.String
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field car_code", values[3])
	} else if value.Valid {
		d.CarCode = value.String
	}
	return nil
}

// Update returns a builder for updating this District.
// Note that, you need to call District.Unwrap() before calling this method, if this District
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *District) Update() *DistrictUpdateOne {
	return (&DistrictClient{config: d.config}).UpdateOne(d)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (d *District) Unwrap() *District {
	tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("entmodels: District is not a transactional entity")
	}
	d.config.driver = tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *District) String() string {
	var builder strings.Builder
	builder.WriteString("District(")
	builder.WriteString(fmt.Sprintf("id=%v", d.ID))
	builder.WriteString(", pid=")
	builder.WriteString(fmt.Sprintf("%v", d.Pid))
	builder.WriteString(", district_name=")
	builder.WriteString(d.DistrictName)
	builder.WriteString(", shorter_name=")
	builder.WriteString(d.ShorterName)
	builder.WriteString(", car_code=")
	builder.WriteString(d.CarCode)
	builder.WriteByte(')')
	return builder.String()
}

// Districts is a parsable slice of District.
type Districts []*District

func (d Districts) config(cfg config) {
	for _i := range d {
		d[_i].config = cfg
	}
}

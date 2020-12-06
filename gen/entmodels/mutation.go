// Code generated by entc, DO NOT EDIT.

package entmodels

import (
	"bifrost/gen/entmodels/district"
	"bifrost/gen/entmodels/predicate"
	"context"
	"fmt"
	"sync"

	"github.com/facebook/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeDistrict = "District"
)

// DistrictMutation represents an operation that mutate the Districts
// nodes in the graph.
type DistrictMutation struct {
	config
	op            Op
	typ           string
	id            *uint64
	pid           *uint64
	addpid        *uint64
	district_name *string
	shorter_name  *string
	car_code      *string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*District, error)
	predicates    []predicate.District
}

var _ ent.Mutation = (*DistrictMutation)(nil)

// districtOption allows to manage the mutation configuration using functional options.
type districtOption func(*DistrictMutation)

// newDistrictMutation creates new mutation for $n.Name.
func newDistrictMutation(c config, op Op, opts ...districtOption) *DistrictMutation {
	m := &DistrictMutation{
		config:        c,
		op:            op,
		typ:           TypeDistrict,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withDistrictID sets the id field of the mutation.
func withDistrictID(id uint64) districtOption {
	return func(m *DistrictMutation) {
		var (
			err   error
			once  sync.Once
			value *District
		)
		m.oldValue = func(ctx context.Context) (*District, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().District.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withDistrict sets the old District of the mutation.
func withDistrict(node *District) districtOption {
	return func(m *DistrictMutation) {
		m.oldValue = func(context.Context) (*District, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m DistrictMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m DistrictMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("entmodels: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that, this
// operation is accepted only on District creation.
func (m *DistrictMutation) SetID(id uint64) {
	m.id = &id
}

// ID returns the id value in the mutation. Note that, the id
// is available only if it was provided to the builder.
func (m *DistrictMutation) ID() (id uint64, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetPid sets the pid field.
func (m *DistrictMutation) SetPid(u uint64) {
	m.pid = &u
	m.addpid = nil
}

// Pid returns the pid value in the mutation.
func (m *DistrictMutation) Pid() (r uint64, exists bool) {
	v := m.pid
	if v == nil {
		return
	}
	return *v, true
}

// OldPid returns the old pid value of the District.
// If the District object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *DistrictMutation) OldPid(ctx context.Context) (v uint64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldPid is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldPid requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPid: %w", err)
	}
	return oldValue.Pid, nil
}

// AddPid adds u to pid.
func (m *DistrictMutation) AddPid(u uint64) {
	if m.addpid != nil {
		*m.addpid += u
	} else {
		m.addpid = &u
	}
}

// AddedPid returns the value that was added to the pid field in this mutation.
func (m *DistrictMutation) AddedPid() (r uint64, exists bool) {
	v := m.addpid
	if v == nil {
		return
	}
	return *v, true
}

// ResetPid reset all changes of the "pid" field.
func (m *DistrictMutation) ResetPid() {
	m.pid = nil
	m.addpid = nil
}

// SetDistrictName sets the district_name field.
func (m *DistrictMutation) SetDistrictName(s string) {
	m.district_name = &s
}

// DistrictName returns the district_name value in the mutation.
func (m *DistrictMutation) DistrictName() (r string, exists bool) {
	v := m.district_name
	if v == nil {
		return
	}
	return *v, true
}

// OldDistrictName returns the old district_name value of the District.
// If the District object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *DistrictMutation) OldDistrictName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldDistrictName is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldDistrictName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDistrictName: %w", err)
	}
	return oldValue.DistrictName, nil
}

// ResetDistrictName reset all changes of the "district_name" field.
func (m *DistrictMutation) ResetDistrictName() {
	m.district_name = nil
}

// SetShorterName sets the shorter_name field.
func (m *DistrictMutation) SetShorterName(s string) {
	m.shorter_name = &s
}

// ShorterName returns the shorter_name value in the mutation.
func (m *DistrictMutation) ShorterName() (r string, exists bool) {
	v := m.shorter_name
	if v == nil {
		return
	}
	return *v, true
}

// OldShorterName returns the old shorter_name value of the District.
// If the District object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *DistrictMutation) OldShorterName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldShorterName is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldShorterName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldShorterName: %w", err)
	}
	return oldValue.ShorterName, nil
}

// ResetShorterName reset all changes of the "shorter_name" field.
func (m *DistrictMutation) ResetShorterName() {
	m.shorter_name = nil
}

// SetCarCode sets the car_code field.
func (m *DistrictMutation) SetCarCode(s string) {
	m.car_code = &s
}

// CarCode returns the car_code value in the mutation.
func (m *DistrictMutation) CarCode() (r string, exists bool) {
	v := m.car_code
	if v == nil {
		return
	}
	return *v, true
}

// OldCarCode returns the old car_code value of the District.
// If the District object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *DistrictMutation) OldCarCode(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCarCode is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCarCode requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCarCode: %w", err)
	}
	return oldValue.CarCode, nil
}

// ResetCarCode reset all changes of the "car_code" field.
func (m *DistrictMutation) ResetCarCode() {
	m.car_code = nil
}

// Op returns the operation name.
func (m *DistrictMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (District).
func (m *DistrictMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during
// this mutation. Note that, in order to get all numeric
// fields that were in/decremented, call AddedFields().
func (m *DistrictMutation) Fields() []string {
	fields := make([]string, 0, 4)
	if m.pid != nil {
		fields = append(fields, district.FieldPid)
	}
	if m.district_name != nil {
		fields = append(fields, district.FieldDistrictName)
	}
	if m.shorter_name != nil {
		fields = append(fields, district.FieldShorterName)
	}
	if m.car_code != nil {
		fields = append(fields, district.FieldCarCode)
	}
	return fields
}

// Field returns the value of a field with the given name.
// The second boolean value indicates that this field was
// not set, or was not define in the schema.
func (m *DistrictMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case district.FieldPid:
		return m.Pid()
	case district.FieldDistrictName:
		return m.DistrictName()
	case district.FieldShorterName:
		return m.ShorterName()
	case district.FieldCarCode:
		return m.CarCode()
	}
	return nil, false
}

// OldField returns the old value of the field from the database.
// An error is returned if the mutation operation is not UpdateOne,
// or the query to the database was failed.
func (m *DistrictMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case district.FieldPid:
		return m.OldPid(ctx)
	case district.FieldDistrictName:
		return m.OldDistrictName(ctx)
	case district.FieldShorterName:
		return m.OldShorterName(ctx)
	case district.FieldCarCode:
		return m.OldCarCode(ctx)
	}
	return nil, fmt.Errorf("unknown District field %s", name)
}

// SetField sets the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *DistrictMutation) SetField(name string, value ent.Value) error {
	switch name {
	case district.FieldPid:
		v, ok := value.(uint64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPid(v)
		return nil
	case district.FieldDistrictName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDistrictName(v)
		return nil
	case district.FieldShorterName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetShorterName(v)
		return nil
	case district.FieldCarCode:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCarCode(v)
		return nil
	}
	return fmt.Errorf("unknown District field %s", name)
}

// AddedFields returns all numeric fields that were incremented
// or decremented during this mutation.
func (m *DistrictMutation) AddedFields() []string {
	var fields []string
	if m.addpid != nil {
		fields = append(fields, district.FieldPid)
	}
	return fields
}

// AddedField returns the numeric value that was in/decremented
// from a field with the given name. The second value indicates
// that this field was not set, or was not define in the schema.
func (m *DistrictMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case district.FieldPid:
		return m.AddedPid()
	}
	return nil, false
}

// AddField adds the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *DistrictMutation) AddField(name string, value ent.Value) error {
	switch name {
	case district.FieldPid:
		v, ok := value.(uint64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddPid(v)
		return nil
	}
	return fmt.Errorf("unknown District numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared
// during this mutation.
func (m *DistrictMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicates if this field was
// cleared in this mutation.
func (m *DistrictMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value for the given name. It returns an
// error if the field is not defined in the schema.
func (m *DistrictMutation) ClearField(name string) error {
	return fmt.Errorf("unknown District nullable field %s", name)
}

// ResetField resets all changes in the mutation regarding the
// given field name. It returns an error if the field is not
// defined in the schema.
func (m *DistrictMutation) ResetField(name string) error {
	switch name {
	case district.FieldPid:
		m.ResetPid()
		return nil
	case district.FieldDistrictName:
		m.ResetDistrictName()
		return nil
	case district.FieldShorterName:
		m.ResetShorterName()
		return nil
	case district.FieldCarCode:
		m.ResetCarCode()
		return nil
	}
	return fmt.Errorf("unknown District field %s", name)
}

// AddedEdges returns all edge names that were set/added in this
// mutation.
func (m *DistrictMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all ids (to other nodes) that were added for
// the given edge name.
func (m *DistrictMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this
// mutation.
func (m *DistrictMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all ids (to other nodes) that were removed for
// the given edge name.
func (m *DistrictMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this
// mutation.
func (m *DistrictMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean indicates if this edge was
// cleared in this mutation.
func (m *DistrictMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value for the given name. It returns an
// error if the edge name is not defined in the schema.
func (m *DistrictMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown District unique edge %s", name)
}

// ResetEdge resets all changes in the mutation regarding the
// given edge name. It returns an error if the edge is not
// defined in the schema.
func (m *DistrictMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown District edge %s", name)
}

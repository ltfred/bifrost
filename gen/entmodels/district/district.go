// Code generated by entc, DO NOT EDIT.

package district

const (
	// Label holds the string label denoting the district type in the database.
	Label = "district"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPid holds the string denoting the pid field in the database.
	FieldPid = "pid"
	// FieldDistrictName holds the string denoting the district_name field in the database.
	FieldDistrictName = "district_name"
	// FieldShorterName holds the string denoting the shorter_name field in the database.
	FieldShorterName = "shorter_name"
	// FieldCarCode holds the string denoting the car_code field in the database.
	FieldCarCode = "car_code"

	// Table holds the table name of the district in the database.
	Table = "districts"
)

// Columns holds all SQL columns for district fields.
var Columns = []string{
	FieldID,
	FieldPid,
	FieldDistrictName,
	FieldShorterName,
	FieldCarCode,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DistrictNameValidator is a validator for the "district_name" field. It is called by the builders before save.
	DistrictNameValidator func(string) error
	// ShorterNameValidator is a validator for the "shorter_name" field. It is called by the builders before save.
	ShorterNameValidator func(string) error
	// CarCodeValidator is a validator for the "car_code" field. It is called by the builders before save.
	CarCodeValidator func(string) error
)
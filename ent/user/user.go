// Code generated by entc, DO NOT EDIT.

package user

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAge holds the string denoting the age field in the database.
	FieldAge = "age"
	// FieldNames holds the string denoting the names field in the database.
	FieldNames = "names"
	// FieldLastnames holds the string denoting the lastnames field in the database.
	FieldLastnames = "lastnames"
	// Table holds the table name of the user in the database.
	Table = "users"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldAge,
	FieldNames,
	FieldLastnames,
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
	// AgeValidator is a validator for the "age" field. It is called by the builders before save.
	AgeValidator func(int) error
	// NamesValidator is a validator for the "names" field. It is called by the builders before save.
	NamesValidator func(string) error
	// LastnamesValidator is a validator for the "lastnames" field. It is called by the builders before save.
	LastnamesValidator func(string) error
)

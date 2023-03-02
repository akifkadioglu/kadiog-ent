// Code generated by ent, DO NOT EDIT.

package notes

const (
	// Label holds the string label denoting the notes type in the database.
	Label = "notes"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// Table holds the table name of the notes in the database.
	Table = "notes"
)

// Columns holds all SQL columns for notes fields.
var Columns = []string{
	FieldID,
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
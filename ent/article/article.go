// Code generated by ent, DO NOT EDIT.

package article

import "github.com/kangana1024/go-graphql-ent/ent/schema"

const (
	// Label holds the string label denoting the article type in the database.
	Label = "article"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the article in the database.
	Table = "articles"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "articles"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
)

// Columns holds all SQL columns for article fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldDescription,
	FieldContent,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "articles"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_id",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultTitle holds the default value on creation for the "title" field.
	DefaultTitle string
	// DefaultDescription holds the default value on creation for the "description" field.
	DefaultDescription string
	// DefaultContent holds the default value on creation for the "content" field.
	DefaultContent *schema.ArticleContent
)

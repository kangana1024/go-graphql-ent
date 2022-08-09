// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldFirstname holds the string denoting the firstname field in the database.
	FieldFirstname = "firstname"
	// FieldLastname holds the string denoting the lastname field in the database.
	FieldLastname = "lastname"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldBirthDate holds the string denoting the birth_date field in the database.
	FieldBirthDate = "birth_date"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// Table holds the table name of the user in the database.
	Table = "users"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldFirstname,
	FieldLastname,
	FieldUsername,
	FieldBirthDate,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
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
	// DefaultFirstname holds the default value on creation for the "firstname" field.
	DefaultFirstname string
	// FirstnameValidator is a validator for the "firstname" field. It is called by the builders before save.
	FirstnameValidator func(string) error
	// DefaultLastname holds the default value on creation for the "lastname" field.
	DefaultLastname string
	// LastnameValidator is a validator for the "lastname" field. It is called by the builders before save.
	LastnameValidator func(string) error
	// DefaultUsername holds the default value on creation for the "username" field.
	DefaultUsername string
	// UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	UsernameValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
)
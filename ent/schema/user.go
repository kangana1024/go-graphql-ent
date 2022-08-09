package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("firstname").Default("").MaxLen(100),
		field.String("lastname").Default("").MaxLen(100),
		field.String("username").Default("").MaxLen(50),
		field.Time("birth_date").Default(nil),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now),
		field.Time("deleted_at").Default(nil),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

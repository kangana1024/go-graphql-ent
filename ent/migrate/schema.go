// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ArticlesColumns holds the columns for the "articles" table.
	ArticlesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "title", Type: field.TypeString, Default: ""},
		{Name: "description", Type: field.TypeString, Default: ""},
		{Name: "content", Type: field.TypeJSON},
		{Name: "user_id", Type: field.TypeInt, Nullable: true},
	}
	// ArticlesTable holds the schema information for the "articles" table.
	ArticlesTable = &schema.Table{
		Name:       "articles",
		Columns:    ArticlesColumns,
		PrimaryKey: []*schema.Column{ArticlesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "articles_users_articles",
				Columns:    []*schema.Column{ArticlesColumns[4]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "firstname", Type: field.TypeString, Size: 100, Default: ""},
		{Name: "password", Type: field.TypeString, Size: 100, Default: ""},
		{Name: "lastname", Type: field.TypeString, Size: 100, Default: ""},
		{Name: "username", Type: field.TypeString, Size: 50, Default: ""},
		{Name: "birth_date", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ArticlesTable,
		UsersTable,
	}
)

func init() {
	ArticlesTable.ForeignKeys[0].RefTable = UsersTable
}

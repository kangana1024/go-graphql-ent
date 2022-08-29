package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Article holds the schema definition for the Article entity.
type Article struct {
	ent.Schema
}

type ArticleContent struct {
	Time   int64 `json:"time"`
	Blocks []struct {
		Type string      `json:"type"`
		Data interface{} `json:"data"`
	} `json:"blocks"`
	Version string `json:"version"`
}

// Fields of the Article.
func (Article) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").Default(""),
		field.String("description").Default(""),
		field.JSON("content", &ArticleContent{}).Default(&ArticleContent{}),
	}
}

// Edges of the Article.
func (Article) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("articles").
			Unique(),
	}
}

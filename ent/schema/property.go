package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Property holds the schema definition for the Property entity.
type Property struct {
	ent.Schema
}

// Fields of the Property.
func (Property) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("type").Values(
			"house",
			"apartment",
			"premises",
			"office",
			"vehicle",
		).Default("apartment"),
		field.String("name").NotEmpty(),
		field.Bool("deleted"),
	}
}

// Edges of the Property.
func (Property) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Ref("properties").Unique(),
	}
}

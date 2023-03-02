package schema

import "entgo.io/ent"

// Notes holds the schema definition for the Notes entity.
type Notes struct {
	ent.Schema
}

// Fields of the Notes.
func (Notes) Fields() []ent.Field {
	return nil
}

// Edges of the Notes.
func (Notes) Edges() []ent.Edge {
	return nil
}

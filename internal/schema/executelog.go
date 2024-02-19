package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type ExecuteLog struct {
	ent.Schema
}

func (ExecuteLog) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "execute_logs"},
		entsql.WithComments(true),
	}
}

func (ExecuteLog) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),

		field.String("category").
			Optional().
			Nillable().
			Comment("수집할 데이터의 유형"),

		field.String("keyword").
			Optional().
			Nillable().
			Comment("검색한 키워드"),

		field.Int("requestCount").
			Optional().Nillable().
			Comment("수집 희망 개수"),

		field.Int("collectedCount").
			Optional().Nillable().
			Comment("수집한 개수"),
	}
}

func (ExecuteLog) Mixin() []ent.Mixin {
	return []ent.Mixin{
		DefaultTimeMixin{},
	}
}

func (ExecuteLog) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("products", Product.Type),
	}
}

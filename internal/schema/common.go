package schema

import (
	"boilerplate/pkg/types"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// ------------------- Default  -------------------

type DefaultTimeMixin struct {
	mixin.Schema
}

func (DefaultTimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("createdAt").
			Immutable().
			SchemaType(
				map[string]string{
					dialect.Postgres: "timestamp",
				},
			).
			GoType(types.TimeString{}).
			Default(types.TimeString{}.Now),

		field.Time("updatedAt").
			GoType(types.TimeString{}).
			SchemaType(
				map[string]string{
					dialect.Postgres: "timestamp",
				},
			).
			Default(types.TimeString{}.Now).
			UpdateDefault(types.TimeString{}.Now),
	}
}

// ------------------- Delete -------------------

type WithDeletedTimeMixin struct {
	mixin.Schema
}

func (WithDeletedTimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("createdAt").
			Immutable().
			SchemaType(
				map[string]string{
					dialect.MySQL: "timestamp",
				},
			).
			GoType(types.TimeString{}).
			Default(types.TimeString{}.Now).StructTag(`sql:"created_at"`).
			Comment("생성일"),

		field.Time("updatedAt").
			GoType(types.TimeString{}).
			SchemaType(
				map[string]string{
					dialect.MySQL: "timestamp",
				},
			).
			Default(types.TimeString{}.Now).
			UpdateDefault(types.TimeString{}.Now).StructTag(`sql:"updated_at"`).
			Comment("수징일"),

		field.Time("deletedAt").
			GoType(types.TimeString{}).
			SchemaType(
				map[string]string{
					dialect.MySQL: "timestamp",
				},
			).
			Nillable().
			Optional().StructTag(`sql:"deleted_at"`).
			Comment("삭제일"),
	}
}

type Status int

const (
	New Status = iota
	Normal
	Hidden
	PriceChange
)

func (s *Status) ToString() string {
	if s == nil {
		return ""
	}

	if *s == New {
		return "신규"
	} else if *s == Normal {
		return "정상"
	} else if *s == Hidden {
		return "품절"
	} else if *s == PriceChange {
		return "가격변동"
	}
	return ""

}

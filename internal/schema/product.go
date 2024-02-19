package schema

import (
	"boilerplate/pkg/types"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type Product struct {
	ent.Schema
}

func (Product) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "products"},
		entsql.WithComments(true),
	}
}

func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),

		field.String("externalID").
			Comment("상품을 가져온 곳의 대표 ID"),

		field.Int("status").
			GoType(Status(0)).
			Optional().Nillable().
			Comment("상태"),

		field.String("url").
			Optional().Nillable().
			Comment("구매 URL"),

		field.String("descriptionText").
			Optional().Nillable().
			Comment("상품 설명 텍스트만"),

		field.String("descriptionHTML").
			Optional().Nillable().
			Comment("상품 설명 HTML"),

		field.String("name").
			Comment("상품 이름"),

		field.String("nameEng").
			Optional().Nillable().
			Comment("상품 영어 이름"),

		field.String("brandName").
			Optional().Nillable().
			Comment("브랜드 이름"),

		field.String("origin").
			Optional().Nillable().
			Comment("원산지"),

		field.String("modelName").
			Optional().Nillable().
			Comment("모델명"),

		field.Int("quantity").
			Optional().Nillable().
			Comment("상품의 현재 수량"),

		field.Int("immediatePrice").
			Comment("현재 구매할 수 있는 상품 가격 필수가격"),

		field.Int("releasePrice").
			Optional().Nillable().
			Comment("출시 가격"),

		field.Int("discountPrice").
			Optional().Nillable().
			Comment("할인 가격"),

		field.Int("recentPrice").
			Optional().Nillable().
			Comment("최근 거래 가격"),

		field.Int("deliveryPrice").
			Optional().Nillable().
			Comment("배송비"),

		field.Float("avgRating").
			Optional().Nillable().
			Comment("평균 별점"),

		field.Int("purchaseCount").
			Optional().Nillable().
			Comment("구매 개수"),

		field.Int("reviewCount").
			Optional().Nillable().
			Comment("리뷰 개수"),

		field.Time("openDate").
			GoType(types.TimeString{}).
			Optional().
			Nillable().
			Comment("등록일"),

		field.Int("wishCount").
			Optional().Nillable().
			Comment("찜 개수"),
	}
}

func (Product) Mixin() []ent.Mixin {
	return []ent.Mixin{
		DefaultTimeMixin{},
	}
}

func (Product) Indexes() []ent.Index {
	return []ent.Index{

		index.Fields("externalID"),
	}
}

func (Product) Edges() []ent.Edge {
	return []ent.Edge{

		// many to many
		edge.From("executeLog", ExecuteLog.Type).
			Ref("products"),
	}
}

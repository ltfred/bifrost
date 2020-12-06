package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// District holds the schema definition for the District entity.
type District struct {
	ent.Schema
}

const (
	nameLen = 50
)



// Fields of the District.
func (District) Fields() []ent.Field {
	return []ent.Field {
		field.Uint64("id").Unique(),field.Uint64("pid").Comment("上级id"),
		field.String("district_name").MaxLen(nameLen).Comment("地区名称"),
		field.String("shorter_name").MaxLen(nameLen).Comment("简称"),
		field.String("car_code").MaxLen(nameLen).Comment("车牌代码"),
	}
}

// Edges of the District.
func (District) Edges() []ent.Edge {
	return nil
}

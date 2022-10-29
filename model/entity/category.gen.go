// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

import (
	"time"

	"github.com/go-sonic/sonic/consts"
)

const TableNameCategory = "category"

// Category mapped from table <category>
type Category struct {
	ID          int32               `gorm:"column:id;type:int;primaryKey;autoIncrement:true" json:"id"`
	CreateTime  time.Time           `gorm:"column:create_time;type:datetime;not null" json:"create_time"`
	UpdateTime  *time.Time          `gorm:"column:update_time;type:datetime" json:"update_time"`
	Description string              `gorm:"column:description;type:varchar(100);not null" json:"description"`
	Name        string              `gorm:"column:name;type:varchar(255);not null;index:category_name,priority:1" json:"name"`
	ParentID    int32               `gorm:"column:parent_id;type:int;not null;index:category_parent_id,priority:1" json:"parent_id"`
	Password    string              `gorm:"column:password;type:varchar(255);not null" json:"password"`
	Slug        string              `gorm:"column:slug;type:varchar(255);not null;uniqueIndex:uniq_category_slug,priority:1" json:"slug"`
	Thumbnail   string              `gorm:"column:thumbnail;type:varchar(1023);not null" json:"thumbnail"`
	Priority    int32               `gorm:"column:priority;type:int;not null" json:"priority"`
	Type        consts.CategoryType `gorm:"column:type;type:int;not null" json:"type"`
}

// TableName Category's table name
func (*Category) TableName() string {
	return TableNameCategory
}
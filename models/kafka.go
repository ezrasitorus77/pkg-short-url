package models

import "time"

type (
	Topic struct {
		CreatedAt *time.Time `gorm:"column:created_at"`
		Name      string     `gorm:"column:name"`
		ID        int8       `gorm:"column:id;primary_key"`
	}
)

type (
	Group struct {
		CreatedAt *time.Time `gorm:"column:created_at"`
		Name      string     `gorm:"column:name"`
		TopicID   int8       `gorm:"column:topic_id"`
		ID        int8       `gorm:"column:id;primary_key"`
	}
)

func (tbl *Topic) TableName() string {
	return "topics"
}

func (tbl *Group) TableName() string {
	return "groups"
}

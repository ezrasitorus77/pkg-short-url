package models

import "time"

type (
	LogDBKafka struct {
		CreatedAt *time.Time `gorm:"column:created_at"`
		Key       string     `gorm:"column:key"`
		Message   string     `gorm:"column:message"`
		Error     string     `gorm:"column:error"`
		Entity    string     `gorm:"column:entity"`
		ID        int16      `gorm:"column:id;primary_key"`
		URLID     int16      `gorm:"column:url_id"`
		TopicID   int8       `gorm:"column:topic_id"`
		GroupID   int8       `gorm:"column:group_id"`
		IsSuccess bool       `gorm:"column:is_success"`
	}
)

func (tbl *LogDBKafka) TableName() string {
	return "log_trail_kafka"
}

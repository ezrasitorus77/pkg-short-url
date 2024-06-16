package models

import (
	"time"
)

type (
	LogDBRequest struct {
		CreatedAt  time.Time `gorm:"column:created_at"`
		URL        string    `gorm:"column:url"`
		Query      string    `gorm:"column:query"`
		Method     string    `gorm:"column:method"`
		Headers    string    `gorm:"column:headers"`
		ReqBody    string    `gorm:"column:request_body"`
		RespCode   string    `gorm:"column:response_code"`
		RespBody   string    `gorm:"column:response_body"`
		Error      string    `gorm:"column:error"`
		ID         int       `gorm:"column:id;primary_key"`
		HTTPStatus int8      `gorm:"column:http_status"`
		UserID     int8      `gorm:"column:user_id"`
	}
)

func (tbl *LogDBRequest) TableName() string {
	return "log_trail_requests"
}

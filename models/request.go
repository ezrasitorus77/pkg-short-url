package models

import "time"

type (
	LogDBRequest struct {
		ID         int       `gorm:"column:id;primary_key"`
		UserID     int       `gorm:"column:user_id"`
		URL        string    `gorm:"column:url"`
		Query      string    `gorm:"column:query"`
		Method     string    `gorm:"column:method"`
		Headers    string    `gorm:"column:headers"`
		ReqBody    string    `gorm:"column:request_body"`
		HTTPStatus int       `gorm:"column:http_status"`
		RespCode   string    `gorm:"column:response_code"`
		RespBody   string    `gorm:"column:response_body"`
		Error      string    `gorm:"column:error"`
		CreatedAt  time.Time `gorm:"column:created_at"`
	}
)

func (tbl *LogDBRequest) TableName() string {
	return "log_trail_requests"
}

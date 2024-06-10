package models

import "time"

type (
	URL struct {
		CreatedAt     *time.Time `gorm:"column:created_at" json:"created_at"`
		UpdatedAt     *time.Time `gorm:"column:updated_at" json:"updated_at"`
		ValidSince    *time.Time `gorm:"column:valid_since" json:"valid_since"`
		ValidUntil    *time.Time `gorm:"column:valid_until" json:"valid_until"`
		Short         string     `gorm:"column:short_url" json:"short_url"`
		Long          string     `gorm:"column:long_url" json:"long_url"`
		CurrentVisits int32      `gorm:"column:curr_visits" json:"curr_visits"`
		MaxVisits     int32      `gorm:"column:max_visits" json:"max_visits"`
		UserID        int16      `gorm:"column:user_id" json:"user_id"`
		ID            int16      `gorm:"column:id;primary_key" json:"id"`
		Status        bool       `gorm:"column:status" json:"status"`
	}
)

func (tbl *URL) TableName() string {
	return "urls"
}

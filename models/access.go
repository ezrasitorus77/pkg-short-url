package models

import "time"

type (
	RouteAccess struct {
		CreatedAt *time.Time `gorm:"column:created_at"`
		APIKey    string     `gorm:"column:api_key"`
		Alias     string     `gorm:"column:alias"`
		Group     string     `gorm:"column:group"`
		ID        int16      `gorm:"column:id;primary_key" json:"id"`
	}

	MethodAccess struct {
		CreatedAt *time.Time `gorm:"column:created_at"`
		APIKey    string     `gorm:"column:api_key"`
		Method    string     `gorm:"column:method"`
		ID        int16      `gorm:"column:id;primary_key" json:"id"`
	}
)

func (tbl *RouteAccess) TableName() string {
	return "routes_access"
}

func (tbl *MethodAccess) TableName() string {
	return "methods_access"
}

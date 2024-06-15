package query

import (
	"gorm.io/gorm"
)

func Q_WHERE_EXACT_ROUTE_GROUP(group string) QChain {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("group = ?", group)
	}
}

func Q_WHERE_EXACT_METHOD_NAME(method string) QChain {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("method = ?", method)
	}
}

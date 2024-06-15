package query

import (
	"gorm.io/gorm"
)

var (
	Q_WHERE_EXACT_ROUTE_GROUP QChain = func(db *gorm.DB) *gorm.DB {
		return db.Where("group = ?", true)
	}

	Q_WHERE_EXACT_METHOD_NAME QChain = func(db *gorm.DB) *gorm.DB {
		return db.Where("method = ?", true)
	}
)

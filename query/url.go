package query

import (
	"gorm.io/gorm"
)

var (
	Q_WHERE_ACTIVE_URL QChain = func(db *gorm.DB) *gorm.DB {
		return db.Where("status = ?", true)
	}
)

func Q_WHERE_INSENSITIVE_CERTAIN_URL(short_url string) QChain {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("LOWER(short_url) = ?", short_url)
	}
}

func Q_WHERE_SENSITIVE_CERTAIN_URL(short_url string) QChain {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("short_url = ?", short_url)
	}
}

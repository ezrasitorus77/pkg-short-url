package logdb

import (
	"github.com/ezrasitorus77/pkg-short-url/models"
	"gorm.io/gorm"
)

type (
	LogKafkaBuilder struct {
		models.LogDBKafka
	}
)

func (lk LogKafkaBuilder) record(db *gorm.DB, log models.LogDBKafka) {
	db.Create(&log)
}

func (lk LogKafkaBuilder) Success(db *gorm.DB, log models.LogDBKafka) {
	log.IsSuccess = true

	go lk.record(db, log)
}

func (lk LogKafkaBuilder) Error(db *gorm.DB, log models.LogDBKafka) {
	log.IsSuccess = false

	go lk.record(db, log)
}

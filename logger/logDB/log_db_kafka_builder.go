package logdb

import (
	"github.com/ezrasitorus77/pkg-short-url/models"
	"gorm.io/gorm"
)

type (
	LogKafkaBuilder struct {
		models.LogDBKafka
		db *gorm.DB
	}
)

func NewLogKafka(db *gorm.DB) LogKafkaBuilder {
	return LogKafkaBuilder{
		db: db,
	}
}

func (lk LogKafkaBuilder) record(log models.LogDBKafka) {
	lk.db.Create(&log)
}

func (lk LogKafkaBuilder) Success(log models.LogDBKafka) error {
	log.IsSuccess = true

	go lk.record(log)

	return nil
}

func (lk LogKafkaBuilder) Error(log models.LogDBKafka) error {
	log.IsSuccess = false

	go lk.record(log)

	return nil
}

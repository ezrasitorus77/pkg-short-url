package query

import (
	"gorm.io/gorm"
)

type (
	QChain     func(db *gorm.DB) *gorm.DB
	QValidator struct{}
	QBuilder   struct {
		Query
		Validator QValidator
	}
)

type Query interface {
	First(db *gorm.DB, chain ...QChain) (data interface{}, err error)
	Last(db *gorm.DB, chain ...QChain) (data interface{}, err error)
	All(db *gorm.DB, chain ...QChain) (data interface{}, err error)
	Write(db *gorm.DB, data interface{}) (err error)
}

func (qv QValidator) validate(db *gorm.DB, chain ...QChain) *gorm.DB {
	if len(chain) > 0 {
		for i := range chain {
			db = chain[i](db)
		}
	}

	return db
}

func (b QBuilder) First(db *gorm.DB, chain ...QChain) (data interface{}, err error) {
	db = b.Validator.validate(db, chain...)

	return b.Query.First(db, chain...)
}

func (b QBuilder) Last(db *gorm.DB, chain ...QChain) (data interface{}, err error) {
	db = b.Validator.validate(db, chain...)

	return b.Query.Last(db, chain...)
}

func (b QBuilder) All(db *gorm.DB, chain ...QChain) (data interface{}, err error) {
	db = b.Validator.validate(db, chain...)

	return b.Query.All(db, chain...)
}

func (b QBuilder) Write(db *gorm.DB, data interface{}) (err error) {
	return b.Query.Write(db, data)
}

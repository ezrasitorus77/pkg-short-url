package models

import (
	"time"

	"github.com/valyala/fasthttp"
)

type (
	Visit struct {
		CreatedAt  time.Time `gorm:"column:created_at" json:"created_at"`
		VisitedURL string    `gorm:"column:visited_url" json:"visited_url"`
		IP         string    `gorm:"column:ip" json:"ip"`
		Referer    string    `gorm:"column:referer" json:"referer"`
		Agent      string    `gorm:"column:agent" json:"agent"`
		URLID      int16     `gorm:"column:url_id" json:"url_id"`
		ID         int16     `gorm:"column:id;primary_key" json:"id"`
		IsValid    bool      `gorm:"column:is_valid" json:"is_valid"`
	}
)

func (tbl *Visit) TableName() string {
	return "visits"
}

func (tbl *Visit) Record(r *fasthttp.Request, ip string) {
	tbl.Agent = string(r.Header.UserAgent())
	tbl.Referer = string(r.Header.Referer())
	tbl.IP = ip
	tbl.VisitedURL = string(r.RequestURI())
	tbl.CreatedAt = time.Now()
}

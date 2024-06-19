package models

import (
	"time"

	"github.com/ipinfo/go/v2/ipinfo"
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

	Location struct {
		CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
		TimeZone    string    `gorm:"column:time_zone" json:"time_zone"`
		OrgName     string    `gorm:"column:org" json:"org"`
		LongLat     string    `gorm:"column:longlat" json:"longlat"`
		CountryCode string    `gorm:"column:country_code" json:"country_code"`
		Region      string    `gorm:"column:region" json:"region"`
		City        string    `gorm:"column:city" json:"city"`
		VisitID     int16     `gorm:"column:visit_id" json:"visit_id"`
		ID          int16     `gorm:"column:id;primary_key" json:"id"`
	}
)

func (tbl *Visit) TableName() string {
	return "visits"
}

func (tbl *Location) TableName() string {
	return "visits_location"
}

func (tbl *Visit) Record(r *fasthttp.Request, ip string) {
	tbl.Agent = string(r.Header.UserAgent())
	tbl.Referer = string(r.Header.Referer())
	tbl.IP = ip
	tbl.VisitedURL = string(r.RequestURI())
}

func (tbl *Location) Set(visitID int16, ipData *ipinfo.Core) {
	tbl.VisitID = visitID
	tbl.City = ipData.City
	tbl.Region = ipData.Region
	tbl.CountryCode = ipData.Country
	tbl.LongLat = ipData.Location
	tbl.OrgName = ipData.Org
	tbl.TimeZone = ipData.Timezone
}

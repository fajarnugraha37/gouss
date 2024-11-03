package model

import (
	"github.com/fajarnugraha37/gouss/pkg/model/embed"
)

type VisitModel struct {
	embed.IdentityModel
	UrlId     string
	Ip        string
	UserAgent string
	Geo       string
	Browser   string
	City      string
	Country   string
	OS        string
	Region    string
	Referrer  string
	embed.AuditModel
}

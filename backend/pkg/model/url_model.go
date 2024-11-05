package model

import (
	"time"
	"github.com/fajarnugraha37/gouss/pkg/model/embed"
)

type UrlModel struct {
	embed.IdentityModel
	Description *string
	OriginalUrl string
	ShortUrl    *string
	CustomUrl	*string
	IsCustom    bool
	IsDisable   bool
	DisableAt   *time.Time
	ActiveAt    *time.Time
	ExpiresAt   *time.Time
	embed.UtmModel
	embed.OwnerModel
	embed.AuditModel
}

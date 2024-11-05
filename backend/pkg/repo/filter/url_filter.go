package filter


type UrlFilter struct {
	IdentityFilter

	OriginalUrl *string
	ShortUrl    *string
	CustomUrl   *string
	IsCustom    *bool
	IsDisable   *bool
	IsActive    *bool

	OwnerFilter
	AuditFilter
}

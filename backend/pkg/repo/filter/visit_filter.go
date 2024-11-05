package filter

type VisitFilter struct {
	IdentityFilter

	UrlId       *string
	OriginalUrl *string
	ShortUrl    *string
	CustomUrl   *string
	Ip          *string
	Browser     *string
	City        *string
	Country     *string
	OS          *string
	Region      *string
	Referrer    *string

	AuditFilter
}

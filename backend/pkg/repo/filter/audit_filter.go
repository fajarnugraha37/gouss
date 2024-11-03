package filter

import "time"

type AuditFilter struct {
	CreatedStart *time.Time
	CreatedEnd   *time.Time
	CreatedBy    *string
	UpdatedStart *time.Time
	UpdatedEnd   *time.Time
	UpdateBy     *string
	DeletedFlag  bool
	DeletedStart *time.Time
	DeletedEnd   *time.Time
	DeletedBy    *string
}

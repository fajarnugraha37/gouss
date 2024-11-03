package embed

import "time"

type AuditModel struct {
	CreatedAt   time.Time
	CreatedBy   string
	UpdatedAt   *time.Time
	UpdateBy    *string
	DeletedFlag bool
	DeletedAt   *time.Time
	DeletedBy   *string
}

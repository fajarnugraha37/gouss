package filter

type PageableFilter struct {
	PageSize int32
	Page     int32
	SortableFilter
}

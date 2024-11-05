package filter

type OrderDirection int

const (
	ASC OrderDirection = iota
	DESC
)

type SortableFilter struct {
	SortBy   *string
	Order    OrderDirection
}

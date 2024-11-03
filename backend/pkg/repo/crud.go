package repo

import (
	"github.com/fajarnugraha37/gouss/pkg/model"
	"github.com/fajarnugraha37/gouss/pkg/repo/filter"
	im "github.com/fajarnugraha37/gouss/pkg/repo/in_memory"
)

type ICrud[M any, F any] interface {
	FindOne(f *F) *M
	FindAll(f *F, page *filter.PageableFilter) *[]M
	Create(m *M) *M
	Update(m *M) *M
	Delete(m *M) *M
}
	
var (
	UrlRepo ICrud[model.UrlModel, filter.UrlFilter]
	VisitRepo ICrud[model.VisitModel, filter.VisitFilter]
)

func init() {
	UrlRepo = &im.UrlRepo{
		Store: map[string]model.UrlModel{},
	}
	VisitRepo = &im.VisitRepo{
		Store: map[string]model.VisitModel{},
	}
}
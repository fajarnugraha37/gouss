package in_memory

import (
	"github.com/fajarnugraha37/gouss/pkg/model"
	"github.com/fajarnugraha37/gouss/pkg/repo/filter"
)

type VisitRepo struct {
	Store map[string]model.VisitModel
}

func (r *VisitRepo) FindOne(f *filter.VisitFilter) *model.VisitModel {
	// TODO: implement
	return &model.VisitModel{}
}

func (r *VisitRepo) FindAll(f *filter.VisitFilter, page *filter.PageableFilter) *[]model.VisitModel {
	// TODO: implement
	var result []model.VisitModel
	return &result 
}

func (r *VisitRepo) Create(m *model.VisitModel) *model.VisitModel {
	// TODO: implement
	return &model.VisitModel{}
}

func (r *VisitRepo) Update(m *model.VisitModel) *model.VisitModel {
	// TODO: implement
	return &model.VisitModel{}
}

func (r *VisitRepo) Delete(m *model.VisitModel) *model.VisitModel {
	return &model.VisitModel{}
}
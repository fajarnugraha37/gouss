package in_memory

import (
	"github.com/fajarnugraha37/gouss/pkg/model"
	"github.com/fajarnugraha37/gouss/pkg/repo/filter"
)

type UrlRepo struct {
	Store map[string]model.UrlModel
}

func (r *UrlRepo) FindOne(f *filter.UrlFilter) *model.UrlModel {
	// TODO: implement
	return &model.UrlModel{}
}

func (r *UrlRepo) FindAll(f *filter.UrlFilter, page *filter.PageableFilter) *[]model.UrlModel {
	// TODO: implement
	var result []model.UrlModel
	return &result 
}

func (r *UrlRepo) Create(m *model.UrlModel) *model.UrlModel {
	// TODO: implement
	return &model.UrlModel{}
}

func (r *UrlRepo) Update(m *model.UrlModel) *model.UrlModel {
	// TODO: implement
	return &model.UrlModel{}
}

func (r *UrlRepo) Delete(m *model.UrlModel) *model.UrlModel {
	// TODO: implement
	return &model.UrlModel{}
}
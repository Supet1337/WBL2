package repository

import (
	models "dev11/internal"
	"errors"
)

type Repository struct {
	data map[int]models.Event
}
type IRepository interface {
	Create(model models.Event) error
	Update(model models.Event) error
	Delete(model models.Event) error
	Get() map[int]models.Event
}

func NewRepository() IRepository {
	repository := Repository{data: make(map[int]models.Event)}
	return &repository
}

func (r *Repository) Create(model models.Event) error {
	if _, exist := r.data[model.Id]; !exist {
		r.data[model.Id] = model
	} else {
		return errors.New("Event already exist")
	}
	return nil
}

func (r *Repository) Update(model models.Event) error {
	if _, exist := r.data[model.Id]; exist {
		r.data[model.Id] = model
	} else {
		return errors.New("Event doesn`t exist")
	}
	return nil
}

func (r *Repository) Delete(model models.Event) error {
	if _, exist := r.data[model.Id]; exist {
		delete(r.data, model.Id)
	} else {
		return errors.New("Event doesn`t exist")
	}
	return nil
}

func (r *Repository) Get() map[int]models.Event {
	return r.data
}

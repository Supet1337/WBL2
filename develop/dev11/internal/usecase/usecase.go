package usecase

import (
	models "dev11/internal"
	"dev11/internal/repository"
	"encoding/json"
	"errors"
	"log"
	"time"
)

type UseCase struct {
	Repository repository.IRepository
}
type IUseCase interface {
	Create(model []byte) error
	Update(model []byte) error
	Delete(model []byte) error
	Get(date string) ([]byte, error)
}

func NewUseCase() IUseCase {
	useCase := UseCase{}
	useCase.Repository = repository.NewRepository()
	return &useCase
}

func (useCase *UseCase) Create(model []byte) error {
	var event models.Event

	err := json.Unmarshal(model, &event)
	if err != nil {
		return err
	}
	if event.Id > 0 {
		err = useCase.Repository.Create(event)
		if err != nil {
			return err
		}
	} else {
		return errors.New("Invalid id")
	}
	return nil
}

func (useCase *UseCase) Update(model []byte) error {
	var event models.Event

	err := json.Unmarshal(model, &event)
	if err != nil {
		return err
	}
	if event.Id > 0 {
		err = useCase.Repository.Update(event)
		if err != nil {
			return err
		}
	} else {
		return errors.New("Invalid id")
	}
	return nil
}

func (useCase *UseCase) Delete(model []byte) error {
	var event models.Event

	err := json.Unmarshal(model, &event)
	if err != nil {
		return err
	}
	if event.Id > 0 {
		err = useCase.Repository.Delete(event)
		if err != nil {
			return err
		}
	} else {
		return errors.New("Invalid id")
	}
	return nil
}

func (useCase *UseCase) Get(date string) ([]byte, error) {
	repo := useCase.Repository.Get()
	events := make([]models.Event, 0)
	switch date {
	case "day":
		for _, event := range repo {
			if event.Date.After(time.Now()) && event.Date.Before(time.Now().Add(time.Hour*24)) {
				events = append(events, event)
			}
		}
	case "week":
		for _, event := range repo {
			if event.Date.After(time.Now()) && event.Date.Before(time.Now().Add(time.Hour*24*7)) {
				events = append(events, event)
			}
		}
	case "month":
		for _, event := range repo {
			if event.Date.After(time.Now()) && event.Date.Before(time.Now().Add(time.Hour*24*7*30)) {
				events = append(events, event)
			}
		}
	}

	if len(events) > 0 {
		res, err := json.Marshal(events)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		return res, nil
	}
	return nil, errors.New("No events")

}

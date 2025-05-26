package service

import (
	"TestTaskEffectiveMobile/internal/names/model"
	namedata "TestTaskEffectiveMobile/pkg/api/nameData"
	"context"
	"fmt"
)

type Repository interface {
	Create(ctx context.Context, p model.Person) error
	Get(ctx context.Context, userID int) (model.Person, error)
	Delete(ctx context.Context, userID int) error
	Update(ctx context.Context, userID int, p model.Person) error
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s Service) CreateUser(ctx context.Context, p model.PersonRequest) error {
	age, err := namedata.GetAge(p.Name)
	if err != nil {
		return fmt.Errorf("service.CreateUser: %w", err)
	}

	gender, err := namedata.GetGender(p.Name)
	if err != nil {
		return fmt.Errorf("service.CreateUser: %w", err)
	}

	country, err := namedata.GetCountry(p.Name)
	if err != nil {
		return fmt.Errorf("service.CreateUser: %w", err)
	}

	person := model.Person{
		Name:       p.Name,
		Surname:    p.Surname,
		Patronymic: p.Patronymic,
		Age:        age,
		Gender:     gender,
		Country:    country,
	}

	err = s.repo.Create(ctx, person)
	if err != nil {
		return fmt.Errorf("service.CreateUser: %w", err)
	}

	return nil
}

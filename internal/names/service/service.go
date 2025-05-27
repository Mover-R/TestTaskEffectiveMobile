package service

import (
	"TestTaskEffectiveMobile/internal/names/model"
	namedata "TestTaskEffectiveMobile/pkg/api/nameData"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Repository interface {
	Create(ctx context.Context, p model.Person) error
	Get(ctx context.Context, userID int) (model.Person, error)
	Delete(ctx context.Context, userID int) error
	Update(ctx context.Context, userID int, p model.Person) error
	FindList(ctx context.Context, filter model.Filter) ([]model.Person, error)
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

	countrys, err := namedata.GetCountry(p.Name)
	if err != nil {
		return fmt.Errorf("service.CreateUser: %w", err)
	}

	country := model.CountryInf{
		CountryID:   "",
		Probability: -1,
	}
	for _, c := range countrys {
		if country.Probability < c.Probability {
			country.Probability = c.Probability
			country.CountryID = c.CountryID
		}
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

func (s Service) GetUser(ctx context.Context, userID int) (model.Person, error) {
	p, err := s.repo.Get(ctx, userID)
	if err == pgx.ErrNoRows {
		return model.Person{}, err
	}
	if err != nil {
		return model.Person{}, fmt.Errorf("serrvice.GetUser: %w", err)
	}

	return p, nil
}

func (s Service) UpdateUser(ctx context.Context, userID int, p model.Person) error {
	err := s.repo.Update(ctx, userID, p)
	if err != nil {
		return fmt.Errorf("service.UpdateUser: %w", err)
	}

	return nil
}

func (s Service) DeleteUser(ctx context.Context, userID int) error {
	err := s.repo.Delete(ctx, userID)
	if err != nil {
		return fmt.Errorf("service.DeleteUser: %w", err)
	}

	return nil
}

func (s Service) FindWithFilter(ctx context.Context, filter model.Filter) ([]model.Person, error) {
	persons, err := s.repo.FindList(ctx, filter)
	if err == pgx.ErrNoRows {
		return persons, err
	}
	if err != nil {
		return persons, fmt.Errorf("service.FindWithFilter: %w", err)
	}

	return persons, nil
}

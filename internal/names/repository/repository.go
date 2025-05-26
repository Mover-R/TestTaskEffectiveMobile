package repository

import (
	query "TestTaskEffectiveMobile/db"
	"TestTaskEffectiveMobile/internal/names/model"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	pgDB *pgxpool.Pool
}

func NewRepository(pgDB *pgxpool.Pool) *Repository {
	return &Repository{pgDB: pgDB}
}

func (r *Repository) Create(ctx context.Context, p model.Person) error {
	var userID int
	err := r.pgDB.QueryRow(ctx, query.InsertUserQuery,
		p.Name,
		p.Surname,
		p.Patronymic,
		p.Age,
		p.Gender).Scan(&userID)
	if err != nil {
		return fmt.Errorf("repository.Create: %w", err)
	}

	for _, country := range p.Country {
		_, err = r.pgDB.Exec(ctx, query.InsertCountryQuery, userID, country.CountryID, country.Probability)
		if err != nil {
			return fmt.Errorf("repository.Create - insert country: %w", err)
		}
	}

	return nil
}

func (r *Repository) Get(ctx context.Context, userID int) (model.Person, error) {
	var p model.Person
	err := r.pgDB.QueryRow(ctx, query.GetUserQuery).Scan(
		&p.UserID,
		&p.Name,
		&p.Surname,
		p.Patronymic,
		&p.Age,
		&p.Gender)
	if err != nil {
		return model.Person{}, fmt.Errorf("repository.Get: %w", err)
	}

	rows, err := r.pgDB.Query(ctx, query.GetCountryQuery, userID)
	if err != nil {
		return model.Person{}, fmt.Errorf("repository.Get: %w", err)
	}

	for rows.Next() {
		var counry model.CountryInf
		err = rows.Scan(&counry.CountryID, &counry.Probability)
		if err != nil {
			return model.Person{}, fmt.Errorf("repository.Get: %w", err)
		}
		p.Country = append(p.Country, counry)
	}

	return p, nil
}

func (r *Repository) Delete(ctx context.Context, userID int) error {
	_, err := r.pgDB.Exec(ctx, query.DeleteQuery, userID)
	if err != nil {
		return fmt.Errorf("repository.Delete: %w", err)
	}
	return nil
}

func (r *Repository) Update(ctx context.Context, userID int, p model.Person) error {
	return nil
}

package repository

import (
	query "TestTaskEffectiveMobile/db"
	"TestTaskEffectiveMobile/internal/names/model"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
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

	_, err = r.pgDB.Exec(ctx, query.InsertCountryQuery, userID, p.Country.CountryID, p.Country.Probability)
	if err != nil {
		return fmt.Errorf("repository.Create - insert country: %w", err)
	}

	return nil
}

func (r *Repository) Get(ctx context.Context, userID int) (model.Person, error) {
	var p model.Person
	err := r.pgDB.QueryRow(ctx, query.GetUserQuery, userID).Scan(
		&p.UserID,
		&p.Name,
		&p.Surname,
		&p.Patronymic,
		&p.Age,
		&p.Gender,
		&p.Country.CountryID,
		&p.Country.Probability)
	if err == pgx.ErrNoRows {
		return model.Person{}, err
	}
	if err != nil {
		return model.Person{}, fmt.Errorf("repository.Get: %w", err)
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
	_, err := r.pgDB.Exec(ctx, query.UpdateUserQuery,
		p.Name,
		p.Surname,
		p.Patronymic,
		p.Age,
		p.Gender,
		p.UserID)
	if err != nil {
		return fmt.Errorf("repository.Update: %w", err)
	}

	_, err = r.pgDB.Exec(ctx, query.UpdateCountryQuery,
		p.Country.CountryID,
		p.Country.Probability,
		p.UserID)
	if err != nil {
		return fmt.Errorf("repository.Update: %w", err)
	}

	return nil
}

func (r *Repository) FindList(ctx context.Context, filter model.Filter) ([]model.Person, error) {
	var res []model.Person

	query := query.FindWithFilter
	paramCount := 1
	params := []interface{}{}
	if filter.Name != nil {
		query += fmt.Sprintf(" AND u.name LIKE $%d", paramCount)
		params = append(params, filter.Name)
		paramCount++
	}
	if filter.Surname != nil {
		query += fmt.Sprintf(" AND u.surname LIKE $%d", paramCount)
		params = append(params, filter.Surname)
		paramCount++
	}
	if filter.Patronymic != nil {
		query += fmt.Sprintf(" AND u.patronymic LIKE $%d", paramCount)
		params = append(params, filter.Patronymic)
		paramCount++
	}
	if filter.AgeMin != nil {
		query += fmt.Sprintf(" AND u.age >= $%d", paramCount)
		params = append(params, filter.AgeMin)
		paramCount++
	}
	if filter.AgeMax != nil {
		query += fmt.Sprintf(" AND u.age <= $%d", paramCount)
		params = append(params, filter.AgeMax)
		paramCount++
	}
	if filter.Gender != nil {
		query += fmt.Sprintf(" AND u.gender = $%d", paramCount)
		params = append(params, filter.Gender)
		paramCount++
	}
	if filter.CountryID != nil {
		query += fmt.Sprintf(" AND uc.country = $%d", paramCount)
		params = append(params, filter.CountryID)
		paramCount++
	}

	offset := (filter.Page - 1) * filter.PerPage
	query += fmt.Sprintf(" LIMIT %d OFFSET %d", filter.PerPage, offset)

	rows, err := r.pgDB.Query(ctx, query, params...)
	if err == pgx.ErrNoRows {
		return res, err
	}
	if err != nil {
		return res, fmt.Errorf("repository.FindList: %w", err)
	}

	for rows.Next() {
		var p model.Person
		err = rows.Scan(
			&p.UserID,
			&p.Name,
			&p.Surname,
			&p.Patronymic,
			&p.Age,
			&p.Gender,
			&p.Country.CountryID,
			&p.Country.Probability)
		if err != nil {
			return res, fmt.Errorf("repository.FindList: %w", err)
		}
		res = append(res, p)
	}

	return res, nil
}

package query

const (
	InsertUserQuery = `INSERT INTO mig.users 
					(name, surname, patronymic, age, gender)
					VALUES($1, $2, $3, $4, $5)
					RETURNING user_id;`
	DeleteQuery        = `DELETE FROM mig.users WHERE user_id = $1;`
	InsertCountryQuery = `INSERT INTO mig.user_country 
                     (user_id, country, probability)
                     VALUES($1, $2, $3);`
	GetUserQuery = `SELECT 
						u.user_id,
						u.name,
						u.surname, 
						u.patronymic,
						u.age,
						u.gender,
						uc.country,
						uc.probability
					FROM mig.users u
					LEFT JOIN mig.user_country uc
					WHERE u.user_id = $1;`
	UpdateUserQuery = `UPDATE mig.users u SET
						u.name = $1,
						u.surname = $2,
						u.patronymic = $3,
						u.age = $4,
						u.gender = $5
						WHERE u.user_id = $6`
	UpdateCountryQuery = `UPDATE mig.user_country uc SET
						uc.country = $1,
						uc.probability = $2
						WHERE uc.user_id = $3`
	FindWithFilter = `SELECT u.user_id,
						u.name,
						u.surname, 
						u.patronymic,
						u.age,
						u.gender,
						uc.country,
						uc.probability
					FROM mig.users u
					LEFT JOIN mig.user_country uc
					WHERE 1 = 1`
)

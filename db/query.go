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
						u.gender
					FROM mig.users u
					WHERE u.user_id = $1;`
	GetCountryQuery = `SELECT
						c.country,
						c.probability,
					FROM mig.user_country c
					WHERE c.user_id = $1;`
)

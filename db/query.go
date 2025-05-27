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
					LEFT JOIN mig.user_country uc ON uc.user_id = u.user_id
					WHERE u.user_id = $1;`
	UpdateUserQuery = `UPDATE mig.users SET
						name = $1,
						surname = $2,
						patronymic = $3,
						age = $4,
						gender = $5
						WHERE user_id = $6`
	UpdateCountryQuery = `UPDATE mig.user_country SET
						country = $1,
						probability = $2
						WHERE user_id = $3`
	FindWithFilter = `SELECT u.user_id,
						u.name,
						u.surname, 
						u.patronymic,
						u.age,
						u.gender,
						uc.country,
						uc.probability
					FROM mig.users u
					JOIN mig.user_country uc ON u.user_id = uc.user_id
					WHERE 1 = 1`
)

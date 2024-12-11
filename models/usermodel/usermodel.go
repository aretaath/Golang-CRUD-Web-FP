package usermodel

import (
	"go-web/config"
	"go-web/entities"
	"time"
)

func GetAll() []entities.User {
	rows, err := config.DB.Query("SELECT * FROM users")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var users []entities.User

	for rows.Next() {
		var user entities.User
		if err := rows.Scan(&user.Id, &user.Name, &user.CreatedAt, &user.UpdatedAt); err != nil {
			panic(err)
		}

		users = append(users, user)
	}

	return users
}

func Create(user entities.User) bool {

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	result, err := config.DB.Exec(`
		INSERT INTO users (name, created_at, updated_at) 
		VALUE (?, ?, ?)`,
		user.Name,
		user.CreatedAt,
		user.UpdatedAt,
	)

	if err != nil {
		panic(err)
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	return lastInsertId > 0
}

func Edit(id int) entities.User {
	row := config.DB.QueryRow(`SELECT id, name FROM users WHERE id = ? `, id)

	var user entities.User

	if err := row.Scan(&user.Id, &user.Name); err != nil {
		panic(err.Error())
	}

	return user
}

func Update(id int, user entities.User) bool {
	query, err := config.DB.Exec(`UPDATE users SET name = ?, updated_at = ? where id = ?`, user.Name, user.UpdatedAt, id)
	if err != nil {
		panic(err)
	}

	result, err := query.RowsAffected()
	if err != nil {
		panic(err)
	}

	return result > 0
}

func Delete(id int) error {
	_, err := config.DB.Exec("DELETE FROM users WHERE id = ?", id)
	return err
}

package repositories

import (
	"fmt"
	"home_work_sql_gin/iternals/models"

	"github.com/jmoiron/sqlx"
)

func CreateUsersTable(db *sqlx.DB) {
	stmt := `
		CREATE TABLE IF NOT EXISTS tb_user (
			id SERIAL PRIMARY KEY,
			name VARCHAR,
			email VARCHAR
		)
	`

	db.MustExec(stmt)
	fmt.Println("Table was created")
}

func AddNewUser(u models.User, db *sqlx.DB) error {
	stmt := `
		INSERT INTO tb_user(name, email)
		VALUES($1, $2)
	`

	_, err := db.Exec(stmt, u.Name, u.Email)

	return err
}

func GetAllUsers(db *sqlx.DB) ([]models.User, error) {
	stmt := `
		SELECT * FROM tb_user
	`

	var users []models.User

	err := db.Select(&users, stmt)

	return users, err
}

func GetUserById(userId int, db *sqlx.DB) ([]models.User, error) {
	stmt := `
		SELECT * FROM Tb_user
		WHERE id = $1
	`

	var user []models.User

	err := db.Select(&user, stmt, userId)

	return user, err
}

func UpdateUserEmail(userId int, newEmail string, db *sqlx.DB) error {
	stmt := `
		UPDATE tb_user SET email = $1
		WHERE id = $2
	`

	_, err := db.Exec(stmt, newEmail, userId)

	return err
}

func DeleteUserByName(userName string, db *sqlx.DB) error {
	stmt := `
		DELETE FROM tb_user
		WHERE name = $1
	`

	_, err := db.Exec(stmt, userName)

	return err
}

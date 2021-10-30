package userRepository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/eyupfatihersoy/app-tryout-1/models"
)

type UserRepository struct{}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (u UserRepository) SignUp(db *sql.DB, user models.User) models.User {
	stmt := "insert into users (email, password, clientType) values($1, $2, $3) RETURNING id"
	err := db.QueryRow(stmt, user.Email, user.Password, user.ClientType).Scan(&user.ID)

	logFatal(err)

	user.Password = ""
	return user
}

func (u UserRepository) LogIn(db *sql.DB, user models.User) (models.User, error) {
	row := db.QueryRow("select * from users where email=$1", user.Email)
	err := row.Scan(&user.ID, &user.Email, &user.Password, &user.ClientType)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (u UserRepository) ChangeUserType(db *sql.DB, uType models.ChangeUserType) (models.ChangeUserType, error) {
	result, err := db.Exec("update users set clientType=$1 where email=$2", uType.ClientType, uType.Email)

	if err != nil {
		return uType, err
	}

	rowsUpdated, err := result.RowsAffected()
	logFatal(err)

	fmt.Println("rows updated", rowsUpdated)

	return uType, nil
}

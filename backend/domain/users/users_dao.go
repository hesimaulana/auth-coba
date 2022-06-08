package users

import (
	"fmt"

	users_db "github.com/hesimaulana/auth-coba/backend/datasource/sqlite/users_db"
	"github.com/hesimaulana/auth-coba/backend/utils/errors"
)

var (
	insertStmt     = "INSERT INTO users (nama, email, password) VALUES (?,?,?);"
	getByEmailStmt = "SELECT id, nama, email, password FROM users WHERE email=?;"
	getByUserStmt  = "SELECT id, nama email FROM users WHERE id=?;"
)

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.DB.Prepare(insertStmt)
	if err != nil {
		return errors.NewBadRequestError("database error")
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Password)
	if saveErr != nil {
		fmt.Println(saveErr)
		return errors.NewInternalServerError("database error")
	}

	userID, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError("database error")
	}
	user.ID = userID
	return nil
}

func (user *User) GetByEmail() *errors.RestErr {
	stmt, err := users_db.DB.Prepare(getByUserStmt)
	if err != nil {
		return errors.NewInternalServerError("invalid email")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Email)
	if getErr := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password); getErr != nil {
		fmt.Println(getErr)
		return errors.NewInternalServerError("database error")
	}
	return nil
}

func (user *User) GetByID() *errors.RestErr {
	stmt, err := users_db.DB.Prepare(getByEmailStmt)
	if err != nil {
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.ID)
	if getErr := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email); getErr != nil {
		return errors.NewInternalServerError("database error")
	}
	return nil
}

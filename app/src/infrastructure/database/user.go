package database

import (
	"Go-Blog/app/src/domain/model"
	"Go-Blog/app/src/domain/repository"
	"log"

	"database/sql"
)

type userDatabase struct {
	Conn *sql.DB
}

func NewUserDabase(conn *sql.DB) repository.UserRepository {
	return &userDatabase{
		Conn: conn,
	}
}

func (ud *userDatabase) Get(name string) (*model.UserModel, error) {

	rows, err := ud.Conn.Query("select * from users where username = ?", name)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	rows.Next()
	var users model.UserModel

	if err = rows.Scan(&users.Id, &users.UserName, &users.UserEmail, &users.UserPassWord, &users.CreatedAt, &users.UpdatedAt); err != nil {
		return nil, err
	}
	log.Print(&users)
	return &users, err
}

func (ud *userDatabase) GetAll() []model.UserModel {
	var users []model.UserModel
	return users
}

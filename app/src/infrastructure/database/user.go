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

	rows, err := ud.Conn.Query("select * from　users")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var users []model.UserModel
	for rows.Next() {
		user := model.UserModel{}
		if err := rows.Scan(&user.Id, &user.UserName, &user.UserEmail, &user.UserPassWord, &user.CreatedAt, &user.UpdatedAt); err != nil {
			panic(err)
		}
		users = append(users, user)
	}
	return users
}

func (ud *userDatabase) FindByName(name string) (string, error) {

	var password string
	var err error
	if err != ud.Conn.QueryRow("select userpassword from users where username = ? LIMIT 1", name).Scan(&password) {
		return "miss", err
	}

	return password, err
}

func (ud *userDatabase) Add(name string, email string, pass string) error {

	_, err := ud.Conn.Exec("insert into users (username,useremail,userpassword) values (?, ?, ?)", name, email, pass)
	if err != nil {
		return err
	}
	return err
}

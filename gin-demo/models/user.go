package models

import "gin-demo/db"

type User struct {
	Id int `json:"id" from:"id"`
	Username string `json:"username" from:"username"`
}

func (u *User) AddUser() (id int64, err error)  {
	result, err := db.SqlDB.Exec("insert into user(nickname) values (?, ?)", u.Username)
	if err != nil {
		return
	}
	id, err = result.LastInsertId()
	return
}

func (u *User) GetUsers() (users []User, err error) {
	result, err := db.SqlDB.Query("select * from user order by id desc")
	defer result.Close()
	if err != nil {
		return
	}
	for result.Next() {
		user := User{}
		users = append(users, user)
	}
	return

}
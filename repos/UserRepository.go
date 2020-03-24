package repos

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	username string `json: "username"`
	password string `json: "password"`
}

func UserIsValid(username, password string) bool {
	var user User
	_isValid := false

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3307)/zocialeye")
	if err != nil {
		panic(err)
	}
	fmt.Println("connected database!")
	defer db.Close()

	result, err := db.Query("SELECT username, password FROM api_login where member_id = ?", 1)
	if err != nil {
		log.Fatal(err)
	}
	defer result.Close()

	for result.Next() {
		err := result.Scan(&user.username, &user.password)
		if err != nil {
			log.Fatal(err)
		}
		_username := user.username
		_password := user.password

		if username == _username && password == _password {
			_isValid = true
		} else {
			_isValid = false
		}
	}
	err = result.Err()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(_isValid)
	return _isValid
}

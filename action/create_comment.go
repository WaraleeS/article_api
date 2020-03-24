package action

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func create_comment() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3307)/zocialeye")
	if err != nil {
		panic(err)
	}
	fmt.Println("connected database!")
	defer db.Close()

	create, err := db.Query("INSERT INTO api_comment VALUES (0, 'There is always the right tool for the job. For sometimes when your Ruby on Rails system needs a very powerful web handler, think a little outside of the ruby eco-system for simpler yet more powerful alternative solutions.', 0)")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("inserted database api_comment !")
	defer create.Close()
}

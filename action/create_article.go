package action

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func create_article() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3307)/zocialeye")
	if err != nil {
		panic(err)
	}
	fmt.Println("connected database!")
	defer db.Close()

	create, err := db.Query("INSERT INTO api_article VALUES (0, 'I am not very big on the language and framework wars that the interwebs are always fighting about. I believe efficiency, productivity and code maintainability relies mostly on how simple you can architect your solution.', 0)")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("inserted database api_article !")
	defer create.Close()
}

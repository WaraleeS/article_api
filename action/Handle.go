package action

import (
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/waralees/article_api/helper"
	"github.com/waralees/article_api/repos"
)

// func user_login() {
// 	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3307)/zocialeye")
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("connected database!")
// 	defer db.Close()

// 	result, err := db.Query("SELECT username, password FROM api_login where member_id = ?", 1)
// 	if err != nil {
// 		return err
// 	} else {
// 		return result
// 	}
// 	defer result.Close()
// }

// for GET
func LoginPageHandle(response http.ResponseWriter, request *http.Request) {
	var body, _ = helper.LoadFile("templates/login.html")
	fmt.Fprintf(response, body)
}

// for POST
func LoginHandle(response http.ResponseWriter, request *http.Request) {
	username := request.FormValue("username")
	password := request.FormValue("password")
	redirectTarget := "/"
	if !helper.IsEmpty(username) && !helper.IsEmpty(password) {
		_userIsValid := repos.UserIsValid(username, password)
		if _userIsValid {
			redirectTarget = "/"
		}
	}
	http.Redirect(response, request, redirectTarget, 302)
}

package main

import (
	"net/http"

	action "./action"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	echo "github.com/labstack/echo/v4"
)

func server(c echo.Context) error {
	return c.String(http.StatusOK, "simon says be cool!")
}

var router = mux.NewRouter()

func main() {
	// fmt.Println("Test main.go in article_api")

	// // Connect Database
	// db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3307)/zocialeye")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("Connected Database!!")
	// defer db.Close()

	// Connect Web Server with echo
	// e := echo.New()
	// e.GET("/", server)

	// e.Start(":1234")
	router.HandleFunc("/", action.LoginPageHandle)

	http.Handle("/", router)
	http.ListenAndServe(":1234", nil)
}

package main

import (
	"net/http"
	"os"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
	_ "github.com/labstack/echo/middleware"
)

type user struct {
	Id   int64  `json:id`
	Name string `json:name`
}

func gormConnect() *gorm.DB {
	var DBMS = "mysql"
	USER := os.Getenv("MYSQL_USER")
	PASS := os.Getenv("MYSQL_PASSWORD")
	var PROTOCOL = "tcp(mysql:3306)"
	DBNAME := os.Getenv("MYSQL_DATABASE")

	var CONNECT = USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return db
}

func main() {
	e := echo.New()
	db := gormConnect()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!!!")
	})
	e.GET("/users/:id", func(c echo.Context) error {
		userEx := user{}
		user_id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		userEx.Id = user_id
		db.First(&userEx)
		return c.String(http.StatusOK, "User: "+userEx.Name)
	})
	e.Logger.Fatal(e.Start(":5000"))
}

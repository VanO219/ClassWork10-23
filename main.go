package main

import (
	"ClassWork10-23/config"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"log"
	"net/url"
)

type DB struct {
	testDB *sql.DB
}

func (db *DB) Connection(dataPass, username, password string) (err error) {
	defer func() { err = errors.Wrap(err, "main.Connection") }()

	u, err := url.Parse(dataPass)
	if err != nil {
		err = errors.Wrap(err, "ошибка парсинга пути до базы данных")
	}
	u.User = url.UserPassword(username, password)
	u.RawQuery = "sslmode=disable"
	fmt.Println(u.String())
	db.testDB, err = sql.Open("postgres", u.String())
	if err != nil {
		err = errors.Wrap(err, "ошибка подключения к бд")
		return
	}
	return err
}

func (db *DB) Close() (err error) {
	defer func() { err = errors.Wrap(err, "main.Close") }()
	err = db.testDB.Close()
	if err != nil {
		err = errors.Wrap(err, "ошибка закрытия бд")
		return
	}
	return
}

var (
	dbase = DB{}
	conf *config.Config
)

func main() {
	err := manage()
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		err = dbase.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	t := dbase.testDB
	rows, err := t.Query("select * from products")
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var id int64
		var model string
		var company string
		var price int64
		err = rows.Scan(&id, &model, &company, &price)
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Println(id, model, company, price)
	}

}

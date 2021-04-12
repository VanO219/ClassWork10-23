package main

import (
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

func (db *DB) Connection(username, password string) (err error) {
	defer func() { err = errors.Wrap(err, "main.Connection") }()

	u := url.URL{}
	u.Scheme = "postgres"
	u.User = url.UserPassword(username, password)
	u.Host = "localhost:5000"
	u.Path = "test_db"
	u.RawQuery = "sslmode=disable"

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

var dbase = DB{}

func main() {
	var err error
	err = dbase.Connection("testuseri", "testuser123")
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
	rows, err := t.Query("select * from notes")
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var id int64
		var text string
		err = rows.Scan(&id, &text)
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Println(id, text)
	}

}

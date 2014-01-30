package backend

import (
	"fmt"
	"database/sql"
	"github.com/go-sql-driver/mysql"
)

func ConnectDatabase(user string, password string, host string, dbname string) (database *mysql.DB) {
	dsn := user + ":" + password + "@tcp(" + host + ":3306)/" + dbname + "?charset=utf8"

	// Opens the connection with the given DSN
	db, err := mysql.Open(dsn)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Pings the connection, as the Open command doesn't ping
	if err = db.Ping(); err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Connected to MySQL database on " + host + ".")

	return &db
}
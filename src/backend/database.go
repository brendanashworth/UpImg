package backend

import (
	"fmt"
	"database/sql"
	"github.com/go-sql-driver/mysql"
)

type Database struct {
	connection *mysql.MySQLDriver
}

// Connects to the database.
func ConnectDatabase(host string, database string, user string, password string) (db *Database) {
	dsn := user + ":" + password + "@tcp(" + host + ":3306)/" + database + "?charset=utf8"
	// username:password@tcp(example.com:3306)/database?charset=utf8

	// Connect to the database!
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	// Pings the database for connection.
	err = db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return nil
	} else {
		fmt.Println("Successfully connected to MySQL.")
	}

	// Return the database, then don't do much.
	db = &Database{ connection: &db }
	return *db
}

func (db *Database) GetImage(url string) (img string) {
	err := *db.connection.QueryRow("SELECT image FROM `images` WHERE `url`='?' LIMIT 1", url).Scan(&img)

	if err != nil {
		return ""
	} else {
		return img
	}
}

func (db *Database) Close() {
	*db.connection.Close()
}
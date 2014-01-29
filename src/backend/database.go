package backend

import (
	"fmt"
	"database/sql"
	_ "github.com/Go-SQL-Driver/MySQL"
)

type Database struct {
	*connection DB
}

// Connects to the database.
func ConnectDatabase(host string, database string, user string, password string) (*database Database) {
	dsn := user + ":" + password + "@tcp(" + host + ":3306)/" + database + "?charset=utf8"

	// Connect to the database!
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	// Pings the database for connection.
	err := db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return nil
	} else {
		fmt.Println("Successfully connected to MySQL.")
	}

	// Return the database, then don't do much.
	return Database{
		*connection: &db
	}

}

func (*db Database) GetImage(url string) (img []byte) {
	var img string

	err := *db.connection.QueryRow("SELECT image FROM `images` WHERE `url`='?' LIMIT 1", url).Scan(&img)

	if err != nil {
		return nil
	} else {
		return img
	}
}

func (*db Database) Close() {
	*db.db.Close()
}
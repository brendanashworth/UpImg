package backend

import (
	"fmt"
)

type Backend struct {

}

func ConnectDatabase(user string, password string, host string, dbname string) *Backend {
	fmt.Println("Connected to MySQL database on " + host)
	return &Backend{

	}
}

func (backend *Backend) GetImage(url string) (image []byte) {
	return make([]byte, '2')
}

/*func ConnectDatabase(user string, password string, host string, dbname string) (database *sql.DB) {
	dsn := user + ":" + password + "@tcp(" + host + ":3306)/" + dbname + "?charset=utf8"

	// Opens the connection with the given DSN
	db, err := sql.Open("mysql", dsn)
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

	return db
}*/
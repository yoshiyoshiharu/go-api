package db

import(
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
)

var Db *sql.DB

func init() {
	var err error
	// [TODO] Use envs
	cfg := mysql.Config{
		DBName:    "test_db",
		User:      "root",
		Passwd:    "root",
		Addr:      "db:3306",
		Net:       "tcp",
		AllowNativePasswords: true,
	}
	Db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	err = Db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Successfly connected to database.")
}

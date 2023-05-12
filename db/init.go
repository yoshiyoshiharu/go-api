package db

import(
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"os"
	"log"
	"time"
)

func Init() *sql.DB {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
	  log.Fatal(err)
	}

	// [TODO] Use envs
	c := mysql.Config{
		DBName:    "test_db",
		User:      "root",
		Passwd:    "root",
		Addr:      "db:3306",
		Net:       "tcp",
		ParseTime: true,
		Loc:       jst,
		AllowNativePasswords: true,
	}
	db, err := sql.Open("mysql", c.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfly connected to database.")
	}
	return db
}

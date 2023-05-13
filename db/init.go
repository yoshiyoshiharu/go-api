package db

import(
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
	"time"
)

var Db *sql.DB

func init() {
	var err
	jst, err = time.LoadLocation("Asia/Tokyo")
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
	Db, err = sql.Open("mysql", c.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	if err = Db.Ping(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfly connected to database.")
	}
}

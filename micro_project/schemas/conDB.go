package schemas

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Connectdb() (*sql.DB, error) {
	cfg := mysql.Config{
		User:   "root",
		Passwd: "smToey140239",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "mobiledb_sit",
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	// defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
	return db, nil
}

// func Connectdb(cfg *mysql.Config) (*sql.DB, error) {
// cfg := mysql.Config{
// 	User:    "root",
// 	Passwd:  "smToey140239",
// 	Net:     "tcp",
// 	Addr:    "127.0.0.1:3306",
// 	DBName:  "mobiledb_sit",
// 	Timeout: 30 * time.Second,
// }
// Get a database handle.
// 	cfgCopy := cfg
// 	// cfgCopy.Addr = fmt.Sprintf("%s:%d", cfg.Addr, 3306)
// 	var err error
// 	db, err = sql.Open("mysql", cfgCopy.FormatDSN())
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	pingErr := db.Ping()
// 	if pingErr != nil {
// 		log.Fatal(pingErr)
// 	}
// 	fmt.Println("Connected!")
// 	return db, nil
// }

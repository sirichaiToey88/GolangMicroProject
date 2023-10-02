package customer

import (
	"github.com/labstack/echo"
)

var err error

// func CrerateTable(e echo.Context) error {
// 	createTb := `
// 	CREATE TABLE IF NOT EXISTS goimdb (
// 	id INT AUTO_INCREMENT,
// 	imdbID TEXT NOT NULL,
// 	title TEXT NOT NULL,
// 	year INT NOT NULL,
// 	rating FLOAT NOT NULL,
// 	isSuperHero BOOLEAN NOT NULL,
// 	PRIMARY KEY (id)
// 	);`

// 	_, err = db.Exec(createTb)

//		if err != nil {
//			fmt.Println("Creat table")
//			log.Fatal("error create DB", err)
//		}
//		return nil
//	}
func CreateTable(c echo.Context) error {
	// Your implementation here
	return nil
}

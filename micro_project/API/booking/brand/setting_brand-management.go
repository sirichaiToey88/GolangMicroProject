package brand

import (
	"log"
	"net/http"
	bookstadium "project/models/book_stadium"
	"project/other"
	"project/schemas"

	"github.com/labstack/echo/v4"
)

func AddNewBrand(e echo.Context) error {
	// Parse the request body to extract the data to be inserted
	body := &bookstadium.Brand{}
	if err := e.Bind(body); err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}

	// Check if any parameter in the body is null or empty
	if body.Brand_title == "" || body.Image_url == "" || body.Time_close == "" || body.Time_open == "" || body.Address == "" {
		return e.JSON(http.StatusBadRequest, "One or more parameters are empty.")
	}

	db, err := schemas.Connectdb()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//Insert statement
	stmt, err := db.Prepare("INSERT INTO brand(id,brand_title,image_url,time_open,time_close,location,tell,address,create_date,modify_date) VALUES(?,?,?,?,?,?,?,?,?,?);")
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	defer stmt.Close()

	// Execute the INSERT statement
	UUID := other.GenerateUUID()
	Createdate := other.DateTime()
	Modifydate := other.DateTime()
	result, err := stmt.Exec(UUID, body.Brand_title, body.Image_url, body.Time_open, body.Time_close, body.Location, body.Tell, body.Address, Createdate, Modifydate)

	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	} else {
		id, _ := result.LastInsertId()
		body.Id = string(id)
		return e.JSON(http.StatusCreated, body)
	}
}

func AddStaduimToBrand(e echo.Context) error {
	body := &bookstadium.Stadium{}
	if err := e.Bind(body); err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}

	// Check if any parameter in the body is null or empty
	if body.Brand_id == "" || body.Stadium_number == "" {
		return e.JSON(http.StatusBadRequest, "One or more parameters are empty.")
	}

	db, err := schemas.Connectdb()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//Query Brand for get time_open and time_close
	rBrand := bookstadium.Brand{}
	query := "SELECT * FROM brand WHERE id=?"

	row := db.QueryRow(query, body.Brand_id)

	err = row.Scan(&rBrand.Id, &rBrand.Brand_title, &rBrand.Image_url, &rBrand.Time_open, &rBrand.Time_close, &rBrand.Location, &rBrand.Tell, &rBrand.Address, &rBrand.Create_date, &rBrand.Modify_date)

	if err != nil {
		return e.JSON(http.StatusNotFound, "Brand ID not found.")
	}

	//Insert statement
	stmt, err := db.Prepare("INSERT INTO stadium(id,brand_id,type_stadium,status,image_url,promotion,price,create_date,modify_date,stadium_number) VALUES(?,?,?,?,?,?,?,?,?,?);")

	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	defer stmt.Close()

	// Execute the INSERT statement
	UUID := other.GenerateUUID()
	Createdate := other.DateTime()
	Modifydate := other.DateTime()

	//Query Stadium_number for check duplicate
	rStadiumNumber := bookstadium.Stadium{}
	queryStadiumNumber := "SELECT * FROM stadium WHERE brand_id=? and stadium_number=? and type_stadium=?"

	rowStadiumNumber := db.QueryRow(queryStadiumNumber, body.Brand_id, body.Stadium_number, body.Type_stadium)

	err = rowStadiumNumber.Scan(&rStadiumNumber.Id, &rStadiumNumber.Brand_id, &rStadiumNumber.Type_stadium, &rStadiumNumber.Status, &rStadiumNumber.Image_url, &rStadiumNumber.Promotion, &rStadiumNumber.Price, &rStadiumNumber.Create_date, &rStadiumNumber.Modify_date,&rStadiumNumber.Stadium_number)

	if err != nil {
		result, err := stmt.Exec(UUID, body.Brand_id, body.Type_stadium, "0", body.Image_url, body.Promotion, body.Price, Createdate, Modifydate, body.Stadium_number)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, err.Error())
		} else {
			id, _ := result.LastInsertId()
			body.Id = string(id)
			return e.JSON(http.StatusCreated, body)
		}
	} else {
		return e.JSON(http.StatusNotFound, "Stadium number it duplicate.")

	}

}

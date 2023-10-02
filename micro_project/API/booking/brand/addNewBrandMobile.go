package brand

import (
	"fmt"
	"log"
	"net/http"
	bookstadium "project/models/book_stadium"
	"project/other"
	"project/schemas"

	"github.com/labstack/echo/v4"
)

func AddNewBrandMobile(e echo.Context) error {
	body := []bookstadium.BrandData{}
	if err := e.Bind(&body); err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}

	if len(body) == 0 {
		return e.JSON(http.StatusBadRequest, "Brand data is empty.")
	}

	db, err := schemas.Connectdb()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Printf("Received Brand Data:\n%+v\n", body)
	for _, body := range body {
		// Check if any parameter in the body is null or empty
		if body.Brand_title == "" || body.Time_close == "" || body.Time_open == "" || body.Address == "" {
			return e.JSON(http.StatusBadRequest, "One or more parameters are empty.")
		}

		if len(body.Stadiums) == 0 {
			return e.JSON(http.StatusBadRequest, "Stadiums data is empty.")
		}

		for _, stadium := range body.Stadiums {
			if stadium.Type_stadium == "" || stadium.Stadium_number == "" || stadium.Price == "" {
				return e.JSON(http.StatusBadRequest, "One or more stadium parameters are empty.")
			}
		}

		//Insert brand statement
		stmt, err := db.Prepare("INSERT INTO brand(id,brand_title,image_url,time_open,time_close,location,tell,address,create_date,modify_date) VALUES(?,?,?,?,?,?,?,?,?,?);")
		if err != nil {
			return e.JSON(http.StatusInternalServerError, err.Error())
		}
		defer stmt.Close()

		// Execute the INSERT statement
		UUID := other.GenerateUUID()
		Createdate := other.DateTime()
		Modifydate := other.DateTime()
		result, errBrand := stmt.Exec(UUID, body.Brand_title, "https://static.vecteezy.com/system/resources/previews/014/589/588/original/stadium-icon-in-simple-style-vector.jpg", body.Time_open, body.Time_close, body.Location, body.Tell, body.Address, Createdate, Modifydate)

		stmtStadium, err := db.Prepare("INSERT INTO stadium(id, brand_id, type_stadium, status, image_url, promotion, price, create_date, modify_date, stadium_number) VALUES(?,?,?,?,?,?,?,?,?,?);")
		if err != nil {
			fmt.Print("Error insert brand", err)
			return err
		}
		defer stmt.Close()

		// พิมพ์ข้อมูล Stadiums
		for _, stadiumData := range body.Stadiums {
			UUIDStadium := other.GenerateUUID()
			// เรียกใช้งานฟังก์ชันที่จะแทรกข้อมูล Stadium
			_, err = stmtStadium.Exec(UUIDStadium, UUID, stadiumData.Type_stadium, "0", "https://static.vecteezy.com/system/resources/previews/014/589/588/original/stadium-icon-in-simple-style-vector.jpg", stadiumData.Promotion, stadiumData.Price, Createdate, Modifydate, stadiumData.Stadium_number)
			if err != nil {
				fmt.Print("Error insert stadium", err)
				return err
			}
			if err != nil {
				fmt.Print("Error insert stadium", err)
				// หากเกิดข้อผิดพลาดในการแทรกข้อมูล Stadium ให้คืนข้อความผิดพลาดกลับไป
				return e.JSON(http.StatusInternalServerError, err.Error())
			}
		}

		if errBrand != nil {
			return e.JSON(http.StatusInternalServerError, err.Error())
		} else {
			id, _ := result.LastInsertId()
			body.Id = string(id)

		}
	}
	return e.JSON(http.StatusCreated, body)
	// for _, stadium := range body.Stadiums {
	// 	fmt.Printf("Type Stadium: %s\n", stadium.Type_stadium)
	// 	fmt.Printf("Stadium Number: %s\n", stadium.Stadium_number)
	// 	fmt.Printf("Price: %s\n", stadium.Price)
	// }

}

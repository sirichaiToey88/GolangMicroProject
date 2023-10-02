package booking

import (
	"log"
	"net/http"
	bookstadium "project/models/book_stadium"
	"project/schemas"

	"github.com/labstack/echo/v4"
)

func ListAllBrand(e echo.Context) error {
	// Parse the request body to extract the data to be inserted
	body := &bookstadium.Brand{}
	if err := e.Bind(body); err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}

	db, err := schemas.Connectdb()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//Query Brand and stadium 
	rBrand := bookstadium.Brand{}
	query := "select * from brand as b inner join stadium as s on s.brand_id = b.id order by b.id"

	rows, err := db.Query(query)

	if err != nil {
		log.Fatal("error executing query:", err)
	}
	defer rows.Close()

	brandWithStadium := make(map[string]map[string]interface{})
	for rows.Next() {
		if err := rows.Scan(&rBrand.Id, &rBrand.Brand_title, &rBrand.Image_url, &rBrand.Time_open, &rBrand.Time_close, &rBrand.Location, &rBrand.Tell, &rBrand.Address, &rBrand.Create_date, &rBrand.Modify_date,
			&rBrand.Stadium.Id, &rBrand.Stadium.Brand_id, &rBrand.Stadium.Type_stadium, &rBrand.Stadium.Status, &rBrand.Stadium.Image_url,
			&rBrand.Stadium.Promotion, &rBrand.Stadium.Price, &rBrand.Stadium.Create_date, &rBrand.Stadium.Modify_date, &rBrand.Stadium.Stadium_number); err != nil {
			log.Fatal("error select into DB", err)
		}
		// Organize data for return to mobile
		brandID := rBrand.Stadium.Brand_id
		stadiumMap := map[string]interface{}{
			"id":             rBrand.Stadium.Id,
			"brand_id":       rBrand.Stadium.Brand_id,
			"type_stadium":   rBrand.Stadium.Type_stadium,
			"status":         rBrand.Stadium.Status,
			"image_url":      rBrand.Stadium.Image_url,
			"stadium_number": rBrand.Stadium.Stadium_number,
			"promotion":      rBrand.Stadium.Promotion,
			"price":          rBrand.Stadium.Price,
		}

		// Check if the brand with the given ID already exists in the result
		if _, ok := brandWithStadium[brandID]; !ok {
			brandWithStadium[brandID] = map[string]interface{}{
				"id":          rBrand.Id,
				"brand_title": rBrand.Brand_title,
				"image_url":   rBrand.Image_url,
				"time_open":   rBrand.Time_open,
				"time_close":  rBrand.Time_close,
				"location":    rBrand.Location,
				"tell":        rBrand.Tell,
				"address":     rBrand.Address,
				"stadium":     map[string]interface{}{},
			}
		}
		brandWithStadium[brandID]["stadium"].(map[string]interface{})[rBrand.Stadium.Id] = stadiumMap
	}
	return e.JSON(http.StatusOK, brandWithStadium)
}

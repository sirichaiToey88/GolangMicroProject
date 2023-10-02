package bookstadium

type Brand struct {
	Id          string  `json:"id"`
	Brand_title string  `json:"brand_title"`
	Image_url   string  `json:"image_url"`
	Time_open   string  `json:"time_open"`
	Time_close  string  `json:"time_close"`
	Location    string  `json:"location"`
	Tell        string  `json:"tell"`
	Address     string  `json:"address"`
	Create_date string  `json:"create_date"`
	Modify_date string  `json:"modify_date"`
	Stadium     Stadium `json:"stadium"`
}

type Stadium struct {
	Id             string `json:"id"`
	Brand_id       string `json:"brand_id"`
	Type_stadium   string `json:"type_stadium"`
	Status         string `json:"status"`
	Image_url      string `json:"image_url"`
	Stadium_number string `json:"stadium_number"`
	Promotion      string `json:"promotion"`
	Price          string `json:"price"`
	Create_date    string `json:"create_date"`
	Modify_date    string `json:"modify_date"`
}

type Booking struct {
	Id                string          `json:"id"`
	User_id           string          `json:"user_id"`
	Brand_id          string          `json:"brand_id"`
	Stadium_id        string          `json:"stadium_id"`
	Reservation_hours string          `json:"reservation_hours"`
	Reservation_date  string          `json:"reservation_date"`
	Start_time        string          `json:"start_time"`
	End_time          string          `json:"end_time"`
	Create_date       string          `json:"create_date"`
	Modify_date       string          `json:"modify_date"`
	Del               string          `json:"del"`
	Status_payment    string          `json:"status_payment"`
	Brand_title       string          `json:"brand_title"`
	Stadium_number    string          `json:"stadium_number"`
	Payment_success   string          `json:"payment_success"`
	Total_Price       string          `json:"total_price"`
	NeedToPay         string          `json:"need_to_pay"`
	OverDue           string          `json:"over_due"`
	Payment           Stadium_payment `json:"Payment"`
}

type Stadium_payment struct {
	Id          string `json:"id"`
	User_id     string `json:"user_id"`
	Booking_id  string `json:"booking_id"`
	Status      string `json:"status"`
	Image_url   string `json:"image_url"`
	Total       string `json:"total"`
	Create_date string `json:"create_date"`
	Modify_date string `json:"modify_date"`
	Source      string `json:"source"`
	Destination string `json:"destination"`
}

type BrandData struct {
	Id          string `json:"id"`
	Brand_title string `json:"brand_title"`
	Image_url   string `json:"image_url"`
	Time_open   string `json:"time_open"`
	Time_close  string `json:"time_close"`
	Location    string `json:"location"`
	Tell        string `json:"tell"`
	Address     string `json:"address"`
	Create_date string `json:"create_date"`
	Modify_date string `json:"modify_date"`
	Stadiums    map[string]struct {
		Id             string `json:"id"`
		Brand_id       string `json:"brand_id"`
		Status         string `json:"status"`
		Type_stadium   string `json:"type_stadium"`
		Image_url      string `json:"image_url"`
		Stadium_number string `json:"stadium_number"`
		Promotion      string `json:"promotion"`
		Price          string `json:"price"`
		Create_date    string `json:"create_date"`
		Modify_date    string `json:"modify_date"`
	} `json:"stadiums"`
}

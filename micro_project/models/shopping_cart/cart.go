package shoppingcart

type Payment struct {
	Account     string `json:"account"`
	Source      string `json:"source"`
	Distination string `json:"distination"`
	Total       string `json:"total"`
}

type Shopping struct {
	User_id          string  `json:"User_id"`
	Product_id       string  `json:"Product_id"`
	Product_title    string  `json:"Product_title"`
	Product_price    string  `json:"Product_price"`
	Total_amount     string  `json:"Total_amount"`
	Quantity_product string  `json:"Quantity_product"`
	Payment_type     string  `json:"Payment_type"`
	Image_url        string  `json:"Image_url"`
	Payment          Payment `json:"Payment"`
}

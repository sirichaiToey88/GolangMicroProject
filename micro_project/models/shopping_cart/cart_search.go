package shoppingcart

type PaymentSearch struct {
	Id          string `json:"id"`
	Order_id    string `json:"order_id"`
	Account     string `json:"account"`
	Source      string `json:"source"`
	Distination string `json:"distination"`
	Total       string `json:"total"`
	Create_date string `json:"create_date"`
	Modify_date string `json:"modify_date"`
}

type ShoppingSearch struct {
	Id               string        `json:"id"`
	Order_id         string        `json:"order_id"`
	User_id          string        `json:"User_id"`
	Product_id       string        `json:"Product_id"`
	Product_title    string        `json:"Product_title"`
	Product_price    string        `json:"Product_price"`
	Total_amount     string        `json:"Total_amount"`
	Quantity_product string        `json:"Quantity_product"`
	Payment_type     string        `json:"Payment_type"`
	Image_url        string        `json:"Image_url"`
	Payment          PaymentSearch `json:"Payment"`
	Create_date      string        `json:"create_date"`
	Modify_date      string        `json:"modify_date"`
}

type MainCart struct {
	Id           string `json:"id"`
	Order_id     string `json:"order_id"`
	User_id      string `json:"User_id"`
	Total_amount string `json:"Total_amount"`
	Payment_type string `json:"Payment_type"`
	Create_date  string `json:"create_date"`
	Modify_date  string `json:"modify_date"`
}

package params

type InsertTransaction struct {
	Fullname string  `json:"fullname"`
	Quantity int     `json:"quantity"`
	Price    string `json:"price"`
}

type Response struct {
	Status  int         `json:"status"`
	Payload interface{} `json:"payload"`
}

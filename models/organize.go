package models

type CreateShow struct {
	Username   string `json:"username"`
	Show       string `json:"show"`
	Price      string `json:"price"`
	Quantity   string `json:"quantity"`
	TicketsRem int    `json:"tickets_rem"`
}

type NFTids struct {
	Ids []string `json:"ids"`
}

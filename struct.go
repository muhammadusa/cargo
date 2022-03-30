package main

type PostInformation struct {
	Sender     SenderInformation
	Receiver   RecieverInformation
	Box        Box
	Items      []Item
	TotalValue TotalValueItem
}


type Box struct {
	FirstName            string `json:"firstName"`
	LastName             string `json:"lastName"`
	PassportSerialNumber string `json:"passportSerialNumber"`
	MainPhone            string `json:"mainPhone"`
	Address              string `json:"address"`
	City                 string `json:"city"`
	Country              string `json:"country"`
}

type Item struct {
	ItemName        string `json:"itemName"`
	Quantity        string    `json:"quantity"`
}

type TotalValueItem struct {
	TotalValueItem      int   `json:"totalValueItem"`
}

type RecieverInformation struct {
	ReceiverId           int   `json:"receiverId"`
}

type SenderInformation struct {
	SenderId         int    `json:"senderId"`
	FirstName        string `json:"firstName"`
	LastName         string `json:"lastName"`
	MainPhone        string `json:"mainPhone"`
	Email            string `json:"email"`
	Address          string `json:"address"`
	City             string `json:"city"`
	Country          string `json:"country"`
	PostCode         int    `json:"postCode"`
}

type PostedInformation struct {
	SenderId          int  `json:"senderId"`
	ReceiverId        int   `json:"receiverId"`
	BoxId             int   `json:"boxId"`
}
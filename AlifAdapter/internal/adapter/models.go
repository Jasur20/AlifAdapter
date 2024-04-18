package adapter

type Resp struct{
	ID string `json:"id"`
	DataTime string `json:"datatime"`
	Code int `json:"code"`
	Message string `json:"message"`
	Status string `json:"status"`
	StatusCode int `json:"statusCode"`
	Amount string `json:"amount"`
	Fx string `json:"fx"`
	ToPay interface{} `json:"topay"`
	AccountInfo string `json:"accountinfo"`
}

type Req struct{
	Service string `json:"service"`
	UserID string `json:"userid"`
	Hash string `json:"hash"`
	Account string `json:"account"`
	Amount string `json:"amount"`
	Currency string `json:"currency"`
	TxnID string `json:"txnid"`
	Phone string `json:"phone"`
	Fee string `json:"fee"`
	Providerid string `json:"proviredid"`
	Last_Name string `json:"last_name"`
	First_Name string `json:"first_name"`
	Middle_Name string `json:"middle_name"`
	Sender_Birthday string`json:"sender_birhtday"`
	Address string `json:"address"`
	Resident_Country string `json:"resident_country"`
	Postal_Code string `json:"postal_code"`
	Recipient_Name string`json:"recipient_name"`
}
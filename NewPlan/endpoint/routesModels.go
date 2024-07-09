package endpoint


type Take struct{
	Number string `json:"number"`
	Language string `json:"language"`
	Insult string `json:"insult"`
	Created string `json:"created"`
	Shown string `json:"shown"`
	CreatedBy string `json:"createdby"`
	Active string `json:"active"`
	Comment string `json:"comment"`
}
type Res struct{
	Err []string `json:"error"`
	Result

}
type Result struct{
	Status string `json:"status"`
	TimeStamp string `json:"timestamp"`
}

type Req1 struct{
	GeckoSays string `json:"gecko_says"`
}
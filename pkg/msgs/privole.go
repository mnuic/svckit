package msgs

type UvjetiKoristenjaReq struct {
	IgracId  string `json:"igrac_id"`
	Verzija  int    `json:"verzija"`
	RemoteIP string `json:"remote_ip"`
}

type NewsletterPostavkeReq struct {
	IgracId  string `json:"igrac_id"`
	Status   int    `json:"status"`
	RemoteIP string `json:"remote_ip"`
}

type SMSPostavkeReq struct {
	IgracId  string `json:"igrac_id"`
	Status   int    `json:"status"`
	RemoteIP string `json:"remote_ip"`
}

package domain

//type DocumentNumber struct {
//}

type Beneficiary struct {
	Id             int    `json:"id"`
	Status         string `json:"status"`
	Name           string `json:"name"`
	DocumentNumber string `json:"document_number"`
	Email          string `json:"email"`
	PixKeyType     string `json:"pix_key_type"`
	PixKeyValue    string `json:"pix_key_value"`
	Bank           string `json:"bank"`
	Agency         string `json:"agency"`
	Account        string `json:"account"`
}

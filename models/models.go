package models

//RetrieveSignatureRequest -
type RetrieveSignatureRequest struct {
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	IP      string `json:"ip"`
	Date    string `json:"date"`
}

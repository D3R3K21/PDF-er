package server

import (
	service "PDF-er/pdf-service"
	"bytes"
	"io"
	"net/http"
)

//RetrieveSignatureRequest -
type RetrieveSignatureRequest struct {
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	IP      string `json:"ip"`
	Date    string `json:"date"`
}

//SignatureFunc -
func SignatureFunc(response http.ResponseWriter, r *http.Request) {
	//localhost/signature?name=derek&email=drose@integrate.com&ip=192.168.1.1&company=integrate inc
	response.Header().Set("Content-Disposition", "attachment; filename=sigFile.pdf")
	response.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	t := RetrieveSignatureRequest{
		Company: r.URL.Query().Get("company"),
		Date:    r.URL.Query().Get("date"),
		Email:   r.URL.Query().Get("email"),
		IP:      r.URL.Query().Get("ip"),
		Name:    r.URL.Query().Get("name"),
	}
	pdf := service.NewSignature(t.Name, t.Company, t.Email, t.IP, t.Date)
	var buf bytes.Buffer
	pdf.Output(&buf)
	var reader io.Reader
	bts := buf.Bytes()
	reader = bytes.NewReader(bts)
	io.Copy(response, reader)
	response.WriteHeader(200)
}

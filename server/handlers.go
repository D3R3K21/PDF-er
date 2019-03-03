package server

import (
	"bytes"
	"integrate-pdf-service/models"
	service "integrate-pdf-service/pdf-service"
	"io"
	"net/http"
)

//SignatureFunc - localhost/signature?name=derek&email=drose@integrate.com&ip=192.168.1.1&company=integrate inc
func SignatureFunc(response http.ResponseWriter, r *http.Request) {
	response.Header().Set("Content-Disposition", "attachment; filename=sigFile.pdf")
	response.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	requestModel := models.RetrieveSignatureRequest{
		Company: r.URL.Query().Get("company"),
		Date:    r.URL.Query().Get("date"),
		Email:   r.URL.Query().Get("email"),
		IP:      r.URL.Query().Get("ip"),
		Name:    r.URL.Query().Get("name"),
	}
	pdf := service.GenerateSignature(requestModel)
	var buf bytes.Buffer
	pdf.Output(&buf)
	var reader io.Reader
	bts := buf.Bytes()
	reader = bytes.NewReader(bts)
	io.Copy(response, reader)
	response.WriteHeader(200)
}

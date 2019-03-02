package main

import (
	"log"

	"github.com/jung-kurt/gofpdf"
)

func main() {

	var name = "Derek Rose"
	var company = "Integrate Inc"
	var email = "drose@integrate.com"
	var ip = "127.0.0.1"
	var date = "March 1 2019"
	sig := newSignature(name, company, email, ip, date)
	err := savePDF(sig)
	if err != nil {
		log.Fatalf("Cannot save PDF: %s|n", err)
	}
}
func savePDF(pdf *gofpdf.Fpdf) error {
	return pdf.OutputFileAndClose("tstSigs.pdf")
}

// ## The Initial PDF document
func newSignature(name, company, email, ip, date string) *gofpdf.Fpdf {

	pdf := gofpdf.New("L", "pt", "Letter", "")
	pdf.AddPage()
	pdf.SetFont("Times", "B", 10)
	var imgWidth float64 = 141
	var imgHeight float64 = 33
	docW, docH := pdf.GetPageSize()
	centerW := ((docW / 2) - (imgWidth / 2)) - 50
	centerH := ((docH / 2) - (imgHeight / 2)) - 50
	cellCenter := 10 + centerW
	pdf.ImageOptions("integrate.png", centerW, centerH, imgWidth, imgHeight, false, gofpdf.ImageOptions{ImageType: "png", ReadDpi: true}, 0, "")
	pdf.Ln(centerH + 15)
	pdf.Cell(cellCenter, 10, " ")
	pdf.Cell(40, 10, "Signee Name : "+name)
	pdf.Ln(25)
	pdf.Cell(cellCenter, 10, " ")
	pdf.Cell(40, 10, "Signee Company : "+company)
	pdf.Ln(25)
	pdf.Cell(cellCenter, 10, " ")
	pdf.Cell(40, 10, "Signee Email : "+email)
	pdf.Ln(25)
	pdf.Cell(cellCenter, 10, " ")
	pdf.Cell(40, 10, "Signee IP : "+ip)
	pdf.Ln(25)
	pdf.Cell(cellCenter, 10, " ")
	pdf.Cell(40, 10, "Signed Date : "+date)
	pdf.Ln(25)
	pdf.Cell(cellCenter, 10, " ")

	return pdf
}

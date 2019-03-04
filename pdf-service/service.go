package service

import (
	"integrate-pdf-service/models"

	"github.com/jung-kurt/gofpdf"
)

//NewSignature - Generate the PDF Signautre using the supplied paramters
func GenerateSignature(model models.RetrieveSignatureRequest) *gofpdf.Fpdf {
	pdf := gofpdf.New("L", "pt", "Letter", "")
	pdf.AddPage()
	pdf.SetFont("Times", "B", 10)
	var imgWidth float64 = 141
	var imgHeight float64 = 33
	docW, docH := pdf.GetPageSize()
	centerW := ((docW / 2) - (imgWidth / 2)) - 50
	centerH := ((docH / 2) - (imgHeight / 2)) - 50
	cellCenter := 10 + centerW
	pdf.ImageOptions("./integrate.png", centerW, centerH, imgWidth, imgHeight, false, gofpdf.ImageOptions{ImageType: "png", ReadDpi: true}, 0, "")
	pdf.Ln(centerH + 15)
	pdf.Cell(cellCenter, 10, " ")
	pdf.Cell(40, 10, "Signee Name : "+model.Name)
	pdf.Ln(25)
	pdf.Cell(cellCenter, 10, " ")
	pdf.Cell(40, 10, "Signee Company : "+model.Company)
	pdf.Ln(25)
	pdf.Cell(cellCenter, 10, " ")
	pdf.Cell(40, 10, "Signee Email : "+model.Email)
	pdf.Ln(25)
	pdf.Cell(cellCenter, 10, " ")
	pdf.Cell(40, 10, "Signee IP : "+model.IP)
	pdf.Ln(25)
	pdf.Cell(cellCenter, 10, " ")
	pdf.Cell(40, 10, "Signed Date : "+model.Date)
	pdf.Ln(25)
	pdf.Cell(cellCenter, 10, " ")
	return pdf
}

package calc

import (
	"fmt"
	"github.com/gioapp/gel"
	"github.com/jung-kurt/gofpdf"
	"github.com/marcetin/wingcal/pkg/latcyr"
	"github.com/skratchdot/open-golang/open"
)

var (
	stampajDugme = new(gel.Button)
)

func (w *WingCal) Stampa() func() {
	return func() {
		if len(w.Suma.Elementi) != 0 {
			btn := w.Tema.Button(latcyr.C("Štampaj", w.Cyr))
			for stampajDugme.Clicked(w.Context) {
				///////////////////////////////////////////
				//	pdf := gofpdf.New("P", "mm", "A4", "")
				//	pdf.AddPage()
				//	pdf.SetFont("Arial", "B", 12)
				//
				//	pdf.SetHeaderFunc(func(){
				//		pdf.Cell(40, 10, latcyr.C("Header tetst ", w.Cyr))
				//	})
				//
				//	pdf.SetFooterFunc(func(){
				//		pdf.Cell(40, 10, latcyr.C("Footer tetst ", w.Cyr))
				//	})
				//
				//	for _, e := range w.Suma.Elementi {
				//		pdf.Cell(40, 10, latcyr.C(e.Element.Naziv, w.Cyr))
				//		pdf.Ln(8)
				//		pdf.Cell(40, 10, latcyr.C("Jedinična cena", w.Cyr))
				//		pdf.Cell(40, 10, fmt.Sprintf("%.2f", e.Element.Cena))
				//		pdf.Cell(40, 10, latcyr.C("Jedinična cena", w.Cyr))
				//		pdf.Ln(8)
				//		pdf.Cell(40, 10, latcyr.C("Količina", w.Cyr))
				//		pdf.Cell(40, 10, fmt.Sprint(e.Kolicina))
				//		pdf.Cell(40, 10, e.Element.Jedinica)
				//		pdf.Cell(40, 10, latcyr.C("Cena suma:", w.Cyr))
				//		pdf.Cell(40, 10, fmt.Sprintf("%.2f", e.SumaCena))
				//		pdf.Ln(8)
				//	}
				//
				//	pdf.Ln(16)
				//	pdf.Cell(40, 10, latcyr.C("Ukupan Neophodan Materijal:", w.Cyr))
				//
				//	for _, m := range w.Suma.UkupanNeophodanMaterijal {
				//		pdf.Cell(40, 10, latcyr.C(m.Materijal.Naziv, w.Cyr))
				//		pdf.Ln(8)
				//		pdf.Cell(40, 10, latcyr.C("Jedinična cena", w.Cyr))
				//		pdf.Cell(40, 10, fmt.Sprint(m.Materijal.Cena))
				//		pdf.Cell(40, 10, m.Materijal.Jedinica)
				//
				//		pdf.Ln(8)
				//		pdf.Cell(40, 10, latcyr.C("Ukupno", w.Cyr))
				//		pdf.Cell(40, 10, fmt.Sprint(m.UkupnoPakovanja))
				//		pdf.Cell(40, 10, latcyr.C("Ukupna cena", w.Cyr))
				//		pdf.Cell(40, 10, fmt.Sprint(m.UkupnaCena))
				//
				//	}

				//pdf.Ln(16)
				//pdf.Cell(40, 10, latcyr.C("Suma:", w.Cyr))
				//pdf.Cell(40, 10, fmt.Sprint(w.Suma.SumaCena))
				//

				/////////////////////////////
				pdf := gofpdf.New("P", "", "", "")
				pdf.SetTopMargin(30)

				marginCell := 2. // margin of top/bottom of cell
				pagew, pageh := pdf.GetPageSize()
				mleft, mright, _, mbottom := pdf.GetMargins()

				pdf.SetHeaderFuncMode(func() {
					//pdf.Image("/usr/home/marcetin/Public/wingcal/NOVOGUI/pdfheader.png", 5, 5, 200, 25, false, "", 0, "")
					//pdf.SetDrawColor(200,200,200)
					pdf.SetFillColor(200, 200, 200)
					pdf.Rect(5, 5, 200, 20, "F")
					pdf.SetY(5)
					pdf.SetFont("Arial", "", 8)
					pdf.CellFormat(47, 6, "MB:20701005", "0", 0, "L", false, 0, "")
					pdf.SetFont("Arial", "B", 10)
					pdf.CellFormat(47, 10, "W-ing Solutions D.o.o.", "0", 0, "R", false, 0, "")
					pdf.SetFont("Arial", "", 8)
					pdf.CellFormat(47, 8, "     SIFRA PROJEKTA", "0", 0, "L", false, 0, "")
					pdf.SetFont("Arial", "", 8)
					pdf.CellFormat(47, 8, "fhe38833", "0", 0, "R", false, 0, "")
					pdf.Ln(5)
					pdf.SetFont("Arial", "", 8)
					pdf.CellFormat(47, 6, "PIB:106892584", "0", 0, "L", false, 0, "")
					pdf.CellFormat(47, 8, "Bulevar Oslobodjenja 30A", "0", 0, "R", false, 0, "")
					pdf.SetFont("Arial", "", 8)
					pdf.CellFormat(47, 8, "     NAZIV PROJEKTA", "0", 0, "L", false, 0, "")
					pdf.SetFont("Arial", "", 8)
					pdf.CellFormat(47, 8, "projekat za evidentiranje", "0", 0, "R", false, 0, "")
					pdf.Ln(5)
					pdf.SetFont("Arial", "", 8)
					pdf.CellFormat(47, 6, "tel:069/222-44-33", "0", 0, "L", false, 0, "")
					pdf.CellFormat(47, 8, "21000 Novi Sad", "0", 0, "R", false, 0, "")
					pdf.SetFont("Arial", "", 8)
					pdf.CellFormat(47, 8, "     DATUM PROJEKTA", "0", 0, "L", false, 0, "")
					pdf.SetFont("Arial", "", 8)
					pdf.CellFormat(47, 10, "mart 2020", "0", 0, "R", false, 0, "")
					pdf.Ln(5)
					pdf.SetFont("Arial", "", 8)
					pdf.CellFormat(47, 6, "email:vukobrat.cedomir@gmail.com", "0", 0, "L", false, 0, "")

				}, true)
				pdf.SetFooterFunc(func() {
					pdf.SetY(-15)
					pdf.SetFont("Arial", "I", 8)
					pdf.CellFormat(0, 10, fmt.Sprintf("Strana %d/{nb}", pdf.PageNo()),
						"", 0, "C", false, 0, "")
				})
				pdf.AliasNbPages("")

				pdf.AddPage()
				pdf.SetFont("Times", "B", 16)
				pdf.CellFormat(0, 10, latcyr.C("Specifikacija aktivnosti", w.Cyr), "0", 0, "", false, 0, "")
				pdf.Ln(20)

				pdf.SetFont("Arial", "", 10)
				for _, e := range w.Suma.Elementi {
					cols := []float64{40, pagew - mleft - mright - 20}
					//rows := [][]string{}

					rows := [][]string{
						[]string{
							"Šifra", e.Sifra,
						},
						[]string{
							"Opis", e.Element.Opis,
						},
						[]string{
							"Jedinica mere", e.Element.Jedinica,
						},
						[]string{
							"Jedinična cena", fmt.Sprint(e.Element.Cena),
						},
						[]string{
							"Količina", fmt.Sprint(e.Kolicina),
						},
						[]string{
							"Vrednost rada", fmt.Sprintf("%.2f", e.SumaCena),
						},
					}
					for _, row := range rows {
						curx, y := pdf.GetXY()
						x := curx
						height := 0.
						_, lineHt := pdf.GetFontSize()
						for i, txt := range row {
							lines := pdf.SplitLines([]byte(txt), cols[i])
							h := float64(len(lines))*lineHt + marginCell*float64(len(lines))
							if h > height {
								height = h
							}
						}
						// add a new page if the height of the row doesn't fit on the page
						if pdf.GetY()+height > pageh-mbottom {
							pdf.AddPage()
							y = pdf.GetY()
						}
						for i, txt := range row {
							width := cols[i]
							//pdf.Rect(x, y, width, height, "")
							pdf.MultiCell(width, lineHt+marginCell, txt, "", "", false)
							x += width
							pdf.SetXY(x, y)
						}
						pdf.SetXY(curx, y+height)
					}
					pdf.Ln(8)
				}

				pdf.SetFont("Times", "B", 16)
				pdf.CellFormat(0, 10, latcyr.C("Suma: ", w.Cyr)+fmt.Sprintf("%.2f", w.Suma.SumaCena), "0", 0, "", false, 0, "")
				pdf.Ln(40)

				/////////////////////////
				//for _, e := range w.Suma.Elementi {
				//	pdf.CellFormat(0, 10, latcyr.C(e.Element.Naziv, w.Cyr), "", 1, "", false, 0, "")
				//	pdf.Ln(8)
				//	pdf.Cell(40, 10, latcyr.C("Opis", w.Cyr))
				//
				//	_, lineHt := pdf.GetFontSize()
				//
				//	lines := pdf.SplitLines([]byte(latcyr.C(e.Element.Opis, w.Cyr)), 120)
				//	//ht := float64(len(lines)) * lineHt
				//	for _, line := range lines {
				//		pdf.CellFormat(0, lineHt, string(line), "", 1, "L", false, 0, "")
				//	}
				//
				//	//pdf.CellFormat(50, 10, latcyr.C(e.Element.Opis, w.Cyr), "", 5, "", false, 0, "")
				//	pdf.Ln(8)
				//	pdf.Cell(40, 10, latcyr.C("Jedinična cena", w.Cyr))
				//	pdf.Cell(40, 10, fmt.Sprintf("%.2f", e.Element.Cena))
				//	pdf.Ln(8)
				//	pdf.Cell(40, 10, latcyr.C("Jedinica mere", w.Cyr))
				//	pdf.Cell(40, 10, e.Element.Jedinica)
				//	pdf.Ln(8)
				//	pdf.Cell(40, 10, latcyr.C("Količina", w.Cyr))
				//	pdf.Cell(40, 10, fmt.Sprint(e.Kolicina))
				//	pdf.Ln(8)
				//	pdf.Cell(40, 10, latcyr.C("Vrednost rada:", w.Cyr))
				//	pdf.Cell(40, 10, fmt.Sprintf("%.2f", e.SumaCena))
				//	pdf.Ln(24)
				//}
				/////
				//for _, e := range w.Suma.Elementi {
				//		pdf.Cell(40, 10, latcyr.C(e.Element.Naziv, w.Cyr))
				//		pdf.Ln(8)
				//		pdf.Cell(40, 10, latcyr.C("Jedinična cena", w.Cyr))
				//		pdf.Cell(40, 10, fmt.Sprintf("%.2f", e.Element.Cena))
				//		pdf.Cell(40, 10, latcyr.C("Jedinična cena", w.Cyr))
				//		pdf.Ln(8)
				//		pdf.Cell(40, 10, latcyr.C("Količina", w.Cyr))
				//		pdf.Cell(40, 10, fmt.Sprint(e.Kolicina))
				//		pdf.Cell(40, 10, e.Element.Jedinica)
				//		pdf.Cell(40, 10, latcyr.C("Cena suma:", w.Cyr))
				//		pdf.Cell(40, 10, fmt.Sprintf("%.2f", e.SumaCena))
				//		pdf.Ln(8)
				//	}
				//fileStr := example.Filename("Fpdf_AddPage")
				//err := pdf.OutputFileAndClose(fileStr)
				//example.Summary(err, fileStr)

				pdf.SetFont("Times", "B", 16)
				pdf.CellFormat(0, 10, latcyr.C("Specifikacija materijala", w.Cyr), "0", 0, "", false, 0, "")
				pdf.Ln(20)

				pdf.SetFont("Arial", "", 10)
				for _, e := range w.Suma.UkupanNeophodanMaterijalPrikaz {
					cols := []float64{40, pagew - mleft - mright - 20}
					//rows := [][]string{}

					rows := [][]string{
						[]string{
							"Šifra", fmt.Sprint(e.Id),
						},
						[]string{
							"Naziv", e.Materijal.Naziv,
						},
						[]string{
							"Osobine i namena", e.Materijal.OsobineNamena,
						},
						[]string{
							"Jedinica mere", e.Materijal.JedinicaPotrosnje,
						},
						[]string{
							"Jedinična cena", fmt.Sprint(e.Materijal.Cena),
						},
						[]string{
							"Količina", fmt.Sprint(e.Kolicina),
						},
						[]string{
							"Vrednost rada", fmt.Sprintf("%.2f", e.UkupnaCena),
						},
					}
					for _, row := range rows {
						curx, y := pdf.GetXY()
						x := curx
						height := 0.
						_, lineHt := pdf.GetFontSize()
						for i, txt := range row {
							lines := pdf.SplitLines([]byte(txt), cols[i])
							h := float64(len(lines))*lineHt + marginCell*float64(len(lines))
							if h > height {
								height = h
							}
						}
						// add a new page if the height of the row doesn't fit on the page
						if pdf.GetY()+height > pageh-mbottom {
							pdf.AddPage()
							y = pdf.GetY()
						}
						for i, txt := range row {
							width := cols[i]
							//pdf.Rect(x, y, width, height, "")
							pdf.MultiCell(width, lineHt+marginCell, txt, "", "", false)
							x += width
							pdf.SetXY(x, y)
						}
						pdf.SetXY(curx, y+height)
					}
					pdf.Ln(8)
				}

				pdf.SetFont("Times", "B", 16)
				pdf.CellFormat(0, 10, latcyr.C("Suma materijal: ", w.Cyr)+fmt.Sprintf("%.2f", w.Suma.SumaCenaMaterijal), "0", 0, "", false, 0, "")
				pdf.Ln(20)

				///////////////////////////////////

				err := pdf.OutputFileAndClose("nalog.pdf")
				if err != nil {
				}
				open.Run("nalog.pdf")

			}
			btn.Layout(w.Context, stampajDugme)
		}
	}
}

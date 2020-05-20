package calc

import (
	"fmt"
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"github.com/gioapp/gel"
)

var (
	stampajDugme = new(gel.Button)
)

func (w *WingCal) SumaIzgled() func() {
	return func() {
		w.Tema.DuoUIcontainer(0, w.Tema.Colors["LightGrayI"]).Layout(w.Context, layout.NW, func() {
			var sumaSumarum float64
			layout.Flex{Axis: layout.Vertical}.Layout(w.Context,
				layout.Flexed(0.5, func() {
					layout.Flex{Axis: layout.Vertical}.Layout(w.Context,
						layout.Rigid(func() {
							ukupan := w.Tema.DuoUIcontainer(16, w.Tema.Colors["Primary"])
							ukupan.FullWidth = true
							ukupan.Layout(w.Context, layout.W, func() {
								suma := w.Tema.H5("Ukupna cena radova")
								suma.Alignment = text.End
								suma.Layout(w.Context)
							})
						}),
						layout.Flexed(1, func() {
							sumList.Layout(w.Context, len(w.Suma.Elementi), func(i int) {
								element := w.Suma.Elementi[i]
								sumaSumarum = sumaSumarum + element.SumaCena
								layout.UniformInset(unit.Dp(4)).Layout(w.Context, func() {
									layout.Flex{Axis: layout.Vertical}.Layout(w.Context,
										layout.Rigid(func() {
											layout.Flex{Axis: layout.Horizontal, Spacing: layout.SpaceBetween}.Layout(w.Context,
												layout.Flexed(0.8, func() {
													w.Tema.Body1(element.Element.Naziv).Layout(w.Context)
												}),
												layout.Flexed(0.1, func() {
													w.Tema.H6(fmt.Sprint(element.Kolicina)).Layout(w.Context)
												}),
												layout.Flexed(0.1, func() {
													w.Tema.H5(fmt.Sprint(element.SumaCena)).Layout(w.Context)
												}))
										}),
									)
								})
							})
						}),
						layout.Rigid(func() {
							suma := w.Tema.H5("Suma: " + fmt.Sprint(sumaSumarum))
							suma.Alignment = text.End
							suma.Layout(w.Context)
						}),
					)
				}),

				layout.Flexed(0.5, func() {
					layout.Flex{Axis: layout.Vertical}.Layout(w.Context,
						layout.Rigid(func() {
							ukupan := w.Tema.DuoUIcontainer(16, w.Tema.Colors["Primary"])
							ukupan.FullWidth = true
							ukupan.Layout(w.Context, layout.W, func() {
								suma := w.Tema.H5("Ukupan neophdni materijal")
								suma.Alignment = text.End
								suma.Layout(w.Context)
							})
						}),
						layout.Flexed(1, w.NeophodanMaterijal(ukupanNeophodanMaterijalList, w.Suma.UkupanNeophodanMaterijal)),

						layout.Rigid(w.Stampa()))
				}),
			)
		})
	}
}

func (w *WingCal) Stampa() func() {
	return func() {

		btn := w.Tema.Button("Stampaj")
		for stampajDugme.Clicked(w.Context) {
			//	pdf := gofpdf.New("P", "mm", "A4", "")
			//	pdf.AddPage()
			//	pdf.SetFont("Arial", "B", 12)
			//	for _, e := range w.Suma.Elementi {
			//		pdf.Cell(40, 10, e.Element.Naziv)
			//		pdf.Ln(8)
			//		pdf.Cell(40, 10, "Jedinicna cena")
			//		pdf.Cell(40, 10, fmt.Sprint(e.Element.Cena))
			//		pdf.Cell(40, 10, "Jedinicna cena")
			//		pdf.Ln(8)
			//		pdf.Cell(40, 10, "Kolicina")
			//		pdf.Cell(40, 10, fmt.Sprint(e.Kolicina))
			//		pdf.Cell(40, 10, e.Element.Jedinica)
			//		pdf.Cell(40, 10, "Cena suma:")
			//		pdf.Cell(40, 10, fmt.Sprint(e.SumaCena))
			//		pdf.Ln(8)
			//	}
			//
			//	pdf.Ln(16)
			//	pdf.Cell(40, 10, "UkupanNeophodanMaterijal:")
			//
			//
			//	for _, e := range w.Suma.UkupanNeophodanMaterijal {
			//		//pdf.Cell(40, 10, e.Materijal.Naziv)
			//		pdf.Ln(8)
			//		pdf.Cell(40, 10, "Jedinicna cena")
			//		pdf.Cell(40, 10, fmt.Sprint(e.Materijal.Cena))
			//		pdf.Cell(40, 10, e.Materijal.Jedinica)
			//
			//		pdf.Ln(8)
			//		pdf.Cell(40, 10, "Ukupno")
			//		pdf.Cell(40, 10, fmt.Sprint(e.UkupnoPakovanja))
			//		pdf.Cell(40, 10, "Ukupna cena")
			//		pdf.Cell(40, 10, fmt.Sprint(e.UkupnaCena))
			//
			//	}
			//
			//
			//
			//	pdf.Ln(16)
			//	pdf.Cell(40, 10, "Suma:")
			//	pdf.Cell(40, 10, fmt.Sprint(w.Suma.SumaCena))
			//
			//	err := pdf.OutputFileAndClose("nalog.pdf")
			//	if err != nil {
			//	}
			fmt.Println("Sume", w.Suma.UkupanNeophodanMaterijal)
		}
		btn.Layout(w.Context, stampajDugme)
	}
}

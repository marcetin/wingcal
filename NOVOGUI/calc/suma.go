package calc

import (
	"fmt"
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"github.com/marcetin/wingcal/model"
	"github.com/marcetin/wingcal/pkg/gelook"
	"github.com/marcetin/wingcal/pkg/latcyr"
)

var (
	tabelaSuma = map[int]int{}
)

func (w *WingCal) SumaIzgled() func() {
	return func() {
		w.Tema.DuoUIcontainer(0, w.Tema.Colors["LightGrayI"]).Layout(w.Context, layout.NW, func() {

			layout.Flex{Axis: layout.Vertical}.Layout(w.Context,
				layout.Flexed(0.5, func() {
					layout.Flex{Axis: layout.Vertical}.Layout(w.Context,
						layout.Rigid(func() {
							ukupan := w.Tema.DuoUIcontainer(16, w.Tema.Colors["Primary"])
							ukupan.FullWidth = true
							ukupan.Layout(w.Context, layout.W, func() {
								suma := w.Tema.H5(latcyr.C("Ukupna cena radova", w.Cyr))
								suma.Alignment = text.End
								suma.Layout(w.Context)
							})
						}),
						layout.Flexed(1, func() {
							layout.Flex{Axis: layout.Vertical}.Layout(w.Context,
								layout.Rigid(func() {
									layout.UniformInset(unit.Dp(4)).Layout(w.Context, func() {
										layout.Flex{Axis: layout.Horizontal, Spacing: layout.SpaceBetween}.Layout(w.Context,
											layout.Flexed(1, func() {
												w.Tema.Caption("Naziv").Layout(w.Context)
											}),
											layout.Rigid(
												func() {
													w.cell("Pojedinačna cena")
													//w.Tema.Caption("Pojedinačna cena").Layout(w.Context)
													w.tabela(0, w.Context.Dimensions.Size.X)
												}),
											layout.Rigid(func() {
												w.cell("Količina")
												//w.Tema.Caption("Količina").Layout(w.Context)
												w.tabela(1, w.Context.Dimensions.Size.X)
											}),
											layout.Rigid(func() {
												w.cell("Cena")
												//w.Tema.Caption("Cena").Layout(w.Context)
												w.tabela(2, w.Context.Dimensions.Size.X)
											}),
											layout.Rigid(func() {
												w.cell("briši")
												//w.Tema.Caption("briši").Layout(w.Context)
												w.tabela(3, w.Context.Dimensions.Size.X)
											}))
									})
								}),
								layout.Rigid(func() {
									sumList.Layout(w.Context, len(w.Suma.Elementi), func(i int) {
										element := w.Suma.Elementi[i]
										layout.UniformInset(unit.Dp(4)).Layout(w.Context, func() {
											layout.Flex{Axis: layout.Horizontal, Spacing: layout.SpaceBetween}.Layout(w.Context,
												layout.Flexed(1, func() {
													w.Tema.Body1(element.Element.Naziv).Layout(w.Context)
												}),
												layout.Rigid(func() {
													w.cell(fmt.Sprint(element.Element.Cena))

													//w.Tema.H6(fmt.Sprint(element.Element.Cena)).Layout(w.Context)
													w.tabela(0, w.Context.Dimensions.Size.X)

												}),
												layout.Rigid(func() {
													w.cell(fmt.Sprint(element.Kolicina))

													//w.Tema.H6(fmt.Sprint(element.Kolicina)).Layout(w.Context)
													w.tabela(1, w.Context.Dimensions.Size.X)

												}),
												layout.Rigid(func() {
													w.cell(fmt.Sprintf("%.2f", element.SumaCena))
													//w.Tema.Body1(fmt.Sprintf("%.2f", element.SumaCena)).Layout(w.Context)
													w.tabela(2, w.Context.Dimensions.Size.X)

												}),
												layout.Rigid(func() {

													btn := w.Tema.Button(latcyr.C("BRIŠI", w.Cyr) + fmt.Sprint(i))
													btn.Inset = layout.Inset{unit.Dp(5), unit.Dp(3), unit.Dp(5), unit.Dp(5)}
													btn.TextSize = unit.Dp(12)
													btn.Color = gelook.HexARGB(w.Tema.Colors["Gray"])
													btn.Background = gelook.HexARGB(w.Tema.Colors["yellow"])
													for element.DugmeBrisanje.Clicked(w.Context) {
														fmt.Println("iii", i)

														//fmt.Println("w.Suma.ElementiPREEEE", w.Suma.Elementi)
														w.Suma.Elementi = append(w.Suma.Elementi[:i], w.Suma.Elementi[i+1:]...)
														//w.Suma.Elementi[fmt.Sprint(i)] =  model.WingIzabraniElement{}
														tabelaSuma = map[int]int{}
														w.NeopodanMaterijal()
														w.SumaRacunica()
													}
													btn.Layout(w.Context, element.DugmeBrisanje)
													w.tabela(3, w.Context.Dimensions.Size.X)
												}))
										})
									})
								}))
						}),
						layout.Rigid(func() {
							suma := w.Tema.H5(latcyr.C("Suma: ", w.Cyr) + fmt.Sprint(w.Suma.SumaCena))
							suma.Alignment = text.End
							suma.Layout(w.Context)
						}),
					)
				}),

				layout.Flexed(0.5, func() {
					width := w.Context.Constraints.Width.Max
					layout.Flex{Axis: layout.Vertical}.Layout(w.Context,
						layout.Rigid(func() {
							ukupan := w.Tema.DuoUIcontainer(16, w.Tema.Colors["Primary"])
							ukupan.FullWidth = true
							ukupan.Layout(w.Context, layout.W, func() {
								suma := w.Tema.H5(latcyr.C("Ukupan neophdni materijal", w.Cyr))
								suma.Alignment = text.End
								suma.Layout(w.Context)
							})
						}),
						layout.Rigid(func() {
							w.Context.Constraints.Width.Min = width
							layout.Flex{
								Axis:    layout.Horizontal,
								Spacing: layout.SpaceBetween,
							}.Layout(w.Context,
								layout.Rigid(func() {
									w.Tema.H6(latcyr.C("Količina:", w.Cyr)).Layout(w.Context)
								}),
								layout.Rigid(func() {
									w.Tema.H6(latcyr.C("Ukupna cena:", w.Cyr)).Layout(w.Context)
								}))
						}),
						layout.Rigid(w.Tema.DuoUIline(w.Context, 0, 0, 2, w.Tema.Colors["Gray"])),
						layout.Flexed(1, w.UkupanNeophodanMaterijal(ukupanNeophodanMaterijalList)),

						layout.Rigid(w.Stampa()))
				}),
			)
		})
	}
}

func (w *WingCal) NeopodanMaterijal() {
	u := make(map[int]model.WingNeophodanMaterijal)
	for _, e := range w.Suma.Elementi {
		for _, n := range e.Element.NeophodanMaterijal {
			u[n.Id] = model.WingNeophodanMaterijal{
				Id:       n.Id,
				Kolicina: w.Suma.UkupanNeophodanMaterijal[n.Id].Kolicina + float64(kolicina.Value),
			}
		}
	}
	w.Suma.UkupanNeophodanMaterijal = u
}

func (w *WingCal) tabela(colona, width int) {
	if width > tabelaSuma[colona] {
		tabelaSuma[colona] = width
	}
	w.Context.Dimensions.Size.X = tabelaSuma[colona]
}

func (w *WingCal) cell(tekst string) {
	layout.Inset{
		Top:    unit.Dp(0),
		Right:  unit.Dp(4),
		Bottom: unit.Dp(0),
		Left:   unit.Dp(4),
	}.Layout(w.Context, func() {
		w.Tema.Caption(tekst).Layout(w.Context)
	})
}

func (w *WingCal) SumaRacunica() {
	s := 0.0
	materijal := make(map[int]model.WingNeophodanMaterijal)
	for _, e := range w.Suma.Elementi {
		s = s + e.SumaCena

		for _, n := range e.Element.NeophodanMaterijal {
			//materijal = w.Suma.UkupanNeophodanMaterijal[i-1]
			nn := model.WingNeophodanMaterijal{
				Id:        n.Materijal.Id - 1,
				Materijal: *w.Materijal[n.Id-1],
			}
			nn.Kolicina = w.Suma.UkupanNeophodanMaterijal[n.Materijal.Id-1].Kolicina + float64(e.Kolicina)
			//kolicina := 0.0
			//if n.Koeficijent > 0 {
			kolicina := nn.Materijal.Potrosnja * float64(e.Kolicina) * n.Koeficijent
			//}
			n.UkupnaCena = nn.Materijal.Cena * float64(kolicina)
			n.UkupnoPakovanja = int(kolicina / float64(n.Materijal.Pakovanje))

			//w.Suma.UkupanNeophodanMaterijal[n.Id].Kolicina = w.Suma.UkupanNeophodanMaterijal[n.Id].Kolicina + float64(kolicina)

			materijal[n.Materijal.Id] = nn

			fmt.Println("kkkn.Materijal.Id:", n.Materijal.Id)
			fmt.Println("kkkn.Id-1:", n.Id)

			fmt.Println("NNNEmaterijalmaterijal:", n.Materijal.Naziv)
			fmt.Println("Potrosnja:", n.Materijal.Potrosnja)
			fmt.Println("kkk:", kolicina)
			fmt.Println("e:", e.Kolicina)
			fmt.Println("UkupnaCena:", n.UkupnaCena)
			fmt.Println("Kolicina:", n.Kolicina)
			fmt.Println("Koeficijent:", n.Koeficijent)

		}
	}
	w.Suma.SumaCena = s
	w.Suma.UkupanNeophodanMaterijal = materijal
}

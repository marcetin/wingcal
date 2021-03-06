package calc

import (
	"fmt"
	"gioui.org/layout"
	"github.com/marcetin/wingcal/model"
	"github.com/marcetin/wingcal/pkg/gelook"
)

func (w *WingCal) PrikazaniElementIzgled() func() {
	return func() {
		layout.Flex{
			Axis: layout.Vertical,
		}.Layout(w.Context,
			layout.Rigid(func() {
				w.Tema.DuoUIcontainer(8, w.Tema.Colors["LightGray"]).Layout(w.Context, layout.W, func() {
					w.Tema.H5(w.PrikazaniElement.Naziv).Layout(w.Context)
				})
			}),
			layout.Flexed(1, func() {
				w.Tema.DuoUIcontainer(8, w.Tema.Colors["LightGray"]).Layout(w.Context, layout.NW, func() {
					//sumaCena := 111.33
					width := w.Context.Constraints.Width.Max
					layout.Flex{Axis: layout.Vertical}.Layout(w.Context,
						layout.Rigid(func() {
							w.Tema.Body1(w.PrikazaniElement.Opis).Layout(w.Context)
						}),
						layout.Rigid(func() {
							w.Tema.Caption(w.PrikazaniElement.Obracun).Layout(w.Context)
						}),
						layout.Rigid(w.Tema.DuoUIline(w.Context, 0, 0, 32, w.Tema.Colors["Gray"])),

						layout.Rigid(func() {
							w.Tema.H6("Neophodan materijal za izvrsenje radova").Layout(w.Context)
						}),
						layout.Rigid(func() {
							w.Context.Constraints.Width.Min = width
							layout.Flex{
								Axis:    layout.Horizontal,
								Spacing: layout.SpaceBetween,
							}.Layout(w.Context,
								layout.Flexed(0.2, func() {
									w.Tema.Caption("Naziv").Layout(w.Context)
								}),
								layout.Flexed(0.2, func() {
									w.Tema.Caption("Potrosnja").Layout(w.Context)
								}),
								layout.Flexed(0.15, func() {
									w.Tema.Caption("Koeficijent:").Layout(w.Context)
								}),
								layout.Flexed(0.15, func() {
									w.Tema.Caption("Merodavna potrosnja").Layout(w.Context)
								}),
								layout.Flexed(0.15, func() {
									w.Tema.Caption("Pakovanja").Layout(w.Context)
								}),
								layout.Flexed(0.15, func() {
									w.Tema.Caption("Cena materijala").Layout(w.Context)
								}))
						}),
						layout.Rigid(w.NeophodanMaterijal(neophodanMaterijalList, w.PrikazaniElement.NeophodanMaterijal, false)),
					)
				})
			}),
			layout.Rigid(func() {
				w.Tema.DuoUIcontainer(0, w.Tema.Colors["Gray"]).Layout(w.Context, layout.NW, func() {
					sumaCena := float64(kolicina.Value) * w.PrikazaniElement.Cena

					layout.Flex{
						Axis:    layout.Horizontal,
						Spacing: layout.SpaceBetween,
					}.Layout(w.Context,
						layout.Flexed(1, func() {
							layout.Flex{
								Axis: layout.Vertical,
							}.Layout(w.Context,
								layout.Rigid(func() {
									w.Tema.DuoUIcontainer(8, w.Tema.Colors["Primary"]).Layout(w.Context, layout.NW, func() {
										w.Tema.H6("Cena:" + fmt.Sprint(w.PrikazaniElement.Cena)).Layout(w.Context)
									})
								}),
								layout.Rigid(w.Tema.DuoUIline(w.Context, 0, 0, 0, w.Tema.Colors["Gray"])),
								layout.Rigid(func() {
									w.Tema.DuoUIcontainer(8, w.Tema.Colors["Primary"]).Layout(w.Context, layout.NW, func() {
										w.Tema.H6("Suma:" + fmt.Sprintf("%.2f", sumaCena)).Layout(w.Context)
									})
								}),
							)
						}),
						layout.Rigid(func() {
							layout.Flex{
								Axis: layout.Vertical,
							}.Layout(w.Context,
								layout.Rigid(func() {
									w.Tema.DuoUIcounter(func() {}).Layout(w.Context, kolicina, "KOLICINA", fmt.Sprint(kolicina.Value))
								}),
								layout.Rigid(func() {
									btn := w.Tema.Button("DODAJ")
									//btn.FullWidth = true
									//btn.FullHeight = true
									btn.Background = gelook.HexARGB(w.Tema.Colors["Secondary"])
									suma := model.WingIzabraniElement{
										Kolicina: kolicina.Value,
										SumaCena: sumaCena,
										Element:  *w.PrikazaniElement,
									}
									for dodajDugme.Clicked(w.Context) {
										w.Suma.Elementi = append(w.Suma.Elementi, suma)
										for _, n := range w.PrikazaniElement.NeophodanMaterijal {
											w.Suma.UkupanNeophodanMaterijal[n.Id] = model.WingNeophodanMaterijal{
												Id:       n.Id,
												Kolicina: w.Suma.UkupanNeophodanMaterijal[n.Id].Kolicina + float64(kolicina.Value),
											}

										}
										//var neophodanmaterijal map[int]model.WingNeophodanMaterijal
									}
									btn.Layout(w.Context, dodajDugme)
								}),
							)
						}))
				})
			}))
	}
}

package calc

import (
	"fmt"
	"gioui.org/layout"
	"github.com/gioapp/gel"
	"github.com/marcetin/wingcal/model"
	"github.com/marcetin/wingcal/pkg/gelook"
	"github.com/marcetin/wingcal/pkg/latcyr"
)

var (
	zatvoriElementDugme = new(gel.Button)
)

func (w *WingCal) PrikazaniElementIzgled() func() {
	return func() {
		layout.Flex{
			Axis: layout.Vertical,
		}.Layout(w.Context,
			layout.Rigid(func() {
				btnZatvoriElement := w.Tema.Button(latcyr.C("zatvori", w.Cyr))
				btnZatvoriElement.Background = gelook.HexARGB(w.Tema.Colors["Secondary"])
				for zatvoriElementDugme.Clicked(w.Context) {
					w.Element = false
				}
				btnZatvoriElement.Layout(w.Context, zatvoriElementDugme)
			}),
			layout.Rigid(func() {
				w.Tema.DuoUIcontainer(8, w.Tema.Colors["LightGray"]).Layout(w.Context, layout.W, func() {
					w.Tema.H5(fmt.Sprint(w.Podvrsta) + "." + fmt.Sprint(w.Roditelj) + "." + fmt.Sprint(w.PrikazaniElement.Id) + " " + latcyr.C(w.PrikazaniElement.Naziv, w.Cyr)).Layout(w.Context)
				})
			}),
			layout.Flexed(1, w.PrikazaniElementOpis()),
			layout.Rigid(w.PrikazaniElementSuma()))
	}
}

func (w *WingCal) PrikazaniElementStavkeMaterijala(width int) func() {
	return func() {
		w.Context.Constraints.Width.Min = width
		layout.Flex{
			Axis:    layout.Horizontal,
			Spacing: layout.SpaceBetween,
		}.Layout(w.Context,
			layout.Flexed(0.4, func() {
				w.Tema.Caption(latcyr.C("Naziv", w.Cyr)).Layout(w.Context)
			}),
			layout.Flexed(0.15, func() {
				w.Tema.Caption(latcyr.C("Potrosnja", w.Cyr)).Layout(w.Context)
			}),
			layout.Flexed(0.15, func() {
				w.Tema.Caption(latcyr.C("Koeficijent", w.Cyr)).Layout(w.Context)
			}),
			layout.Flexed(0.15, func() {
				w.Tema.Caption(latcyr.C("Merodavna potrosnja", w.Cyr)).Layout(w.Context)
			}),
			layout.Flexed(0.15, func() {
				w.Tema.Caption(latcyr.C("Cena materijala", w.Cyr)).Layout(w.Context)
			}))
	}
}

func (w *WingCal) PrikazaniElementDugmeDodaj(sumaCena float64) func() {
	return func() {
		btn := w.Tema.Button(latcyr.C("DODAJ", w.Cyr))
		//btn.FullWidth = true
		//btn.FullHeight = true
		btn.Background = gelook.HexARGB(w.Tema.Colors["Secondary"])
		var varijacijaRada int

		for dodajDugme.Clicked(w.Context) {
			if kolicina.Value > 0 {
				for _, s := range w.Suma.Elementi {
					if s.Element.Id == w.PrikazaniElement.Id {
						varijacijaRada = varijacijaRada + 1
						fmt.Println("varijacijaRada:", varijacijaRada)
					}
					fmt.Println("elem:", s.Element.Id)
				}
				fmt.Println("varijacijaRadaIIIIIIIIIIIII:", varijacijaRada)
				suma := model.WingIzabraniElement{
					Sifra:         fmt.Sprint(w.Podvrsta) + "." + fmt.Sprint(w.Roditelj) + "." + fmt.Sprint(w.PrikazaniElement.Id) + "." + fmt.Sprint(varijacijaRada+1),
					Kolicina:      kolicina.Value,
					SumaCena:      sumaCena,
					Element:       *w.PrikazaniElement,
					DugmeBrisanje: new(gel.Button),
				}
				w.Suma.Elementi = append(w.Suma.Elementi, &suma)
				//w.Suma.Elementi[len(w.Suma.Elementi)] = suma
				//for _, n := range w.PrikazaniElement.NeophodanMaterijal {
				//	w.Suma.UkupanNeophodanMaterijal[n.Id] = model.WingNeophodanMaterijal{
				//		Id:       n.Id,
				//		Kolicina: w.Suma.UkupanNeophodanMaterijal[n.Id].Kolicina + float64(kolicina.Value),
				//	}
				//}
			}
			//var neophodanmaterijal map[int]model.WingNeophodanMaterijal
			//w.NeopodanMaterijal()
			w.SumaRacunica()
		}
		btn.Layout(w.Context, dodajDugme)
	}
}

func (w *WingCal) PrikazaniElementOpis() func() {
	return func() {
		w.Tema.DuoUIcontainer(8, w.Tema.Colors["LightGray"]).Layout(w.Context, layout.NW, func() {
			//sumaCena := 111.33
			width := w.Context.Constraints.Width.Max
			layout.Flex{Axis: layout.Vertical}.Layout(w.Context,
				layout.Rigid(func() {
					w.Tema.Body1(latcyr.C(w.PrikazaniElement.Opis, w.Cyr)).Layout(w.Context)
				}),
				layout.Rigid(func() {
					w.Tema.Caption(latcyr.C(w.PrikazaniElement.Obracun, w.Cyr)).Layout(w.Context)
				}),
				layout.Rigid(w.Tema.DuoUIline(w.Context, 0, 0, 32, w.Tema.Colors["Gray"])),

				layout.Rigid(func() {
					w.Tema.H6(latcyr.C("Neophodan materijal za izvrsenje radova", w.Cyr)).Layout(w.Context)
				}),
				layout.Rigid(w.PrikazaniElementStavkeMaterijala(width)),
				layout.Rigid(w.RadNeophodanMaterijal(neophodanMaterijalList)),
			)
		})
	}
}

func (w *WingCal) PrikazaniElementSuma() func() {
	return func() {
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
								w.Tema.H6(latcyr.C("Cena:", w.Cyr) + fmt.Sprint(w.PrikazaniElement.Cena)).Layout(w.Context)
							})
						}),
						layout.Rigid(w.Tema.DuoUIline(w.Context, 0, 0, 0, w.Tema.Colors["Gray"])),
						layout.Rigid(func() {
							w.Tema.DuoUIcontainer(8, w.Tema.Colors["Primary"]).Layout(w.Context, layout.NW, func() {
								w.Tema.H6(latcyr.C("Suma:", w.Cyr) + fmt.Sprintf("%.2f", sumaCena)).Layout(w.Context)
							})
						}),
					)
				}),
				layout.Rigid(func() {
					layout.Flex{
						Axis: layout.Vertical,
					}.Layout(w.Context,
						layout.Rigid(func() {
							w.Tema.DuoUIcounter(func() {}).Layout(w.Context, kolicina, latcyr.C("KOLIČINA", w.Cyr), fmt.Sprint(kolicina.Value))
						}),
						layout.Rigid(w.PrikazaniElementDugmeDodaj(sumaCena)),
					)
				}))
		})
	}
}

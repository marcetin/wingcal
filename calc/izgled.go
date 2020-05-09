package calc

import (
	"fmt"
	"gioui.org/layout"
	"gioui.org/unit"
	"github.com/gioapp/gel"
	"github.com/gioapp/gelook"
	"github.com/marcetin/wingcal/model"
)

var (
	IzborVrsteRadovaPanelElement = gel.NewPanel()
)

func (w *WingCal) IzborVrsteRadova() func() {
	return func() {
		//if w.Context.Constraints.Width.Max > 300 {
		//w.Context.Constraints.Width.Min = 300
		//w.Context.Constraints.Width.Max = 300
		//}
		w.Tema.DuoUIcontainer(0, w.Tema.Colors["DarkGrayI"]).Layout(w.Context, layout.W, func() {
			layout.Flex{Axis: layout.Vertical}.Layout(w.Context,
				layout.Rigid(w.NazivRoditelja()),
				layout.Rigid(w.Nazad()),
				layout.Flexed(1, func() {

					IzborVrsteRadovaPanelElement.PanelObject = w.IzbornikRadova.PodvrsteRadova
					IzborVrsteRadovaPanelElement.PanelObjectsNumber = len(w.IzbornikRadova.PodvrsteRadova)
					izborVrsteRadovaPanel := w.Tema.DuoUIpanel()
					izborVrsteRadovaPanel.ScrollBar = w.Tema.ScrollBar(0)
					izborVrsteRadovaPanel.Layout(w.Context, IzborVrsteRadovaPanelElement, func(i int, in interface{}) {
						//if in != nil {
						//addresses := in.([]model.DuoUIaddress)

						//*w.IzbornikRadova.Izbor.PodvrsteRadova[i].Roditelj =w.IzbornikRadova.Izbor
						vrstarada := w.IzbornikRadova.PodvrsteRadova[i]
						//if vrstarada.Element && vrstarada.NeophodanMaterijal != nil  {
						//vrstarada.Roditelj = &w.IzbornikRadova
						layout.UniformInset(unit.Dp(0)).Layout(w.Context, func() {
							layout.Flex{Axis: layout.Vertical}.Layout(w.Context,
								layout.Rigid(func() {

									btn := w.Tema.Button(vrstarada.Naziv)
									btn.CornerRadius = unit.Dp(0)
									//btn.FullWidth = true
									btn.Background = gelook.HexARGB(w.Tema.Colors["Gray"])
									//btn.CornerRadius = 0
									//btn.Align = layout.W
									for w.LinkoviIzboraVrsteRadova[i].Clicked(w.Context) {
										w.Putanja = append(w.Putanja, &vrstarada)

										if vrstarada.Baza {
											elementi := w.Db.DbRead(w.IzbornikRadova.Slug, vrstarada.Slug)
											vrstarada.PodvrsteRadova = elementi.PodvrsteRadova

										}
										//if vrstarada.PodvrsteRadova != nil {
										w.Podvrsta(&vrstarada)

										w.EditPolja.Id.SetText(fmt.Sprint(w.PrikazaniElement.Id))
										w.EditPolja.Naziv.SetText(w.PrikazaniElement.Naziv)
										w.EditPolja.Opis.SetText(w.PrikazaniElement.Opis)
										w.EditPolja.Obracun.SetText(w.PrikazaniElement.Obracun)
										w.EditPolja.Jedinica.SetText(fmt.Sprint(w.PrikazaniElement.Jedinica))
										w.EditPolja.Cena.SetText(fmt.Sprint(w.PrikazaniElement.Cena))

										for _, neophodanMaterijal := range w.PrikazaniElement.NeophodanMaterijal {
											w.EditPolja.Materijal[neophodanMaterijal.Id].SetText("0")
											w.EditPolja.Materijal[neophodanMaterijal.Id].SetText(fmt.Sprint(neophodanMaterijal.Kolicina))
										}
										//}
									}
									btn.Layout(w.Context, w.LinkoviIzboraVrsteRadova[i])
								}),
								layout.Rigid(w.Tema.DuoUIline(w.Context, 0, 0, 0, "Dark")),
							)
						})
						//}
					})
				}))
		})
	}
}

func (w *WingCal) Nazad() func() {
	return func() {
		if len(w.Putanja) > 0 {
			btnNazad := w.Tema.Button("NAZAD")
			btnNazad.Background = gelook.HexARGB(w.Tema.Colors["Secondary"])
			for nazadDugme.Clicked(w.Context) {
				w.IzbornikRadova = w.Putanja[len(w.Putanja)-1]
				w.GenerisanjeLinkova(w.Putanja[len(w.Putanja)-1].PodvrsteRadova)
				w.Putanja = w.Putanja[:len(w.Putanja)-1]
				w.Roditelj()
				fmt.Println("IzbornikroditeL::" + w.IzbornikRadova.Slug)
				//fmt.Println("roditeL::" + w.IzbornikRadova.Roditelj.Slug)
			}
			btnNazad.Layout(w.Context, nazadDugme)
		}
	}
}

func (w *WingCal) NeophodanMaterijal(l *layout.List, n map[int]model.WingNeophodanMaterijal) func() {
	return func() {
		width := w.Context.Constraints.Width.Max
		l.Layout(w.Context, len(n), func(i int) {
			materijal := n[i]
			materijal.Materijal = w.Materijal[materijal.Id]
			materijal.UkupnaCena = materijal.Materijal.Cena * float64(materijal.Kolicina)
			materijal.UkupnoPakovanja = materijal.Kolicina / materijal.Materijal.Pakovanje
			w.Context.Constraints.Width.Min = width
			layout.Flex{
				Axis: layout.Vertical,
			}.Layout(w.Context,
				layout.Rigid(func() {
					layout.Flex{
						Axis:    layout.Horizontal,
						Spacing: layout.SpaceBetween,
					}.Layout(w.Context,
						layout.Rigid(func() {
							w.Tema.Body1(materijal.Materijal.Naziv).Layout(w.Context)
						}),
						layout.Rigid(func() {
							w.Tema.Caption(fmt.Sprint(materijal.Materijal.Obracun)).Layout(w.Context)
						}),
					)
				}),
				layout.Rigid(func() {
					layout.Flex{
						Axis:    layout.Horizontal,
						Spacing: layout.SpaceBetween,
					}.Layout(w.Context,
						layout.Rigid(func() {
							w.Tema.Body1(fmt.Sprint(materijal.Kolicina)).Layout(w.Context)
						}),
						layout.Rigid(func() {
							w.Tema.Body1(fmt.Sprint(materijal.UkupnoPakovanja)).Layout(w.Context)
						}),
						layout.Rigid(func() {
							w.Tema.Body1("Ukupna cena:" + fmt.Sprint(materijal.UkupnaCena)).Layout(w.Context)
						}),
					)
				}),
				layout.Rigid(w.Tema.DuoUIline(w.Context, 0, 0, 8, "Gray")),
			)
		})
	}
}

package calc

import (
	"fmt"
	"gioui.org/layout"
	"gioui.org/unit"
	"github.com/marcetin/wingcal/model"
)

func (w *WingCal) IzborVrsteRadova() func() {
	return func() {
		//if w.Context.Constraints.Width.Max > 300 {
		//w.Context.Constraints.Width.Min = 300
		//w.Context.Constraints.Width.Max = 300
		//}
		w.Tema.DuoUIitem(0, w.Tema.Colors["DarkGrayI"]).Layout(w.Context, layout.W, func() {
			layout.Flex{Axis: layout.Vertical}.Layout(w.Context,
				layout.Rigid(w.NazivRoditelja()),
				layout.Rigid(w.Nazad()),
				layout.Flexed(1, func() {
					izbornikRadova.Layout(w.Context, len(w.IzbornikRadova.PodvrsteRadova), func(i int) {
						//*w.IzbornikRadova.Izbor.PodvrsteRadova[i].Roditelj =w.IzbornikRadova.Izbor
						vrstarada := w.IzbornikRadova.PodvrsteRadova[i]
						//if vrstarada.Element && vrstarada.NeophodanMaterijal != nil  {
						//vrstarada.Roditelj = &w.IzbornikRadova
						layout.UniformInset(unit.Dp(0)).Layout(w.Context, func() {
							layout.Flex{Axis: layout.Vertical}.Layout(w.Context,
								layout.Rigid(func() {

									btn := w.Tema.Button(vrstarada.Naziv)
									btn.FullWidth = true
									btn.Background = w.Tema.Colors["Gray"]
									btn.CornerRadius = 0
									btn.Align = layout.W
									for w.LinkoviIzboraVrsteRadova[i].Clicked(w.Context) {

										if vrstarada.Baza {
											elementi := w.Db.DbRead(w.IzbornikRadova.Slug, vrstarada.Slug)
											vrstarada.PodvrsteRadova = elementi.PodvrsteRadova

										}
										//if vrstarada.PodvrsteRadova != nil {
										w.Podvrsta(&vrstarada)
										//}
									}
									btn.Layout(w.Context, w.LinkoviIzboraVrsteRadova[i])
								}),
								layout.Rigid(w.Tema.DuoUIline(w.Context, 0, "Dark")),
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
		if w.IzbornikRadova.Roditelj != nil {
			btnNazad := w.Tema.Button("NAZAD")
			btnNazad.Background = w.Tema.Colors["Secondary"]
			for nazadDugme.Clicked(w.Context) {
				w.GenerisanjeLinkova(w.IzbornikRadova.Roditelj.PodvrsteRadova)
				w.Roditelj()
				w.IzbornikRadova = w.IzbornikRadova.Roditelj
				fmt.Println("IzbornikroditeL::" + w.IzbornikRadova.Slug)
				fmt.Println("roditeL::" + w.IzbornikRadova.Roditelj.Slug)
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
				layout.Rigid(w.Tema.DuoUIline(w.Context, 8, "Gray")),
			)
		})
	}
}

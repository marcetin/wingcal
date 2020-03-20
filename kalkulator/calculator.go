package kalkulator

import (
	"fmt"
	"gioui.org/layout"
	"gioui.org/unit"
	"github.com/marcetin/wingcal/calc"
	"github.com/marcetin/wingcal/pkg/gel"
)

var (
	thingEditTitle   = new(gel.Editor)
	thingEditContent = new(gel.Editor)
	//post             = new(model.DuoCMSpost)
	stampajDugme    = new(gel.Button)
	kalkulatorDugme = new(gel.Button)
	materijalDugme  = new(gel.Button)
	headerMenuList  = &layout.List{
		Axis: layout.Horizontal,
	}
	materijalList = &layout.List{
		Axis: layout.Vertical,
	}
)

type DuoCMSadmin struct {
	Menu *DuoCMSmenu
}

type DuoCMSmenu struct {
	Title string
	Items map[string]DuoCMSmenuItem
}

type DuoCMSmenuItem struct {
	Title       string
	Description string
	Icon        string
	Link        func()
	subItems    map[string]DuoCMSmenuItem
}

var selected int

func glavniEkran(w *calc.WingCal) func() {
	return func() {
		layout.Flex{
			Axis: layout.Vertical,
		}.Layout(w.Context,
			layout.Rigid(func() {
				w.Tema.DuoUIitem(0, w.Tema.Colors["DarkGrayI"]).Layout(w.Context, layout.Center, header(w))
			}),
			layout.Flexed(1, strana(w)),
			layout.Rigid(func() {
				w.Tema.DuoUIitem(0, w.Tema.Colors["DarkGray"]).Layout(w.Context, layout.Center, func() {
					w.Tema.H3("Footer").Layout(w.Context)
				})
			}))
	}
}

func strana(w *calc.WingCal) (s func()) {
	switch w.Strana {
	case "materijal":
		s = material(w)
	case "kalkulator":
		s = calclulator(w)
	}
	return
}

func header(w *calc.WingCal) func() {
	return func() {
		headerMenu := []func(){
			func() {
				btnKalkulator := w.Tema.Button("KALKULATOR")
				btnKalkulator.Background = w.Tema.Colors["Secondary"]
				for kalkulatorDugme.Clicked(w.Context) {
					w.Strana = "kalkulator"
				}
				btnKalkulator.Layout(w.Context, kalkulatorDugme)
			},
			func() {
				btnMaterijal := w.Tema.Button("MATERIJAL")
				btnMaterijal.Background = w.Tema.Colors["Secondary"]
				for materijalDugme.Clicked(w.Context) {
					w.Strana = "materijal"
				}
				btnMaterijal.Layout(w.Context, materijalDugme)
			},
		}
		headerMenuList.Layout(w.Context, len(headerMenu), func(i int) {
			layout.UniformInset(unit.Dp(0)).Layout(w.Context, headerMenu[i])
		})
	}
}

func calclulator(w *calc.WingCal) func() {
	return func() {
		layout.Flex{
			Axis: layout.Horizontal,
		}.Layout(w.Context,
			layout.Flexed(0.3, w.IzborVrsteRadova()),
			layout.Flexed(0.3, w.PrikazaniElementIzgled()),
			layout.Flexed(0.4, w.SumaIzgled()),
		)
	}
}

func material(w *calc.WingCal) func() {
	return func() {
		materijalList.Layout(w.Context, len(w.Materijal), func(i int) {
			m := w.Materijal[i]

			layout.Flex{
				Axis: layout.Vertical,
			}.Layout(w.Context,
				layout.Rigid(func() {
					layout.Flex{
						Axis: layout.Horizontal,
					}.Layout(w.Context,
						layout.Flexed(0.02, func() {
							w.Tema.Caption(fmt.Sprint(m.Id)).Layout(w.Context)
						}),
						layout.Flexed(0.3, func() {
							w.Tema.H6(m.Naziv).Layout(w.Context)
						}),
						layout.Flexed(0.3, func() {
							w.Tema.Body1(m.Opis).Layout(w.Context)
						}),
						layout.Flexed(0.2, func() {
							w.Tema.Caption(m.Obracun).Layout(w.Context)
						}),
						layout.Flexed(0.04, func() {
							w.Tema.Body2(fmt.Sprint(m.Pakovanje)).Layout(w.Context)
						}),
						layout.Flexed(0.04, func() {
							w.Tema.Body1(m.Jedinica).Layout(w.Context)
						}),
						layout.Flexed(0.1, func() {
							w.Tema.H6(fmt.Sprint(m.Cena)).Layout(w.Context)
						}),
					)
				}),
				layout.Rigid(w.Tema.DuoUIline(w.Context, 4, "Dark")),
			)
		})
	}
}

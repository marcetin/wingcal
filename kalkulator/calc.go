package kalkulator

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"github.com/gioapp/gel"
	"github.com/gioapp/gelook"
	"github.com/marcetin/wingcal/calc"
)

var (
	thingEditTitle   = new(gel.Editor)
	thingEditContent = new(gel.Editor)
	//post             = new(model.DuoCMSpost)
	stampajDugme    = new(gel.Button)
	kalkulatorDugme = new(gel.Button)
	materijalDugme  = new(gel.Button)
	projekatDugme   = new(gel.Button)
	editorDugme     = new(gel.Button)
	headerMenuList  = &layout.List{
		Axis: layout.Horizontal,
	}
	materijalList = &layout.List{
		Axis: layout.Vertical,
	}
	putanjaList = &layout.List{
		Axis: layout.Horizontal,
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
				w.Tema.DuoUIcontainer(0, w.Tema.Colors["DarkGrayI"]).Layout(w.Context, layout.Center, header(w))
			}),
			layout.Flexed(1, strana(w)),
			layout.Rigid(func() {
				w.Tema.DuoUIcontainer(0, w.Tema.Colors["DarkGray"]).Layout(w.Context, layout.Center, func() {
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
		s = kalkulator(w)
	case "projekat":
		s = projekat(w)
	}
	return
}

func header(w *calc.WingCal) func() {
	return func() {
		headerMenu := []func(){
			func() {
				if w.PrikazaniElement.Element {
					btnEditor := w.Tema.Button("EDITOR")
					btnEditor.Background = gelook.HexARGB(w.Tema.Colors["Secondary"])
					for editorDugme.Clicked(w.Context) {
						if w.Edit {
							w.Edit = false
						} else {
							w.Edit = true
						}
					}
					btnEditor.Layout(w.Context, editorDugme)
				}
			},
			func() {
				btnKalkulator := w.Tema.Button("KALKULATOR")
				btnKalkulator.Background = gelook.HexARGB(w.Tema.Colors["Secondary"])
				for kalkulatorDugme.Clicked(w.Context) {
					w.Strana = "kalkulator"
				}
				btnKalkulator.Layout(w.Context, kalkulatorDugme)
			},
			func() {
				btnMaterijal := w.Tema.Button("MATERIJAL")
				btnMaterijal.Background = gelook.HexARGB(w.Tema.Colors["Secondary"])
				for materijalDugme.Clicked(w.Context) {
					w.Strana = "materijal"
				}
				btnMaterijal.Layout(w.Context, materijalDugme)
			},
			func() {
				btnMaterijal := w.Tema.Button("PROJEKAT")
				btnMaterijal.Background = gelook.HexARGB(w.Tema.Colors["Secondary"])
				for projekatDugme.Clicked(w.Context) {
					w.Strana = "projekat"
				}
				btnMaterijal.Layout(w.Context, projekatDugme)
			},
			func() {
				putanjaList.Layout(w.Context, len(w.Putanja), func(i int) {
					w.Tema.Caption(w.Putanja[i].Naziv).Layout(w.Context)
				})

			},
		}
		headerMenuList.Layout(w.Context, len(headerMenu), func(i int) {
			layout.UniformInset(unit.Dp(0)).Layout(w.Context, headerMenu[i])
		})
	}
}

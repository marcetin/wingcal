package main

import (
	"gioui.org/layout"
	"github.com/gioapp/gel"
	"github.com/marcetin/wingcal/NOVOGUI/calc"
)

var (
	thingEditTitle   = new(gel.Editor)
	thingEditContent = new(gel.Editor)
	//post             = new(model.DuoCMSpost)
	stampajDugme  = new(gel.Button)
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

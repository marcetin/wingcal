package kalkulator

import (
	"gioui.org/layout"
	"github.com/marcetin/wingcal/calc"
)

func kalkulator(w *calc.WingCal) func() {
	return func() {
		layout.Flex{
			Axis: layout.Horizontal,
		}.Layout(w.Context,
			layout.Flexed(0.3, w.IzborVrsteRadova()),
			layout.Flexed(0.7, glavniDeo(w)),
		)
	}
}
func glavniDeo(w *calc.WingCal) func() {
	return func() {
		if w.Edit {
			layout.Flex{
				Axis: layout.Horizontal,
			}.Layout(w.Context,
				layout.Flexed(1, w.EditorElementaIzgled()),
			)
		} else {
			layout.Flex{
				Axis: layout.Horizontal,
			}.Layout(w.Context,
				layout.Flexed(0.4, w.PrikazaniElementIzgled()),
				layout.Flexed(0.6, w.SumaIzgled()),
			)
		}
	}
}

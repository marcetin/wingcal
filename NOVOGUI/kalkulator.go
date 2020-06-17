package main

import (
	"gioui.org/layout"
	"github.com/marcetin/wingcal/NOVOGUI/calc"
)

func kalkulator(w *calc.WingCal) func() {
	return func() {
		if w.Edit {
			layout.Flex{
				Axis: layout.Horizontal,
			}.Layout(w.Context,
				layout.Flexed(1, w.EditorElementaIzgled()),
			)
		}
		levo := w.IzborVrsteRadova()
		if w.Element {
			levo = w.PrikazaniElementIzgled()
		}

		layout.Flex{
			Axis: layout.Horizontal,
		}.Layout(w.Context,
			layout.Flexed(0.5, levo),
			layout.Flexed(0.5, w.SumaIzgled()),
		)

	}
}

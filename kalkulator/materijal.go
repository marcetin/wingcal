package kalkulator

import (
	"fmt"
	"gioui.org/layout"
	"github.com/marcetin/wingcal/calc"
)

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
				layout.Rigid(w.Tema.DuoUIline(w.Context, 0, 0, 4, "Dark")),
			)
		})
	}
}

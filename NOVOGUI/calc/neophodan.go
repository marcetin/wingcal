package calc

import (
	"fmt"
	"gioui.org/layout"
	"github.com/marcetin/wingcal/pkg/latcyr"
)

func (w *WingCal) RadNeophodanMaterijal(l *layout.List) func() {
	return func() {
		//var materijal model.WingNeophodanMaterijal
		nm := w.PrikazaniElement.NeophodanMaterijal
		width := w.Context.Constraints.Width.Max
		l.Layout(w.Context, len(nm), func(i int) {
			materijal := nm[i]
			id := materijal.Id - 1
			materijal.Koeficijent = materijal.Koeficijent
			materijal.Materijal = *w.Materijal[id]
			if materijal.Koeficijent > 0 {
				materijal.Kolicina = materijal.Materijal.Potrosnja * float64(kolicina.Value) * materijal.Koeficijent
			}
			materijal.UkupnaCena = materijal.Materijal.Cena * float64(materijal.Kolicina)
			materijal.UkupnoPakovanja = int(materijal.Kolicina / float64(materijal.Materijal.Pakovanje))

			w.Context.Constraints.Width.Min = width

			layout.Flex{
				Axis: layout.Vertical,
			}.Layout(w.Context,
				layout.Rigid(func() {
					layout.Flex{
						Axis:    layout.Horizontal,
						Spacing: layout.SpaceBetween,
					}.Layout(w.Context,
						layout.Flexed(0.4, func() {
							w.Tema.Body1(latcyr.C(materijal.Materijal.Naziv, w.Cyr)).Layout(w.Context)
						}),
						layout.Flexed(0.15, func() {
							w.Tema.Body1(fmt.Sprintf("%.2f", materijal.Materijal.Potrosnja)).Layout(w.Context)
						}),
						layout.Flexed(0.15, func() {
							w.Tema.Body1(fmt.Sprint(materijal.Koeficijent)).Layout(w.Context)
						}),
						layout.Flexed(0.15, func() {
							w.Tema.Body1(fmt.Sprintf("%.2f", materijal.Kolicina)).Layout(w.Context)
						}),
						layout.Flexed(0.15, func() {
							w.Tema.Body1(fmt.Sprintf("%.2f", materijal.UkupnaCena)).Layout(w.Context)
						}),
					)
				}),
				layout.Rigid(w.Tema.DuoUIline(w.Context, 0, 0, 1, w.Tema.Colors["Gray"])),
			)
		})
	}
}

func (w *WingCal) UkupanNeophodanMaterijal(l *layout.List) func() {
	return func() {

		//fmt.Println(":::::UkupanNeophodanMaterijal", w.Suma.UkupanNeophodanMaterijal)
		//var materijal model.WingNeophodanMaterijal
		width := w.Context.Constraints.Width.Max
		l.Layout(w.Context, len(w.Suma.UkupanNeophodanMaterijal), func(i int) {

			materijal := w.Suma.UkupanNeophodanMaterijalPrikaz[i]
			//materijal.Materijal = *w.Materijal[materijal.Id-1]
			//fmt.Println(":::::NazivNazivNaziv", materijal.Materijal.Naziv)
			//fmt.Println(":::::IDidididi111", materijal.Id)
			//fmt.Println(":::::IDidididi", materijal.Materijal.Id)
			w.Context.Constraints.Width.Min = width

			layout.Flex{
				Axis: layout.Vertical,
			}.Layout(w.Context,
				layout.Rigid(func() {
					layout.Flex{
						Axis:    layout.Horizontal,
						Spacing: layout.SpaceBetween,
					}.Layout(w.Context,
						layout.Flexed(0.4, func() {
							w.Tema.Body1(latcyr.C(materijal.Materijal.Naziv, w.Cyr)).Layout(w.Context)
						}),
						layout.Flexed(0.15, func() {
							w.Tema.Body1(fmt.Sprint(materijal.Materijal.Cena)).Layout(w.Context)
						}),
						layout.Flexed(0.15, func() {
							w.Tema.Body1(fmt.Sprintf("%.2f", materijal.Kolicina)).Layout(w.Context)
						}),
						layout.Flexed(0.15, func() {
							w.Tema.Body1(fmt.Sprintf("%.2f", materijal.UkupnaCena)).Layout(w.Context)
						}),
					)
				}),
				layout.Rigid(w.Tema.DuoUIline(w.Context, 0, 0, 1, w.Tema.Colors["Gray"])),
			)
		})
	}
}

func (w *WingCal) SumaStavkeMaterijala(width int) func() {
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
				w.Tema.Caption(latcyr.C("Pojedinacna cena", w.Cyr)).Layout(w.Context)
			}),
			layout.Flexed(0.15, func() {
				w.Tema.Caption(latcyr.C("Kolicina", w.Cyr)).Layout(w.Context)
			}),
			layout.Flexed(0.15, func() {
				w.Tema.Caption(latcyr.C("Ukupna cena", w.Cyr)).Layout(w.Context)
			}))
	}
}

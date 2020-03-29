package calc

import (
	"fmt"
	"gioui.org/layout"
	"gioui.org/unit"
	"github.com/marcetin/wingcal/model"
	"github.com/marcetin/wingcal/pkg/gel"
	"github.com/marcetin/wingcal/pkg/gelook"
)

func (w *WingCal) EditorElementaIzgled() func() {
	return func() {
		layout.Flex{
			Axis: layout.Vertical,
		}.Layout(w.Context,
			layout.Rigid(func() {
				w.Tema.DuoUIcontainer(8, w.Tema.Colors["LightGray"]).Layout(w.Context, layout.W, func() {
					w.Tema.H5(w.PrikazaniElement.Naziv).Layout(w.Context)
				})
			}),
			layout.Flexed(1, func() {
				w.Tema.DuoUIcontainer(8, w.Tema.Colors["LightGray"]).Layout(w.Context, layout.NW, func() {
					//sumaCena := 111.33
					//layout.Flex{Axis: layout.Vertical}.Layout(w.Context,
					//	layout.Rigid(func() {
					//		w.Tema.Body1(w.PrikazaniElement.Opis).Layout(w.Context)
					//	}),
					//	layout.Rigid(func() {
					//		w.Tema.Caption(w.PrikazaniElement.Obracun).Layout(w.Context)
					//	}),
					//	layout.Rigid(w.Tema.DuoUIline(w.Context, 32, "Dark")),
					//
					//	layout.Rigid(func() {
					//		w.Tema.H6("Neophodan materijal za izvrsenje radova").Layout(w.Context)
					//	}),
					//	layout.Rigid(w.NeophodanMaterijal(neophodanMaterijalList, w.PrikazaniElement.NeophodanMaterijal)),
					//)
					layout.Flex{Axis: layout.Vertical}.Layout(w.Context,
						layout.Rigid(func() {
							layout.Flex{Axis: layout.Vertical}.Layout(w.Context,
								layout.Rigid(Editor(w.Context, w.Tema, w.EditPolja.Id, "Id", func(e gel.EditorEvent) {})),
								layout.Rigid(Editor(w.Context, w.Tema, w.EditPolja.Naziv, "Naziv", func(e gel.EditorEvent) {})),
								layout.Rigid(Editor(w.Context, w.Tema, w.EditPolja.Obracun, "Obracun", func(e gel.EditorEvent) {})),
								layout.Rigid(Editor(w.Context, w.Tema, w.EditPolja.Jedinica, "Jedinica", func(e gel.EditorEvent) {})),
								layout.Rigid(Editor(w.Context, w.Tema, w.EditPolja.Cena, "Cena", func(e gel.EditorEvent) {})),
								layout.Rigid(Editor(w.Context, w.Tema, w.EditPolja.Opis, "Opis", func(e gel.EditorEvent) {})))
						}),

						//layout.Rigid(Editor(w.Context, w.Tema, w.EditabilnaPoljaVrsteRadova[w.PrikazaniElement.Id].Opis, fmt.Sprint(w.PrikazaniElement.PodvrsteRadova[w.PrikazaniElement.Id].Opis), func(e gel.EditorEvent) {})),
					)
				})
			}),
			layout.Rigid(func() {
				w.Tema.DuoUIcontainer(0, w.Tema.Colors["Gray"]).Layout(w.Context, layout.NW, func() {
					sumaCena := float64(kolicina.Value) * w.PrikazaniElement.Cena
					layout.Flex{
						Axis:    layout.Horizontal,
						Spacing: layout.SpaceBetween,
					}.Layout(w.Context,
						layout.Rigid(func() {
							layout.Flex{
								Axis: layout.Vertical,
							}.Layout(w.Context,
								layout.Rigid(func() {
									w.Tema.DuoUIcounter(func() {}).Layout(w.Context, kolicina, "KOLICINA", fmt.Sprint(kolicina.Value))
								}),
								layout.Rigid(func() {
									btn := w.Tema.Button("DODAJ")
									//btn.FullWidth = true
									//btn.FullHeight = true
									btn.Background = gelook.HexARGB(w.Tema.Colors["Secondary"])
									suma := model.WingIzabraniElement{
										Kolicina: kolicina.Value,
										SumaCena: sumaCena,
										Element:  *w.PrikazaniElement,
									}
									for dodajDugme.Clicked(w.Context) {
										w.Suma.Elementi = append(w.Suma.Elementi, suma)
										for _, n := range w.PrikazaniElement.NeophodanMaterijal {
											w.Suma.UkupanNeophodanMaterijal[n.Id] = model.WingNeophodanMaterijal{
												Id:       n.Id,
												Kolicina: w.Suma.UkupanNeophodanMaterijal[n.Id].Kolicina + n.Kolicina*kolicina.Value,
											}
										}
										//var neophodanmaterijal map[int]model.WingNeophodanMaterijal
									}
									btn.Layout(w.Context, dodajDugme)
								}),
							)
						}))
				})
			}))
	}
}

func Editor(gtx *layout.Context, th *gelook.DuoUItheme, editorControler *gel.Editor, label string, handler func(gel.EditorEvent)) func() {
	return func() {
		layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func() {
				th.H6(label).Layout(gtx)
			}),
			layout.Rigid(func() {
				th.DuoUIcontainer(8, "ffffffff").Layout(gtx, layout.NW, func() {
					e := th.DuoUIeditor(label)
					e.Font.Typeface = th.Fonts["Mono"]
					e.TextSize = unit.Dp(12)
					layout.UniformInset(unit.Dp(4)).Layout(gtx, func() {
						e.Layout(gtx, editorControler)
					})
					for _, e := range editorControler.Events(gtx) {
						switch e.(type) {
						case gel.ChangeEvent:
							handler(e)
						}
					}
				})
			}))
	}
}

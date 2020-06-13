package calc

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"github.com/gioapp/gel"
	"github.com/gioapp/gelook"
)

var (
	materijalElementaPanelElement = gel.NewPanel()
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
								layout.Rigid(Editor(w.Context, w.Tema, w.EditPolja.Naziv, layout.Vertical, "Naziv", func(e gel.EditorEvent) {})),
								layout.Rigid(func() {
									layout.Flex{
										Axis:    layout.Horizontal,
										Spacing: layout.SpaceBetween,
									}.Layout(w.Context,
										layout.Rigid(Editor(w.Context, w.Tema, w.EditPolja.Id, layout.Vertical, "Id", func(e gel.EditorEvent) {})),
										layout.Rigid(Editor(w.Context, w.Tema, w.EditPolja.Obracun, layout.Vertical, "Obracun", func(e gel.EditorEvent) {})),
										layout.Rigid(Editor(w.Context, w.Tema, w.EditPolja.Jedinica, layout.Vertical, "Jedinica", func(e gel.EditorEvent) {})),
										layout.Rigid(Editor(w.Context, w.Tema, w.EditPolja.Cena, layout.Vertical, "Cena", func(e gel.EditorEvent) {})))
								}),
								layout.Rigid(Editor(w.Context, w.Tema, w.EditPolja.Opis, layout.Vertical, "Opis", func(e gel.EditorEvent) {})),
								layout.Flexed(1, func() {
									materijalElementaPanelElement.PanelObject = w.Materijal
									materijalElementaPanelElement.PanelObjectsNumber = len(w.Materijal)
									materijalElementaPanel := w.Tema.DuoUIpanel()
									materijalElementaPanel.ScrollBar = w.Tema.ScrollBar(0)
									materijalElementaPanel.Layout(w.Context, materijalElementaPanelElement, func(i int, in interface{}) {
										//if in != nil {
										//addresses := in.([]model.DuoUIaddress)
										materijal := w.Materijal[i]
										layout.Flex{Axis: layout.Vertical}.Layout(w.Context,
											layout.Rigid(func() {
												layout.Flex{
													Alignment: layout.Middle,
												}.Layout(w.Context,
													layout.Rigid(Editor(w.Context, w.Tema, w.EditPolja.Materijal[materijal.Id], layout.Horizontal, materijal.Naziv, func(e gel.EditorEvent) {})),
													layout.Rigid(func() {
														w.Tema.Body1(materijal.Jedinica).Layout(w.Context)
													}),
													layout.Rigid(func() {
														w.Tema.Caption(materijal.Obracun).Layout(w.Context)
													}),
												)
											}),
											layout.Rigid(w.Tema.DuoUIline(w.Context, 1, 0, 1, w.Tema.Colors["Gray"])),
										)
										//}
									})
								}),
							)
						}),

						//layout.Rigid(Editor(w.Context, w.Tema, w.EditabilnaPoljaVrsteRadova[w.PrikazaniElement.Id].Opis, fmt.Sprint(w.PrikazaniElement.PodvrsteRadova[w.PrikazaniElement.Id].Opis), func(e gel.EditorEvent) {})),
					)
				})
			}),
			layout.Rigid(func() {

			}))
	}
}

func Editor(gtx *layout.Context, th *gelook.DuoUItheme, editorControler *gel.Editor, axis layout.Axis, label string, handler func(gel.EditorEvent)) func() {
	return func() {
		layout.Flex{Axis: axis}.Layout(gtx,
			layout.Rigid(func() {
				th.H6(label).Layout(gtx)
			}),
			layout.Rigid(func() {
				th.DuoUIcontainer(8, "ffffffff").Layout(gtx, layout.NW, func() {
					//width := gtx.Constraints.Width.Max
					width := 555
					e := th.DuoUIeditor(label, "", "", width)
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

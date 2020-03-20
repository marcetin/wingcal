package main

import (
	"fmt"
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"github.com/marcetin/wingcal/calc"
	"github.com/marcetin/wingcal/pkg/gel"
	"github.com/marcetin/wingcal/pkg/gelook"
)

var (
	izbornikRadova = &layout.List{
		Axis: layout.Vertical,
	}
	editLista = &layout.List{
		Axis: layout.Vertical,
	}
	thingEditTitle   = new(gel.Editor)
	thingEditContent = new(gel.Editor)
	//post             = new(model.DuoCMSpost)
	dodajDugme   = new(gel.Button)
	stampajDugme = new(gel.Button)
	nazadDugme   = new(gel.Button)
	kolicina     = &gel.DuoUIcounter{
		Value:        1,
		OperateValue: 1,
		From:         1,
		To:           100,
		CounterInput: &gel.Editor{
			Alignment:  text.Middle,
			SingleLine: true,
		},
		CounterIncrease: new(gel.Button),
		CounterDecrease: new(gel.Button),
		CounterReset:    new(gel.Button),
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

func admin(w *calc.WingCal) func() {
	return func() {

		//post := &model.DuoCMSpost{
		//	Title:    "",
		//	Subtitle: "",
		//}

		layout.Flex{
			Axis: layout.Vertical,
		}.Layout(w.Context,
			layout.Rigid(func() {
				w.Tema.DuoUIitem(0, w.Tema.Colors["Secondary"]).Layout(w.Context, layout.Center, func() {
					w.Tema.H3("EDITOR").Layout(w.Context)
				})
			}),
			layout.Flexed(1, func() {
				layout.Flex{
					Axis: layout.Horizontal,
				}.Layout(w.Context,
					layout.Rigid(IzborVrsteRadova(w)),
					layout.Flexed(1, func() {
						layout.Flex{
							Axis: layout.Vertical,
						}.Layout(w.Context,
							layout.Rigid(func() {
								w.Tema.DuoUIitem(0, w.Tema.Colors["Danger"]).Layout(w.Context, layout.Center, func() {
									w.Tema.H3("Content header").Layout(w.Context)
								})
							}),
							layout.Flexed(1, func() {
								w.Tema.DuoUIitem(0, w.Tema.Colors["Secondary"]).Layout(w.Context, layout.N, func() {

									editLista.Layout(w.Context, len(w.PrikazaniElement.PodvrsteRadova), func(i int) {
										layout.Flex{Axis: layout.Vertical}.Layout(w.Context,
											layout.Rigid(func() {
												layout.Flex{Spacing: layout.SpaceBetween}.Layout(w.Context,
													layout.Rigid(Editor(w.Context, w.Tema, w.EditabilnaPoljaVrsteRadova[i].Id, fmt.Sprint(w.PrikazaniElement.PodvrsteRadova[i].Id), func(e gel.EditorEvent) {})),
													layout.Rigid(Editor(w.Context, w.Tema, w.EditabilnaPoljaVrsteRadova[i].Naziv, fmt.Sprint(w.PrikazaniElement.PodvrsteRadova[i].Naziv), func(e gel.EditorEvent) {})),
													layout.Rigid(Editor(w.Context, w.Tema, w.EditabilnaPoljaVrsteRadova[i].Obracun, fmt.Sprint(w.PrikazaniElement.PodvrsteRadova[i].Obracun), func(e gel.EditorEvent) {})),
													layout.Rigid(Editor(w.Context, w.Tema, w.EditabilnaPoljaVrsteRadova[i].Jedinica, fmt.Sprint(w.PrikazaniElement.PodvrsteRadova[i].Jedinica), func(e gel.EditorEvent) {})),
													layout.Rigid(Editor(w.Context, w.Tema, w.EditabilnaPoljaVrsteRadova[i].Cena, fmt.Sprint(w.PrikazaniElement.PodvrsteRadova[i].Id), func(e gel.EditorEvent) {})))
											}),
											layout.Rigid(Editor(w.Context, w.Tema, w.EditabilnaPoljaVrsteRadova[i].Opis, fmt.Sprint(w.PrikazaniElement.PodvrsteRadova[i].Opis), func(e gel.EditorEvent) {})),
										)
									})

								})

							}),
							layout.Rigid(func() {
								w.Tema.DuoUIitem(0, w.Tema.Colors["Primary"]).Layout(w.Context, layout.Center, func() {
									w.Tema.H3("Content footer").Layout(w.Context)
								})
							}))
					}))
			}),
			layout.Rigid(func() {
				w.Tema.DuoUIitem(0, w.Tema.Colors["DarkGray"]).Layout(w.Context, layout.Center, func() {
					w.Tema.H3("Footer").Layout(w.Context)
				})
			}))
	}
}

func IzborVrsteRadova(w *calc.WingCal) func() {
	return func() {
		if w.Context.Constraints.Width.Max > 300 {
			w.Context.Constraints.Width.Max = 300
		}
		w.Tema.DuoUIitem(0, w.Tema.Colors["Warning"]).Layout(w.Context, layout.N, func() {
			layout.Flex{Axis: layout.Vertical}.Layout(w.Context,
				layout.Rigid(w.NazivRoditelja()),
				layout.Rigid(w.Nazad()),
				layout.Flexed(1, func() {
					izbornikRadova.Layout(w.Context, len(w.IzbornikRadova.PodvrsteRadova), func(i int) {
						//*w.IzbornikRadova.Izbor.PodvrsteRadova[i].Roditelj =w.IzbornikRadova.Izbor
						vrstarada := w.IzbornikRadova.PodvrsteRadova[i]
						vrstarada.Roditelj = &w.IzbornikRadova
						layout.UniformInset(unit.Dp(0)).Layout(w.Context, func() {
							layout.Flex{Axis: layout.Vertical}.Layout(w.Context,
								layout.Rigid(func() {
									btn := w.Tema.Button(vrstarada.Naziv)

									for w.LinkoviIzboraVrsteRadova[i].Clicked(w.Context) {
										for _, rad := range vrstarada.PodvrsteRadova {
											rad.Roditelj = &w.IzbornikRadova
										}
										if vrstarada.Baza {
											vrstarada = w.Db.DbRead(vrstarada.Roditelj.Slug, vrstarada.Slug)
											fmt.Println("Baza IMA")
											w.PrikazaniElement = &vrstarada
											w.GenerisanjeEdita(vrstarada.PodvrsteRadova)
										} else {

											//if vrstarada.PodvrsteRadova != nil {

											w.Podvrsta(vrstarada)
										}
										//}
									}
									btn.Layout(w.Context, w.LinkoviIzboraVrsteRadova[i])
								}),
							)
						})
					})
				}))
		})
	}
}

func Editor(gtx *layout.Context, th *gelook.DuoUItheme, editorControler *gel.Editor, label string, handler func(gel.EditorEvent)) func() {
	return func() {
		layout.UniformInset(unit.Dp(4)).Layout(gtx, func() {
			//gelook.DuoUIdrawRectangle(gtx, cs.Width.Max, cs.Height.Max,
			//	th.Colors["Gray"],
			//	[4]float32{0, 0, 0, 0}, [4]float32{0, 0, 0, 0})
			layout.UniformInset(unit.Dp(0)).Layout(gtx, func() {
				//cs := gtx.Constraints
				gelook.DuoUIdrawRectangle(gtx, 30, 30,
					th.Colors["Light"], [4]float32{0, 0, 0, 0}, [4]float32{0, 0,
						0, 0})
				e := th.DuoUIeditor(label)
				e.Font.Typeface = th.Fonts["Mono"]
				// e.Font.Style = text.Italic
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
		})
	}
}

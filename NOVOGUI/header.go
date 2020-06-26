package main

import (
	"fmt"
	"gioui.org/layout"
	"gioui.org/unit"
	"github.com/gioapp/gel"
	"github.com/marcetin/wingcal/NOVOGUI/calc"
	"github.com/marcetin/wingcal/pkg/gelook"
	"github.com/marcetin/wingcal/pkg/latcyr"
)

var (
	latcyrDugme     = new(gel.Button)
	kalkulatorDugme = new(gel.Button)
	materijalDugme  = new(gel.Button)
	projekatDugme   = new(gel.Button)
	editorDugme     = new(gel.Button)

	putanjaRadoviDugme   = new(gel.Button)
	putanjaPodvrstaDugme = new(gel.Button)
	putanjaElementDugme  = new(gel.Button)

	headerMenuList = &layout.List{
		Axis: layout.Horizontal,
	}
)

func header(w *calc.WingCal) func() {
	return func() {
		layout.Flex{
			Axis:    layout.Horizontal,
			Spacing: layout.SpaceBetween,
		}.Layout(w.Context,
			layout.Flexed(0.5, func() {
				putanjaList.Layout(w.Context, len(w.Putanja), func(i int) {
					switch i {
					case 0:
						w.Tema.Button(latcyr.C(w.Putanja[i], w.Cyr)).Layout(w.Context, putanjaRadoviDugme)
						for putanjaRadoviDugme.Clicked(w.Context) {

							komanda := ""
							//if len(w.Putanja) == 3 {
							//	komanda = "/" + fmt.Sprint(w.Roditelj)
							//	//podvrstaradova = fmt.Sprint(w.Roditelj)
							//	fmt.Println("roddddditeL111::", w.Roditelj)
							//}
							//if len(w.Putanja) == 4 {
							//	komanda = "/" + calc.Podvrstaradova + "/" + fmt.Sprint(w.Roditelj)
							//}
							w.APIpozivIzbornik("radovi" + komanda)
							//w.LinkoviIzboraVrsteRadova = GenerisanjeLinkova(w.IzbornikRadova)
							w.GenerisanjeLinkova(w.IzbornikRadova)
							//w.Putanja = nil
							//w.Putanja = append(w.Putanja, w.Radovi.Naziv)

						}
					case 1:
						w.Tema.Button(latcyr.C(w.Putanja[i], w.Cyr)).Layout(w.Context, putanjaPodvrstaDugme)
						for putanjaPodvrstaDugme.Clicked(w.Context) {
							komanda := ""

							komanda = "/" + fmt.Sprint(w.Roditelj)
							//podvrstaradova = fmt.Sprint(w.Roditelj)
							//fmt.Println("roddddditeL111::", w.Roditelj)
							//komanda = "/" + calc.Podvrstaradova + "/" + fmt.Sprint(w.Roditelj)

							w.APIpozivIzbornik("radovi" + komanda)
							//w.LinkoviIzboraVrsteRadova = GenerisanjeLinkova(w.IzbornikRadova)
							w.GenerisanjeLinkova(w.IzbornikRadova)
							w.Putanja = w.Putanja[:len(w.Putanja)-1]

						}
					case 2:
						w.Tema.Button(latcyr.C(w.Putanja[i], w.Cyr)).Layout(w.Context, putanjaElementDugme)
						for putanjaElementDugme.Clicked(w.Context) {
							w.Element = false
						}
					}
				})
			}),
			layout.Flexed(0.5, func() {

				headerMenu := []func(){
					func() {
						if w.PrikazaniElement.Element {
							btnEditor := w.Tema.Button(latcyr.C("EDITOR", w.Cyr))
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
						btnKalkulator := w.Tema.Button(latcyr.C("KALKULATOR", w.Cyr))
						btnKalkulator.Background = gelook.HexARGB(w.Tema.Colors["Secondary"])
						for kalkulatorDugme.Clicked(w.Context) {
							w.Strana = "kalkulator"
						}
						btnKalkulator.Layout(w.Context, kalkulatorDugme)
					},
					func() {
						btnMaterijal := w.Tema.Button(latcyr.C("MATERIJAL", w.Cyr))
						btnMaterijal.Background = gelook.HexARGB(w.Tema.Colors["Secondary"])
						for materijalDugme.Clicked(w.Context) {
							w.Strana = "materijal"
						}
						btnMaterijal.Layout(w.Context, materijalDugme)
					},
					func() {
						btnMaterijal := w.Tema.Button(latcyr.C("PROJEKAT", w.Cyr))
						btnMaterijal.Background = gelook.HexARGB(w.Tema.Colors["Secondary"])
						for projekatDugme.Clicked(w.Context) {
							w.Strana = "projekat"
						}
						btnMaterijal.Layout(w.Context, projekatDugme)
					},
				}
				headerMenuList.Layout(w.Context, len(headerMenu), func(i int) {
					layout.UniformInset(unit.Dp(0)).Layout(w.Context, headerMenu[i])
				})
			}),
			layout.Rigid(func() {

				latcyr := "Ћирилица"
				if w.Cyr {
					latcyr = "Latinica"
				}

				btnLatcyr := w.Tema.Button(latcyr)
				btnLatcyr.Background = gelook.HexARGB(w.Tema.Colors["Secondary"])
				for latcyrDugme.Clicked(w.Context) {
					if w.Cyr {
						w.Cyr = false
					} else {
						w.Cyr = true
					}
				}
				btnLatcyr.Layout(w.Context, latcyrDugme)
			}))
	}
}

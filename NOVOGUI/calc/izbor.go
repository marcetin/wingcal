package calc

import (
	"fmt"
	"gioui.org/layout"
	"gioui.org/unit"
	"github.com/gioapp/gel"
	"github.com/marcetin/wingcal/pkg/gelook"
	"github.com/marcetin/wingcal/pkg/latcyr"
)

var (
	IzborVrsteRadovaPanelElement = gel.NewPanel()
	podvrstaradova               string
	elementi                     string
)

func (w *WingCal) IzborVrsteRadova() func() {
	return func() {
		//if w.Context.Constraints.Width.Max > 300 {
		//w.Context.Constraints.Width.Min = 300
		//w.Context.Constraints.Width.Max = 300
		//}
		w.Tema.DuoUIcontainer(0, w.Tema.Colors["DarkGrayI"]).Layout(w.Context, layout.W, func() {
			layout.Flex{Axis: layout.Vertical}.Layout(w.Context,
				//layout.Rigid(w.NazivRoditelja()),
				layout.Rigid(w.Nazad()),
				layout.Flexed(1, func() {

					IzborVrsteRadovaPanelElement.PanelObject = w.IzbornikRadova
					IzborVrsteRadovaPanelElement.PanelObjectsNumber = len(w.IzbornikRadova)
					izborVrsteRadovaPanel := w.Tema.DuoUIpanel()
					izborVrsteRadovaPanel.ScrollBar = w.Tema.ScrollBar(0)
					izborVrsteRadovaPanel.Layout(w.Context, IzborVrsteRadovaPanelElement, func(i int, in interface{}) {
						//if in != nil {
						//addresses := in.([]model.DuoUIaddress)

						//*w.IzbornikRadova.Izbor.PodvrsteRadova[i].Roditelj =w.IzbornikRadova.Izbor
						vrstarada := w.IzbornikRadova[i]
						//if vrstarada.Element && vrstarada.NeophodanMaterijal != nil  {
						//vrstarada.Roditelj = &w.IzbornikRadova
						layout.UniformInset(unit.Dp(0)).Layout(w.Context, func() {
							layout.Flex{Axis: layout.Vertical}.Layout(w.Context,
								layout.Rigid(func() {

									btn := w.Tema.Button(latcyr.C(vrstarada.Title, w.Cyr))

									btn.CornerRadius = unit.Dp(0)
									btn.Background = gelook.HexARGB(w.Tema.Colors["Gray"])
									if vrstarada.Materijal {
										btn.Background = gelook.HexARGB(w.Tema.Colors["DarkGray"])
									}
									for w.LinkoviIzboraVrsteRadova[i].Clicked(w.Context) {
										fmt.Println(i)

										komanda := fmt.Sprint(i + 1)

										if len(w.Putanja) == 1 {
											komanda = fmt.Sprint(i + 1)
											podvrstaradova = fmt.Sprint(i + 1)
											w.Podvrsta = i + 1
										}
										if len(w.Putanja) == 2 {
											komanda = podvrstaradova + "/" + fmt.Sprint(i+1)
											elementi = fmt.Sprint(i + 1)
											w.Roditelj = i + 1
										}
										if len(w.Putanja) == 3 {
											komanda = podvrstaradova + "/" + elementi + "/" + fmt.Sprint(i+1)
										}
										if len(w.Putanja) == 1 {
											w.APIpozivIzbornik("radovi/" + komanda)
										}
										if len(w.Putanja) == 2 {
											w.APIpozivElementi("radovi/" + komanda)
										}
										if len(w.Putanja) == 3 {
											w.APIpozivElement("radovi/" + komanda)
											w.Element = true
										}

										if len(w.Putanja) < 3 {
											w.Putanja = append(w.Putanja, vrstarada.Title)
										}
										//if vrstarada.Baza {
										//	elementi := w.Db.DbRead(w.IzbornikRadova.Slug, vrstarada.Slug)
										//	vrstarada.PodvrsteRadova = elementi.PodvrsteRadova
										//}
										//if vrstarada.PodvrsteRadova != nil {
										//w.Podvrsta(&vrstarada)
										//w.EditPolja.Id.SetText(fmt.Sprint(w.PrikazaniElement.Id))
										//w.EditPolja.Naziv.SetText(w.PrikazaniElement.Naziv)
										//w.EditPolja.Opis.SetText(w.PrikazaniElement.Opis)
										//w.EditPolja.Obracun.SetText(w.PrikazaniElement.Obracun)
										//w.EditPolja.Jedinica.SetText(fmt.Sprint(w.PrikazaniElement.Jedinica))
										//w.EditPolja.Cena.SetText(fmt.Sprint(w.PrikazaniElement.Cena))
										//for _, neophodanMaterijal := range w.PrikazaniElement.NeophodanMaterijal {
										//	w.EditPolja.Materijal[neophodanMaterijal.Id].SetText("0")
										//	w.EditPolja.Materijal[neophodanMaterijal.Id].SetText(fmt.Sprint(neophodanMaterijal.Kolicina))
										//}
										//}

										//w.LinkoviIzboraVrsteRadova = GenerisanjeLinkova(w.IzbornikRadova)
										w.GenerisanjeLinkova(w.IzbornikRadova)
										kolicina.Value = 0

									}

									btn.Layout(w.Context, w.LinkoviIzboraVrsteRadova[i])
								}),
								layout.Rigid(w.Tema.DuoUIline(w.Context, 0, 0, 0, w.Tema.Colors["Gray"])),
							)
						})
						//}
					})
				}))
		})
	}
}

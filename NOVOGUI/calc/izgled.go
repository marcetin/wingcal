package calc

import (
	"encoding/json"
	"fmt"
	"gioui.org/layout"
	"gioui.org/unit"
	"github.com/gioapp/gel"
	"github.com/gioapp/gelook"
	"github.com/marcetin/wingcal/model"
	"io/ioutil"
	"log"
	"net/http"
	"time"
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
									btn := w.Tema.Button(vrstarada)
									btn.CornerRadius = unit.Dp(0)
									btn.Background = gelook.HexARGB(w.Tema.Colors["Gray"])
									for w.LinkoviIzboraVrsteRadova[i].Clicked(w.Context) {
										fmt.Println(i)

										komanda := fmt.Sprint(i + 1)

										if len(w.Putanja) == 1 {
											komanda = fmt.Sprint(i + 1)
											podvrstaradova = fmt.Sprint(i + 1)
										}
										if len(w.Putanja) == 2 {
											komanda = podvrstaradova + "/" + fmt.Sprint(i+1)
											elementi = fmt.Sprint(i + 1)
										}
										if len(w.Putanja) == 3 {
											komanda = podvrstaradova + "/" + elementi + "/" + fmt.Sprint(i+1)
										}
										if len(w.Putanja) < 3 {
											w.APIpozivIzbornik("radovi/" + komanda)
										}
										if len(w.Putanja) == 3 {
											w.APIpozivElement("radovi/" + komanda)
										}

										if len(w.Putanja) < 3 {
											w.Putanja = append(w.Putanja, vrstarada)
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
										w.GenerisanjeLinkova(w.IzbornikRadova)

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

func (w *WingCal) Nazad() func() {
	return func() {
		if len(w.Putanja) > 0 {
			btnNazad := w.Tema.Button("NAZAD")
			btnNazad.Background = gelook.HexARGB(w.Tema.Colors["Secondary"])
			for nazadDugme.Clicked(w.Context) {
				//w.IzbornikRadova = w.Putanja[len(w.Putanja)-1]
				//w.GenerisanjeLinkova(w.Putanja[len(w.Putanja)-1].PodvrsteRadova)
				w.Putanja = w.Putanja[:len(w.Putanja)-1]
				//w.Roditelj()
				//fmt.Println("IzbornikroditeL::" + w.IzbornikRadova)
				//fmt.Println("roditeL::" + w.IzbornikRadova.Roditelj.Slug)
			}
			btnNazad.Layout(w.Context, nazadDugme)
		}
	}
}

func (w *WingCal) NeophodanMaterijal(l *layout.List, n map[int]model.WingNeophodanMaterijal, s bool) func() {
	return func() {
		var materijal model.WingNeophodanMaterijal
		width := w.Context.Constraints.Width.Max
		l.Layout(w.Context, len(n), func(i int) {
			if s {
				materijal = n[i]
			} else {
				materijal = n[i-1]
				materijal.Koeficijent = n[i].Koeficijent
			}

			materijal.Materijal = w.Materijal[materijal.Id]
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
						layout.Flexed(0.2, func() {
							w.Tema.Body1(materijal.Materijal.Naziv).Layout(w.Context)
						}),
						layout.Flexed(0.2, func() {
							w.Tema.Body1(fmt.Sprintf("%.2f", materijal.Materijal.Potrosnja)).Layout(w.Context)
						}),
						layout.Flexed(0.15, func() {
							w.Tema.Body1(fmt.Sprint(materijal.Koeficijent)).Layout(w.Context)
						}),
						layout.Flexed(0.15, func() {
							w.Tema.Body1(fmt.Sprintf("%.2f", materijal.Kolicina)).Layout(w.Context)
						}),
						layout.Flexed(0.15, func() {
							w.Tema.Body1(fmt.Sprint(materijal.UkupnoPakovanja)).Layout(w.Context)
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

func (w *WingCal) APIpozivIzbornik(komanda string) {
	radovi := map[int]string{}
	jsonErr := json.Unmarshal(APIpoziv(komanda), &radovi)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	w.IzbornikRadova = radovi
}

func (w *WingCal) APIpozivElement(komanda string) {
	rad := &model.WingVrstaRadova{}
	jsonErr := json.Unmarshal(APIpoziv(komanda), &rad)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	w.PrikazaniElement = rad
}

func APIpoziv(komanda string) []byte {
	url := "http://212.62.35.158:9909/" + komanda
	fmt.Println("url", url)
	spaceClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", "wing")
	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	if body != nil {
		//defer body.Close()
	}
	return body
}

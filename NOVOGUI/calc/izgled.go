package calc

import (
	"fmt"
	"github.com/marcetin/wingcal/pkg/gelook"
	"github.com/marcetin/wingcal/pkg/latcyr"
)

func (w *WingCal) Nazad() func() {
	return func() {
		if len(w.Putanja) > 1 {
			btnNazad := w.Tema.Button(latcyr.C("NAZAD", w.Cyr))
			btnNazad.Background = gelook.HexARGB(w.Tema.Colors["Secondary"])
			for nazadDugme.Clicked(w.Context) {
				komanda := ""
				if len(w.Putanja) == 3 {
					komanda = "/" + fmt.Sprint(w.Roditelj)
					//podvrstaradova = fmt.Sprint(w.Roditelj)
					fmt.Println("roddddditeL111::", w.Roditelj)
				}
				if len(w.Putanja) == 4 {
					komanda = "/" + podvrstaradova + "/" + fmt.Sprint(w.Roditelj)
				}
				w.APIpozivIzbornik("radovi" + komanda)
				//w.LinkoviIzboraVrsteRadova = GenerisanjeLinkova(w.IzbornikRadova)
				w.GenerisanjeLinkova(w.IzbornikRadova)
				w.Putanja = w.Putanja[:len(w.Putanja)-1]
			}
			btnNazad.Layout(w.Context, nazadDugme)
		}
	}
}

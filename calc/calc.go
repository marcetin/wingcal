package calc

import (
	"fmt"
	"github.com/gioapp/gel"
	"github.com/marcetin/wingcal/model"
)

func (w *WingCal) GenerisanjeLinkova(radovi map[int]model.WingVrstaRadova) {
	for rad, _ := range radovi {
		w.LinkoviIzboraVrsteRadova[rad] = new(gel.Button)
	}
	return
}

//func  (w *WingCal)Roditelj(roditelj *model.WingVrstaRadova,radovi map[int]model.WingVrstaRadova) {
func (w *WingCal) Roditelj() {
	for _, rad := range w.IzbornikRadova.PodvrsteRadova {
		rad.Roditelj = w.IzbornikRadova
		fmt.Println(rad.Naziv)
		fmt.Println(w.IzbornikRadova.Naziv)
	}
	return
}

func (w *WingCal) Podvrsta(vrstarada *model.WingVrstaRadova) {
	//vrstarada.Roditelj = w.IzbornikRadova

	w.GenerisanjeLinkova(vrstarada.PodvrsteRadova)
	w.Roditelj()

	if vrstarada.Element {
		w.PrikazaniElement = vrstarada
	} else {
		w.IzbornikRadova = vrstarada
	}
	return
}

func (w *WingCal) NazivRoditelja() func() {
	return func() {
		if w.IzbornikRadova.Roditelj != nil {
			w.Tema.H4(w.IzbornikRadova.Roditelj.Naziv)
		}
	}
}

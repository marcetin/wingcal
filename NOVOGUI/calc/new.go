package calc

import (
	"gioui.org/app"
	"gioui.org/unit"
	"github.com/gioapp/gel"
	"github.com/marcetin/wingcal/db"
	"github.com/marcetin/wingcal/model"
	"github.com/marcetin/wingcal/pkg/gelook"
)

func NewWingCal() *WingCal {
	//gofont.Register()
	wing := &WingCal{
		Naziv: "W-ing Solutions - Kalkulator",
		Window: app.NewWindow(
			app.Size(unit.Dp(1280), unit.Dp(1024)),
			app.Title("W-ing Solutions - Kalkulator"),
		),
		Tema:             gelook.NewDuoUItheme(),
		Strana:           "kalkulator",
		Db:               db.DuoUIdbInit("./BAZA"),
		PrikazaniElement: &model.WingVrstaRadova{},
		Suma: &model.WingIzabraniElementi{
			Elementi:                 []*model.WingIzabraniElement{},
			UkupanNeophodanMaterijal: make(map[int]model.WingNeophodanMaterijal),
		},
	}
	wing.NewMaterijal()
	wing.Radovi = model.WingVrstaRadova{
		Id:       0,
		Naziv:    "Radovi",
		Slug:     "radovi",
		Omogucen: false,
		Baza:     false,
		Element:  false,
		//PodvrsteRadova: wing.Db.DbReadAll("radovi"),
	}

	wing.EditPolja = &model.EditabilnaPoljaVrsteRadova{
		Id:       new(gel.Editor),
		Naziv:    new(gel.Editor),
		Opis:     new(gel.Editor),
		Obracun:  new(gel.Editor),
		Jedinica: new(gel.Editor),
		Cena:     new(gel.Editor),
		Slug:     new(gel.Editor),
		Omogucen: new(gel.CheckBox),
	}
	wing.EditPolja.Materijal = make(map[int]*gel.Editor)
	for _, materijal := range wing.Materijal {
		wing.EditPolja.Materijal[materijal.Id] = new(gel.Editor)
	}
	wing.Putanja = append(wing.Putanja, wing.Radovi.Naziv)
	return wing
}

func (w *WingCal) GenerisanjeEdita() (edit *model.EditabilnaPoljaVrsteRadova) {
	//w.EditabilnaPoljaVrsteRadova = make(map[int]*model.EditabilnaPoljaVrsteRadova)
	//for rad, _ := range radovi {
	//	w.EditabilnaPoljaVrsteRadova[rad] =
	return &model.EditabilnaPoljaVrsteRadova{
		Id:    new(gel.Editor),
		Naziv: new(gel.Editor),
		Opis: &gel.Editor{
			SingleLine: false,
		},
		Obracun:  new(gel.Editor),
		Jedinica: new(gel.Editor),
		Cena:     new(gel.Editor),
		Slug:     new(gel.Editor),
		Omogucen: new(gel.CheckBox),
	}
}

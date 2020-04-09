package calc

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/text"
	"github.com/gioapp/gel"
	"github.com/gioapp/gelook"
	"github.com/marcetin/wingcal/db"
	"github.com/marcetin/wingcal/model"
)

var (
	izbornikRadova = &layout.List{
		Axis: layout.Vertical,
	}
	sumList = &layout.List{
		Axis: layout.Vertical,
	}
	neophodanMaterijalList = &layout.List{
		Axis: layout.Vertical,
	}
	ukupanNeophodanMaterijalList = &layout.List{
		Axis: layout.Vertical,
	}
	nazadDugme = new(gel.Button)
	dodajDugme = new(gel.Button)
	kolicina   = &gel.DuoUIcounter{
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

type WingCal struct {
	Naziv                    string
	Window                   *app.Window
	Context                  *layout.Context
	Tema                     *gelook.DuoUItheme
	Strana                   string
	Edit                     bool
	LinkoviIzboraVrsteRadova map[int]*gel.Button
	EditPolja                *model.EditabilnaPoljaVrsteRadova
	Materijal                map[int]*model.WingMaterijal
	Radovi                   model.WingVrstaRadova
	Putanja                  []*model.WingVrstaRadova
	IzbornikRadova           *model.WingVrstaRadova
	Transfered               model.WingCalGrupaRadova
	Db                       *db.DuoUIdb
	Client                   *model.Client
	PrikazaniElement         *model.WingVrstaRadova
	Suma                     *model.WingIzabraniElementi
}

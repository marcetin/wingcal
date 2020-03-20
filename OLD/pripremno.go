package OLD

import (
	"github.com/marcetin/wingcal/model"
)

func pripremnozavrsniradovi() map[int]*model.WingVrstaRadova {
	return map[int]*model.WingVrstaRadova{
		0: &model.WingVrstaRadova{
			Id:    1,
			Naziv: "Dozvole",
			Slug:  "dozvole",
			Baza:  true,
		},
		1: &model.WingVrstaRadova{
			Id:    2,
			Naziv: "Priključci",
			Slug:  "prikljucci",
			Baza:  true,
		},
		2: &model.WingVrstaRadova{
			Id:    3,
			Naziv: "Saglasnosti",
			Slug:  "saglasnosti",
			Baza:  true,
		},
		3: &model.WingVrstaRadova{
			Id:    4,
			Naziv: "Priprema gradilišta",
			Slug:  "priprema_gradilista",
			Baza:  true,
		},
		4: &model.WingVrstaRadova{
			Id:    5,
			Naziv: "Skele",
			Slug:  "skele",
			Baza:  true,
		},
		5: &model.WingVrstaRadova{
			Id:    6,
			Naziv: "Ograde",
			Slug:  "ograde",
			Baza:  true,
		},
		6: &model.WingVrstaRadova{
			Id:    7,
			Naziv: "Nadstrešnice",
			Slug:  "nadstresnice",
			Baza:  true,
		},
		7: &model.WingVrstaRadova{
			Id:    8,
			Naziv: "Signalizacija",
			Slug:  "signalizacija",
			Baza:  true,
		},
		8: &model.WingVrstaRadova{
			Id:    9,
			Naziv: "Zaštite",
			Slug:  "zastite",
			Baza:  true,
		},
		9: &model.WingVrstaRadova{
			Id:    10,
			Naziv: "Čišćenja",
			Slug:  "ciscenja",
			Baza:  true,
		},
		10: &model.WingVrstaRadova{
			Id:    11,
			Naziv: "Pranja",
			Slug:  "pranja",
			Baza:  true,
		},
		11: &model.WingVrstaRadova{
			Id:    12,
			Naziv: "Šut",
			Slug:  "sut",
			Baza:  true,
		},
		12: &model.WingVrstaRadova{
			Id:    13,
			Naziv: "Ostali pripremno završni radovi",
			Slug:  "ostali_pripremno_zavrsni_radovi",
			Baza:  true,
		},
	}
}

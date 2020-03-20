package OLD

import (
	"github.com/marcetin/wingcal/model"
)

func molerskofarbarskiradovi() model.WingVrstaRadova {
	molerskofarbarskiradovi := map[int]model.WingVrstaRadova{
		0: model.WingVrstaRadova{
			Id:    1,
			Naziv: "Pripremni",
			Slug:  "pripremni",
			Baza:  true,
		},
		1: model.WingVrstaRadova{
			Id:    2,
			Naziv: "Vrata i prozori",
			Slug:  "vrata_i_prozori",
			Baza:  true,
		},
		2: model.WingVrstaRadova{
			Id:    3,
			Naziv: "Zidovi i plafoni",
			Slug:  "zidovi_i_plafoni",
			Baza:  true,
		},
		3: model.WingVrstaRadova{
			Id:    4,
			Naziv: "Tapete",
			Slug:  "tapete",
			Baza:  true,
		},
		4: model.WingVrstaRadova{
			Id:    5,
			Naziv: "Lamperije",
			Slug:  "lamperije",
			Baza:  true,
		},
		5: model.WingVrstaRadova{
			Id:    6,
			Naziv: "Strehe",
			Slug:  "strehe",
			Baza:  true,
		},
		6: model.WingVrstaRadova{
			Id:    7,
			Naziv: "Stepeništa",
			Slug:  "stepenista",
			Baza:  true,
		},
		7: model.WingVrstaRadova{
			Id:    8,
			Naziv: "Podovi",
			Slug:  "podovi",
			Baza:  true,
		},
		8: model.WingVrstaRadova{
			Id:    9,
			Naziv: "Obrada fasada",
			Slug:  "obrada_fasada",
			Baza:  true,
		},
		9: model.WingVrstaRadova{
			Id:    10,
			Naziv: "Bojenja fasada",
			Slug:  "bojenja_fasada",
			Baza:  true,
		},
		10: model.WingVrstaRadova{
			Id:    11,
			Naziv: "Zaštita fasada",
			Slug:  "zastita_fasada",
			Baza:  true,
		},
		11: model.WingVrstaRadova{
			Id:    12,
			Naziv: "Limarija",
			Slug:  "limarija",
			Baza:  true,
		},
		12: model.WingVrstaRadova{
			Id:    13,
			Naziv: "Ograde",
			Slug:  "ograde",
			Baza:  true,
		},
		13: model.WingVrstaRadova{
			Id:    14,
			Naziv: "Instalacije",
			Slug:  "instalacije",
			Baza:  true,
		},
		14: model.WingVrstaRadova{
			Id:    15,
			Naziv: "Mobilijar",
			Slug:  "mobilijar",
			Baza:  true,
		},
		15: model.WingVrstaRadova{
			Id:    16,
			Naziv: "Ostali molersko - farbarski radovi",
			Slug:  "ostali_molersko_farbarski_radovi",
			Baza:  true,
		},
	}
	return model.WingVrstaRadova{
		Id:             26,
		Naziv:          "MOLERSKO FARBARSKI RADOVI",
		Slug:           "molersko_farbarski_radovi",
		Omogucen:       false,
		Baza:           false,
		Element:        false,
		Roditelj:       nil,
		PodvrsteRadova: molerskofarbarskiradovi,
	}
}

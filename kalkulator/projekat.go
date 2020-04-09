package kalkulator

import (
	"gioui.org/layout"
	"github.com/marcetin/wingcal/calc"
	"github.com/marcetin/wingcal/model"
)

func projekat(w *calc.WingCal) func() {
	return func() {
		layout.Flex{
			Axis: layout.Horizontal,
		}.Layout(w.Context,
			layout.Flexed(1, glavniDeoProjekat(w)),
		)
	}
}

func glavniDeoProjekat(w *calc.WingCal) func() {
	return func() {
		projekat := projekteTest()
		if w.Edit {
			layout.Flex{
				Axis: layout.Horizontal,
			}.Layout(w.Context,
				layout.Flexed(1, w.EditorElementaIzgled()),
			)
		} else {
			layout.Flex{
				Axis: layout.Horizontal,
			}.Layout(w.Context,
				layout.Rigid(func() {
					w.Tema.Body1(projekat.Naziv)
				}),
			)
		}
	}
}

func projekteTest() *model.WingProjekat {
	objekat := &model.WingObjekat{
		BrojObjekta:          "1",
		KategorijaObjekta:    "A",
		KlasifikacijaObjekta: "100101",
		Funkcija:             "Apartman",
		Gradjenje:            "Nova gradnja",
		Spratnost:            "Po2+Po1+P+2",
		Lokacija:             "Ćuštica",
		Ulica:                "Nova",
		Broj:                 "bb",
		KP:                   "KP 4091/6",
		KO:                   "K.O. Ćuštica",
	}

	investitor := &model.WingInvestitor{
		NazivInvestitora:         "Mijajlović Boban",
		LokacijaInvestitoraUlica: "Ristosija Ristića",
		LokacijaInvestitoraBroj:  "12",
		LokacijaInvestitoraGrad:  "Svrljig",
		PIB:                      "",
		MB:                       "",
		OdgovornoLiceIme:         "Boban",
		OdgovornoLicePrezime:     "Mijajlović ",
		FunkcijaOdgovornogLica:   "vlasnik",
	}

	projektant := &model.WingProjektant{
		NazivProjektanta:         "W-ING SOLUTIONS DOO",
		LokacijaProjektantaUlica: "Bulevar Oslobođenja",
		LokacijaProjektantaBroj:  "30a",
		LokacijaProjektantaGrad:  "Novi Sad",
		PIB:                      "106892584",
		MB:                       "20701005",
		OdgovornoLiceIme:         "Čedomir",
		OdgovornoLicePrezime:     "Vukobrat",
		FunkcijaOdgovornoglica:   "direktor",
	}

	odgovorniProjektant := &model.WingOdgovorniProjektant{
		OdgovornoLiceIime:      "Čedomir",
		OdgovornoLicePrezime:   "Vukobrat",
		JMBG:                   "2302978750021",
		BrojLicence:            "311 J420 10",
		FunkcijaOdgovornogLica: "dipl.inž.građ.",
		IDProjektant:           "",
	}

	projekat := &model.WingProjekat{
		Id:             0,
		Naziv:          "IDP Idejni projekat",
		Opis:           "IDP Idejni projekat IDP Idejni projekat IDP Idejni projekat",
		IdProjekta:     "E – 002/2020.IDP.",
		BrojDokumenta:  "005",
		Sveska:         "2",
		VrstaDokumenta: "projekat_konstrukcije",
		DatumDokumenta: "januar 2020. god",
		//Objekti              []WingObjekat
		//Investitori          []WingInvestitor
		//Projektant           []WingProjektant
		OdgovorniProjektant: odgovorniProjektant,
	}
	projekat.Objekti = append(projekat.Objekti, objekat)
	projekat.Investitori = append(projekat.Investitori, investitor)
	projekat.Projektant = append(projekat.Projektant, projektant)

	return projekat
}

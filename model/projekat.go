package model

type WingProjekat struct {
	Id                  int    `json:"id"`
	Naziv               string `json:"naziv"`
	Opis                string `json:"opis"`
	IdProjekta          string `json:"id_projekta"`
	BrojDokumenta       string `json:"broj_dokumenta"`
	Sveska              string `json:"sveska"`
	VrstaDokumenta      string `json:"vrsta_dokumenta"`
	DatumDokumenta      string `json:"datum"`
	Objekti             []*WingObjekat
	Investitori         []*WingInvestitor
	Projektant          []*WingProjektant
	OdgovorniProjektant *WingOdgovorniProjektant
	Dokumntacija        WingDokumentacija
}

type WingObjekat struct {
	BrojObjekta          string `json:"broj_objekta"`
	KategorijaObjekta    string `json:"kategorija_objekta"`
	KlasifikacijaObjekta string `json:"klasifikacija_objekta"`
	Funkcija             string `json:"funkcija"`
	Gradjenje            string `json:"gradjenje"`
	Spratnost            string `json:"spratnost"`
	Lokacija             string `json:"lokacija"`
	Ulica                string `json:"Ulica"`
	Broj                 string `json:"broj"`
	KP                   string `json:"KP"`
	KO                   string `json:"KO"`
}

type WingInvestitor struct {
	NazivInvestitora         string `json:"naziv_investitora"`
	LokacijaInvestitoraUlica string `json:"lokacija_investitora_ulica"`
	LokacijaInvestitoraBroj  string `json:"lokacija_investitora_broj"`
	LokacijaInvestitoraGrad  string `json:"lokacija_investitora_grad"`
	PIB                      string `json:"PIB"`
	MB                       string `json:"MB"`
	OdgovornoLiceIme         string `json:"odgovorno_lice_ime"`
	OdgovornoLicePrezime     string `json:"odgovorno_lice_prezime"`
	FunkcijaOdgovornogLica   string `json:"funkcija_odgovornog_lica"`
}

type WingProjektant struct {
	NazivProjektanta         string `json:"naziv_projektanta"`
	LokacijaProjektantaUlica string `json:"lokacija_projektanta_ulica"`
	LokacijaProjektantaBroj  string `json:"lokacija_projektanta_broj"`
	LokacijaProjektantaGrad  string `json:"lokacija_projektanta_grad"`
	PIB                      string `json:"PIB"`
	MB                       string `json:"MB"`
	OdgovornoLiceIme         string `json:"odgovorno_lice_ime"`
	OdgovornoLicePrezime     string `json:"odgovorno_lice_prezime"`
	FunkcijaOdgovornoglica   string `json:"funkcija_odgovornog_lica"`
}

type WingOdgovorniProjektant struct {
	OdgovornoLiceIime      string `json:"odgovorno_lice_ime"`
	OdgovornoLicePrezime   string `json:"odgovorno_lice_prezime"`
	JMBG                   string `json:"jmbg"`
	BrojLicence            string `json:"broj_licence"`
	FunkcijaOdgovornogLica string `json:"funkcija_odgovornog_lica"`
	IDProjektant           string `json:"id"`
}

type WingDokumentacija struct {
	Tekstualna string
	Numericka  string
	Graficka   string
}

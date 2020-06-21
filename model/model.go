package model

import (
	"github.com/gioapp/gel"
	"net"
)

type Client struct {
	Socket net.Conn
	data   chan []byte
}

type WingVrstaRadova struct {
	Id                 int                            `json:"id"`
	Naziv              string                         `json:"naziv"`
	Opis               string                         `json:"opis"`
	Obracun            string                         `json:"obracun"`
	Jedinica           string                         `json:"jedinica"`
	Cena               float64                        `json:"cena"`
	Slug               string                         `json:"slug"`
	Omogucen           bool                           `json:"omogucen"`
	Baza               bool                           `json:"baza"`
	Element            bool                           `json:"element"`
	PodvrsteRadova     map[int]WingVrstaRadova        `json:"podvrsteradova"`
	NeophodanMaterijal map[int]WingNeophodanMaterijal `json:"neophodanmaterijal"`
}

type WingIzabraniElementi struct {
	Id                             string
	SumaCena                       float64
	SumaCenaMaterijal              float64
	Elementi                       []*WingIzabraniElement
	UkupanNeophodanMaterijal       map[int]WingNeophodanMaterijal
	UkupanNeophodanMaterijalPrikaz map[int]WingNeophodanMaterijal
}

type WingMaterijal struct {
	Id      int    `json:"id"`
	Naziv   string `json:"naziv"`
	Opis    string `json:"opis"`
	Obracun string `json:"obracun"`

	Proizvodjac       string  `json:"proizvodjac"`
	OsobineNamena     string  `json:"osobinenamena"`
	NacinRada         string  `json:"nacinrada"`
	JedinicaPotrosnje string  `json:"jedinicapotrosnje"`
	Potrosnja         float64 `json:"potrosnja"`
	RokUpotrebe       string  `json:"rokupotreba"`

	Jedinica  string  `json:"jedinica"`
	Pakovanje int     `json:"pakovanje"`
	Cena      float64 `json:"cena"`
	Slug      string  `json:"slug"`
}

type WingNeophodanMaterijal struct {
	Id              int           `json:"id"`
	Kolicina        float64       `json:"kolicina"`
	Koeficijent     float64       `json:"koeficijent"`
	UkupnoPakovanja int           `json:"ukupnopakovanja"`
	UkupnaCena      float64       `json:"ukupnacena"`
	Materijal       WingMaterijal `json:"materijal"`
}

type WingIzabraniElement struct {
	Sifra         string
	Kolicina      int
	SumaCena      float64
	Element       WingVrstaRadova
	DugmeBrisanje *gel.Button
}

type WingCalGrupaRadova struct {
	Id       string                  `json:"id"`
	Slug     string                  `json:"slug"`
	Elementi map[int]WingVrstaRadova `json:"elementi"`
}

type WingCalEcommands map[int]WingCalEcommand

type WingCalEcommand struct {
	Id       string
	Type     string
	Name     string      `json:"name"`
	Enabled  bool        `json:"enabled"`
	CompType string      `json:"comptype"`
	SubType  string      `json:"subtype"`
	Command  interface{} `json:"command"`
	Data     interface{} `json:"data"`
}
type EditabilnaPoljaVrsteRadova struct {
	Id        *gel.Editor
	Naziv     *gel.Editor
	Opis      *gel.Editor
	Obracun   *gel.Editor
	Jedinica  *gel.Editor
	Cena      *gel.Editor
	Slug      *gel.Editor
	Omogucen  *gel.CheckBox
	Materijal map[int]*gel.Editor
}

type ElementMenu struct {
	Title     string
	Materijal bool
}

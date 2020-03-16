package main

import (
	"gioui.org/font/gofont"
	"gioui.org/layout"
	"github.com/marcetin/wingcal/pkg/gel"
	"github.com/marcetin/wingcal/pkg/gelook"
	"net"
)

type Client struct {
	Socket net.Conn
	data   chan []byte
}

type WingCal struct {
	Title           string
	Context         *layout.Context
	Theme           *gelook.DuoUItheme
	ElementsButtons map[int]*gel.Button
	Transfered      WingCalGrupaElemenata
	Db              *DuoUIdb
	Client          *Client

	PrikazaniElement *WingCalElement
	Suma             *WingIzabraniElementi
}

type WingIzabraniElementi struct {
	Id       string
	SumaCena float64
	Elementi []WingIzabraniElement
}

type WingIzabraniElement struct {
	Kolicina int
	SumaCena float64
	Element  WingCalElement
}

type WingCalElement struct {
	Id       uint8   `json:"id"`
	Naziv    string  `json:"naziv"`
	Opis     string  `json:"opis"`
	Obracun  string  `json:"obracun"`
	Jedinica string  `json:"jedinica"`
	Cena     float64 `json:"cena"`
	Slug     string  `json:"slug"`
}

type WingCalGrupaElemenata struct {
	Id       string                 `json:"id"`
	Slug     string                 `json:"slug"`
	Elementi map[int]WingCalElement `json:"elementi"`
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

func NewWingCal() *WingCal {
	gofont.Register()
	return &WingCal{
		Title:            "ParallelCoinGUI",
		Theme:            gelook.NewDuoUItheme(),
		Db:               DuoUIdbInit("db"),
		PrikazaniElement: &WingCalElement{},
		Suma:             &WingIzabraniElementi{},
	}
}

package main

import (
	"fmt"
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"github.com/gioapp/gel"
	"github.com/gioapp/gelook"
	"github.com/marcetin/wingcal/calc"
)

var (
	izbornikRadova = &layout.List{
		Axis: layout.Vertical,
	}
	editLista = &layout.List{
		Axis: layout.Vertical,
	}
	thingEditTitle   = new(gel.Editor)
	thingEditContent = new(gel.Editor)
	//post             = new(model.DuoCMSpost)
	dodajDugme   = new(gel.Button)
	stampajDugme = new(gel.Button)
	nazadDugme   = new(gel.Button)
	kolicina     = &gel.DuoUIcounter{
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

type DuoCMSadmin struct {
	Menu *DuoCMSmenu
}

type DuoCMSmenu struct {
	Title string
	Items map[string]DuoCMSmenuItem
}

type DuoCMSmenuItem struct {
	Title       string
	Description string
	Icon        string
	Link        func()
	subItems    map[string]DuoCMSmenuItem
}

func Editor(gtx *layout.Context, th *gelook.DuoUItheme, editorControler *gel.Editor, label string, handler func(gel.EditorEvent)) func() {
	return func() {
		layout.UniformInset(unit.Dp(4)).Layout(gtx, func() {
			//gelook.DuoUIdrawRectangle(gtx, cs.Width.Max, cs.Height.Max,
			//	th.Colors["Gray"],
			//	[4]float32{0, 0, 0, 0}, [4]float32{0, 0, 0, 0})
			layout.UniformInset(unit.Dp(0)).Layout(gtx, func() {
				//cs := gtx.Constraints
				gelook.DuoUIdrawRectangle(gtx, 30, 30,
					th.Colors["Light"], [4]float32{0, 0, 0, 0}, [4]float32{0, 0,
						0, 0})
				e := th.DuoUIeditor(label)
				e.Font.Typeface = th.Fonts["Mono"]
				// e.Font.Style = text.Italic
				layout.UniformInset(unit.Dp(4)).Layout(gtx, func() {
					e.Layout(gtx, editorControler)
				})
				for _, e := range editorControler.Events(gtx) {
					switch e.(type) {
					case gel.ChangeEvent:
						handler(e)
					}
				}
			})
		})
	}
}

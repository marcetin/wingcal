// SPDX-License-Identifier: Unlicense OR MIT

package gelook

import (
	"fmt"
	"gioui.org/font"
	"gioui.org/font/opentype"
	"gioui.org/text"
	"gioui.org/unit"
	"golang.org/x/image/font/gofont/gobold"
	"golang.org/x/image/font/gofont/gobolditalic"
	"golang.org/x/image/font/gofont/goitalic"
	"golang.org/x/image/font/gofont/gomedium"
	"golang.org/x/image/font/gofont/gomediumitalic"
	"golang.org/x/image/font/gofont/gomono"
	"golang.org/x/image/font/gofont/gomonobold"
	"golang.org/x/image/font/gofont/gomonobolditalic"
	"golang.org/x/image/font/gofont/gomonoitalic"
	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/image/font/gofont/gosmallcaps"
	"golang.org/x/image/font/gofont/gosmallcapsitalic"
)

type DuoUItheme struct {
	Shaper        text.Shaper
	TextSize      unit.Value
	Colors        map[string]string
	Fonts         map[string]text.Typeface
	Icons         map[string]*DuoUIicon
	scrollBarSize int
}

func init() {
	register(text.Font{}, goregular.TTF)
	register(text.Font{Style: text.Italic}, goitalic.TTF)
	register(text.Font{Weight: text.Bold}, gobold.TTF)
	register(text.Font{Style: text.Italic, Weight: text.Bold}, gobolditalic.TTF)
	register(text.Font{Weight: text.Medium}, gomedium.TTF)
	register(text.Font{Weight: text.Medium, Style: text.Italic}, gomediumitalic.TTF)
	register(text.Font{Variant: "Mono"}, gomono.TTF)
	register(text.Font{Variant: "Mono", Weight: text.Bold}, gomonobold.TTF)
	register(text.Font{Variant: "Mono", Weight: text.Bold, Style: text.Italic}, gomonobolditalic.TTF)
	register(text.Font{Variant: "Mono", Style: text.Italic}, gomonoitalic.TTF)
	register(text.Font{Variant: "Mono", Style: text.Italic}, gomonoitalic.TTF)
	register(text.Font{Variant: "Smallcaps"}, gosmallcaps.TTF)
	register(text.Font{Variant: "Smallcaps", Style: text.Italic}, gosmallcapsitalic.TTF)
}

func register(fnt text.Font, ttf []byte) {
	face, err := opentype.Parse(ttf)
	if err != nil {
		panic(fmt.Sprintf("failed to parse font: %v", err))
	}
	fnt.Typeface = "Go"
	font.Register(fnt, face)
}

func NewDuoUItheme() *DuoUItheme {
	t := &DuoUItheme{
		Shaper: font.Default(),
	}
	t.Colors = NewDuoUIcolors()
	t.Fonts = NewDuoUIfonts()
	t.TextSize = unit.Sp(16)
	t.Icons = NewDuoUIicons()
	return t
}

func NewDuoUIfonts() (f map[string]text.Typeface) {
	f = make(map[string]text.Typeface)
	f["Primary"] = "go"
	f["Secondary"] = "plan9"
	f["Mono"] = "go"
	return f
}

func (t *DuoUItheme) ChangeLightDark() {
	light := t.Colors["Light"]
	dark := t.Colors["Dark"]
	lightGray := t.Colors["LightGrayIII"]
	darkGray := t.Colors["DarkGrayII"]
	t.Colors["Light"] = dark
	t.Colors["Dark"] = light
	t.Colors["LightGrayIII"] = darkGray
	t.Colors["DarkGrayII"] = lightGray
}

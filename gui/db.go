package main

import (
	"encoding/json"
	"fmt"
	scribble "github.com/nanobox-io/golang-scribble"
	"golang.org/x/text/unicode/norm"
	"unicode"
)

type DuoUIdb struct {
	DB     *scribble.Driver
	Folder string      `json:"folder"`
	Name   string      `json:"name"`
	Data   interface{} `json:"data"`
}

type Ddb interface {
	DbReadAllTypes()
	DbRead(folder, name string)
	DbReadAll(folder string) WingCalGrupaElemenata
	DbWrite(folder, name string, data interface{})
}

func DuoUIdbInit(dataDir string) (d *DuoUIdb) {
	d = new(DuoUIdb)
	db, err := scribble.New(dataDir, nil)
	if err != nil {
		fmt.Println("Error", err)
	}
	d.DB = db
	return
}

var skip = []*unicode.RangeTable{
	unicode.Mark,
	unicode.Sk,
	unicode.Lm,
}

var safe = []*unicode.RangeTable{
	unicode.Letter,
	unicode.Number,
}

var _ Ddb = &DuoUIdb{}

func (d *DuoUIdb) DbReadAllTypes() {
	items := make(map[string]WingCalGrupaElemenata)
	types := []string{"assets", "config", "apps"}
	for _, t := range types {
		items[t] = d.DbReadAll(t)
	}
	d.Data = items
	fmt.Println("ooooooooooooooooooooooooooooodaaa", d.Data)

}
func (d *DuoUIdb) DbReadTypeAll(f string) {
	d.Data = d.DbReadAll(f)
}

func (d *DuoUIdb) DbReadAll(folder string) WingCalGrupaElemenata {
	itemsRaw, err := d.DB.ReadAll(folder)
	if err != nil {
		fmt.Println("Error", err)
	}
	items := make(map[int]WingCalElement)
	for _, bt := range itemsRaw {
		item := WingCalElement{}
		if err := json.Unmarshal([]byte(bt), &item); err != nil {
			fmt.Println("Error", err)
		}
		//items[item.Slug] = item
	}
	return WingCalGrupaElemenata{
		Slug:     folder,
		Elementi: items,
	}
}

func (d *DuoUIdb) DbRead(folder, name string) {
	item := WingCalGrupaElemenata{}
	if err := d.DB.Read(folder, name, &item); err != nil {
		fmt.Println("Error", err)
	}
	d.Data = item
}
func (d *DuoUIdb) DbWrite(folder, name string, data interface{}) {
	d.DB.Write(folder, name, data)
}

func slug(text string) string {
	buf := make([]rune, 0, len(text))
	dash := false
	for _, r := range norm.NFKD.String(text) {
		switch {
		case unicode.IsOneOf(safe, r):
			buf = append(buf, unicode.ToLower(r))
			dash = true
		case unicode.IsOneOf(skip, r):
		case dash:
			buf = append(buf, '-')
			dash = false
		}
	}
	if i := len(buf) - 1; i >= 0 && buf[i] == '-' {
		buf = buf[:i]
	}
	return string(buf)
}

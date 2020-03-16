package main

import "C"
import (
	"bufio"
	"encoding/json"
	"fmt"
	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/unit"
	"github.com/marcetin/wingcal/gui/js"
	"github.com/marcetin/wingcal/pkg/gel"
	"github.com/marcetin/wingcal/pkg/gelook"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

const (
	Address = ":9999"

	addr = "localhost:4242"
)

var (
	btn = new(gel.Button)
)

func main() {
	wing := NewWingCal()
	js.BuildJS()
	//c, err := net.Dial("tcp", "localhost:19999")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//go func() {
	//	wing.StartClientModule(c, wing.Transfered)
	//}()
	//go connection.StartClientModule()

	wing.Db.DbRead("zidarski_radovi", "demontaze")
	//api.Db.Data.(model.WingCalGrupaElemenata)
	//api.Db.DbWrite("zidarski_radovi", "demontaze", demontaze)
	//fmt.Println("RECEIVEDzestr: ", thing.Data.(model.DuoCMSpost))
	//fmt.Println("RECEIVEDzestr: ", wing.Db.Data.(WingCalGrupaElemenata))
	wing.Transfered = wing.Db.Data.(WingCalGrupaElemenata)

	wing.ElementsButtons = map[int]*gel.Button{}

	for id, _ := range wing.Transfered.Elementi {
		wing.ElementsButtons[id] = new(gel.Button)
	}
	http.Handle("/", http.FileServer(http.Dir("./html")))
	//go panic(http.ListenAndServe(":8080", http.FileServer(http.Dir("/usr/share/doc"))))

	log.Println("Listening on :3333...")
	go http.ListenAndServe(":3333", nil)
	//err := http.ListenAndServe(":3333", nil)
	//if err != nil {
	//	log.Fatal(err)
	//}

	go func() {

		w := app.NewWindow(
			app.Size(unit.Dp(999), unit.Dp(999)),
			app.Title("ParallelCoin"),
		)
		wing.Context = layout.NewContext(w.Queue())
		for e := range w.Events() {
			if e, ok := e.(system.FrameEvent); ok {
				wing.Context.Reset(e.Config, e.Size)
				gelook.DuoUIfill(wing.Context, wing.Theme.Colors["DarkGrayII"])
				layout.Flex{
					Axis: layout.Vertical,
				}.Layout(wing.Context,
					layout.Rigid(func() {
						wing.Theme.DuoUIitem(0, wing.Theme.Colors["Primary"]).Layout(wing.Context, layout.Center, func() {
							wing.Theme.H6("Title").Layout(wing.Context)
						})
					}),
					layout.Flexed(1, wing.admin()),
					//func() {
					//	//cs := wing.Context.Constraints
					//	//helpers.DrawRectangle(wing.Context, cs.Width.Max, cs.Height.Max, helpers.HexARGB("ff3030cf"), [4]float32{0, 0, 0, 0}, unit.Dp(0))
					//	for btn.Clicked(wing.Context) {
					//		c.Write([]byte(strings.TrimRight("testa", "\n")))
					//	}
					//	wing.Theme.Button("Click me!").Layout(wing.Context, btn)
					//}),
					layout.Rigid(func() {
						gelook.DuoUIfill(wing.Context, wing.Theme.Colors["Hint"])
						layout.Center.Layout(wing.Context, func() {
							if wing.Transfered.Elementi != nil {
								//wing.Transfered.(func())()
								//wing.Theme.H6(wing.Transfered.Elementi).Layout(wing.Context)
								//for _, element := range wing.Transfered.Elementi {
								//	fmt.Println(element.Naziv)
								//}
							} else {
								wing.Theme.H6("Nonohhhhhnono").Layout(wing.Context)
							}
						})
					}),
				)
				e.Frame(wing.Context.Ops)
			}
		}
	}()
	app.Main()
	log.Print("Starting server...")
}

func (w *WingCal) Receive(t WingCalGrupaElemenata) {
	for {
		message := make([]byte, 4096)
		length, err := w.Client.Socket.Read(message)
		if err != nil {
			w.Client.Socket.Close()
			break
		}
		//var network bytes.Buffer
		//dec := gob.NewDecoder(&network)
		//err = dec.Decode(&t)
		//if err != nil {
		//	log.Fatal("decode error:", err)
		//}
		t = WingCalGrupaElemenata{}
		if err := json.Unmarshal([]byte(message), &t); err != nil {
			fmt.Println("Error", err)
		}
		//t = interface{}(message).(*model.DuoCMSthing)
		if length > 0 {
			fmt.Println("RECEIVED111: " + string(message))
			fmt.Println("RECEIVED444", t)
		}
	}
}

func (w *WingCal) StartClientModule(c net.Conn, t WingCalGrupaElemenata) {
	fmt.Println("Starting client...")
	w.Client = &Client{Socket: c}
	go w.Receive(t)
	for {
		reader := bufio.NewReader(os.Stdin)
		message, _ := reader.ReadString('\n')
		c.Write([]byte(strings.TrimRight(message, "\n")))
	}
}

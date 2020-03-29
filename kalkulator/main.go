package kalkulator

import "C"
import (
	"fmt"
	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"github.com/marcetin/wingcal/calc"
	"github.com/marcetin/wingcal/pkg/gel"
	"github.com/marcetin/wingcal/pkg/gelook"
	"log"
)

const (
	Address = ":9999"

	addr = "localhost:4242"
)

var (
	btn = new(gel.Button)
)

func C() {
	wing := calc.NewWingCal()
	//js.BuildJS()
	//c, err := net.Dial("tcp", "localhost:19999")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//go func() {
	//	wing.StartClientModule(c, wing.Transfered)
	//}()
	//go connection.StartClientModule()

	//wing.Db.DbRead("molerski_radovi", "pripremni")
	//api.Db.Data.(model.WingCalGrupaElemenata)
	//fmt.Println("RECEIVEDzestr: ", thing.Data.(model.DuoCMSpost))
	//fmt.Println("RECEIVEDzestr: ", wing.Db.Data.(WingCalGrupaElemenata))
	//wing.Transfered = wing.Db.Data.(WingCalGrupaElemenata)

	//wing.Radovi = NewRadovi()
	//wing.Db.DbWrite("radovi", "26", molerskofarbarskiradovi())

	//wing.Radovi = NewRadovi()

	//for _, rr := range wing.Radovi.PodvrsteRadova{
	//	fmt.Println("Slug")
	//	fmt.Println(rr.Slug)
	//	fmt.Println("id")
	//	fmt.Println(rr.Id)
	//}

	//wing.Transfered = sendvic()
	//fmt.Println(wing.Transfered)
	wing.LinkoviIzboraVrsteRadova = map[int]*gel.Button{}

	wing.GenerisanjeLinkova(wing.Radovi.PodvrsteRadova)
	wing.IzbornikRadova = &wing.Radovi
	wing.Roditelj()

	for _, p := range wing.IzbornikRadova.PodvrsteRadova {
		fmt.Println("------")
		fmt.Println(p.Naziv)
		//fmt.Println(p.Roditelj.Naziv)
	}
	//http.Handle("/", http.FileServer(http.Dir("./html")))
	//go panic(http.ListenAndServe(":8080", http.FileServer(http.Dir("/usr/share/doc"))))

	//log.Println("Listening on :3333...")
	//go http.ListenAndServe(":3333", nil)
	//err := http.ListenAndServe(":3333", nil)`
	//if err != nil {
	//	log.Fatal(err)
	//}

	go func() {

		wing.Context = layout.NewContext(wing.Window.Queue())
		for e := range wing.Window.Events() {
			if e, ok := e.(system.FrameEvent); ok {
				wing.Context.Reset(e.Config, e.Size)
				gelook.DuoUIfill(wing.Context, wing.Tema.Colors["Light"])
				layout.Flex{
					Axis: layout.Vertical,
				}.Layout(wing.Context,
					layout.Rigid(func() {
						wing.Tema.DuoUIcontainer(0, wing.Tema.Colors["Primary"]).Layout(wing.Context, layout.Center, func() {
							wing.Tema.H6("W-ing Solutions ").Layout(wing.Context)
						})
					}),
					layout.Flexed(1, glavniEkran(wing)),
					//func() {
					//	//cs := wing.Context.Constraints
					//	//helpers.DrawRectangle(wing.Context, cs.Width.Max, cs.Height.Max, helpers.HexARGB("ff3030cf"), [4]float32{0, 0, 0, 0}, unit.Dp(0))
					//	for btn.Clicked(wing.Context) {
					//		c.Write([]byte(strings.TrimRight("testa", "\n")))
					//	}
					//	wing.Tema.Button("Click me!").Layout(wing.Context, btn)
					//}),
					layout.Rigid(func() {
						gelook.DuoUIfill(wing.Context, wing.Tema.Colors["Hint"])
						layout.Center.Layout(wing.Context, func() {
							if wing.Transfered.Elementi != nil {
								//wing.Transfered.(func())()
								//wing.Tema.H6(wing.Transfered.Elementi).Layout(wing.Context)
								//for _, element := range wing.Transfered.Elementi {
								//	fmt.Println(element.Naziv)
								//}
							} else {
								wing.Tema.H6("Nonohhhhhnono").Layout(wing.Context)
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

//
//func (w *WingCal) Receive(t WingCalGrupaRadova) {
//	for {
//		message := make([]byte, 4096)
//		length, err := w.Client.Socket.Read(message)
//		if err != nil {
//			w.Client.Socket.Close()
//			break
//		}
//		//var network bytes.Buffer
//		//dec := gob.NewDecoder(&network)
//		//err = dec.Decode(&t)
//		//if err != nil {
//		//	log.Fatal("decode error:", err)
//		//}
//		t = WingCalGrupaRadova{}
//		if err := json.Unmarshal([]byte(message), &t); err != nil {
//			fmt.Println("Error", err)
//		}
//		//t = interface{}(message).(*model.DuoCMSthing)
//		if length > 0 {
//			fmt.Println("RECEIVED111: " + string(message))
//			fmt.Println("RECEIVED444", t)
//		}
//	}
//}
//
//func (w *WingCal) StartClientModule(c net.Conn, t WingCalGrupaRadova) {
//	fmt.Println("Starting client...")
//	w.Client = &Client{Socket: c}
//	go w.Receive(t)
//	for {
//		reader := bufio.NewReader(os.Stdin)
//		message, _ := reader.ReadString('\n')
//		c.Write([]byte(strings.TrimRight(message, "\n")))
//	}
//}

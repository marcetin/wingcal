package main

import (
	"encoding/json"
	"fmt"
	"github.com/marcetin/wingcal/model"
	"net"
)

const (
	Address = ":9999"

	addr = "localhost:4242"

	message = "foobar"
)

type DuoCMSapi struct {
	Title string
	//Theme      *gelook.DuoUItheme
	//Transfered *DuoCMSthing
	ClientManager *ClientManager
	Db            *model.DuoUIdb
}

func NewDuoCMSapi() *DuoCMSapi {
	return &DuoCMSapi{
		Title: "ParallelCoinApi",
		Db:    model.DuoUIdbInit("db"),
	}
}

type ClientManager struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}

type Client struct {
	Socket net.Conn
	data   chan []byte
}

func (manager *ClientManager) start() {
	for {
		select {
		case connection := <-manager.register:
			manager.clients[connection] = true
			fmt.Println("Added new connection!")
		case connection := <-manager.unregister:
			if _, ok := manager.clients[connection]; ok {
				close(connection.data)
				delete(manager.clients, connection)
				fmt.Println("A connection has terminated!")
			}
		case message := <-manager.broadcast:
			for connection := range manager.clients {
				select {
				case connection.data <- message:
				default:
					close(connection.data)
					delete(manager.clients, connection)
				}
			}
		}
	}
}

func (api *DuoCMSapi) Receive(client *Client) {
	for {
		message := make([]byte, 4096)
		length, err := client.Socket.Read(message)
		if err != nil {
			api.ClientManager.unregister <- client
			client.Socket.Close()
			break
		}
		if length > 0 {
			fmt.Println("RECEIVEDserver: " + string(message))

			//thing := model.DuoCMSthing{}
			//if err := json.Unmarshal(message, &thing); err != nil {
			//}
			//demontaze := &model.WingCalGrupaElemenata{
			//	Id:   "1",
			//	Slug: "demontaze",
			//	Elementi: map[string]model.WingCalElement{
			//		"0": {
			//			Naziv:    "Demontaža prozorskih rešetki",
			//			Opis:     "Demontaža prozorskih rešetki sa čišćenjem i odvozom na deponiju koju odredi nadzorni organ, na udaljenost do 15 km.",
			//			Obracun:  "Obračun po komadu rešetke.",
			//			Jedinica: "Kom",
			//			Cena:     5.00,
			//		},
			//		"1": {
			//			Naziv:    "Demontaža metalne zastakljene nadstrešnice.",
			//			Opis:     "Demontaža metalne zastakljene nadstrešnice. Prvo demontirati staklo a zatim metalne delove nadstrešnice, sve očistiti i složiti za ponovnu ugradnju ili utovariti u kamion i odvesti na deponiju koju odredi investitor udaljenu do 15 km. Šut prikupiti, izneti, utovariti na kamion i odvesti na gradsku deponiju",
			//			Obracun:  "Obračun po m2 nadstrešnice.",
			//			Jedinica: "m2",
			//			Cena:     6.00,
			//		},
			//		"2": {
			//			Naziv:    "",
			//			Opis:     "Pažljiva demontaža parketa zajedno sa lajsnama. Parket i lajsne pažljivo demontirati, očistiti, složiti po vrsti, upakovati, utovariti u kamion i odvesti na deponiju koju odredi investitor udaljenu do 15 km.Šut prikupiti, izneti, utovariti na kamion i odvesti na gradsku deponiju.",
			//			Obracun:  "Obračun po m2 poda.",
			//			Jedinica: "m2",
			//			Cena:     4.00,
			//		},
			//	},
			//}
			api.Db.DbRead("zidarski_radovi", "demontaze")
			//api.Db.Data.(model.WingCalGrupaElemenata)
			//api.Db.DbWrite("zidarski_radovi", "demontaze", demontaze)
			//fmt.Println("RECEIVEDzestr: ", thing.Data.(model.DuoCMSpost))
			fmt.Println("RECEIVEDzestr: ", api.Db.Data.(model.WingCalGrupaElemenata))
			//spew.Dump(thing)
			//func(){cms.Theme.H6("Nononono").Layout(gtx)}
			//send := &model.DuoCMSthing{
			//	Id:       "999",
			//	Type:     "sysTest",
			//	Name:     "ParallelCoin",
			//	Enabled:  true,
			//	Slug:     "duo",
			//	Version:  "0.0",
			//	CompType: "system",
			//	SubType:  "",
			//}
			//var network bytes.Buffer        // Stand-in for a network connection
			//enc := gob.NewEncoder(&network) // Will write to network.
			//
			//err := enc.Encode(send)
			//if err != nil {
			//	log.Fatal("encode error:", err)
			//}
			b, err := json.MarshalIndent(api.Db.Data.(model.WingCalGrupaElemenata), "", "\t")
			if err != nil {
			}

			api.ClientManager.broadcast <- b
			//api.ClientManager.broadcast <- []byte(strings.TrimRight(string(message)+"KaoCadatesta", "\n"))
		}
	}
}

func main() {
	api := NewDuoCMSapi()

	fmt.Println("Starting server...")
	listener, error := net.Listen("tcp", ":19999")
	if error != nil {
		fmt.Println(error)
	}
	api.ClientManager = &ClientManager{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
	go api.ClientManager.start()
	for {
		conn, _ := listener.Accept()
		if error != nil {
			fmt.Println(error)
		}
		client := &Client{Socket: conn, data: make(chan []byte)}
		api.ClientManager.register <- client
		go api.Receive(client)
		go api.ClientManager.Send(client)
	}
}

func (manager *ClientManager) Send(client *Client) {
	defer client.Socket.Close()
	for {
		select {
		case message, ok := <-client.data:
			if !ok {
				return
			}
			client.Socket.Write(message)
		}
	}
}

package connection

import (
	"bufio"
	"fmt"
	"github.com/marcetin/wingcal/model"
	"net"
	"os"
	"strings"
)

//type ClientManager struct {
//	clients    map[*Client]bool
//	broadcast  chan []byte
//	register   chan *Client
//	unregister chan *Client
//}
//
//type Client struct {
//	Socket net.Conn
//	data   chan []byte
//}
//
//func (manager *ClientManager) start() {
//	for {
//		select {
//		case connection := <-manager.register:
//			manager.clients[connection] = true
//			fmt.Println("Added new connection!")
//		case connection := <-manager.unregister:
//			if _, ok := manager.clients[connection]; ok {
//				close(connection.data)
//				delete(manager.clients, connection)
//				fmt.Println("A connection has terminated!")
//			}
//		case message := <-manager.broadcast:
//			for connection := range manager.clients {
//				select {
//				case connection.data <- message:
//				default:
//					close(connection.data)
//					delete(manager.clients, connection)
//				}
//			}
//		}
//	}
//}
//
//func (manager *ClientManager) Receive(client *Client) {
//	for {
//		message := make([]byte, 4096)
//		length, err := client.Socket.Read(message)
//		if err != nil {
//			manager.unregister <- client
//			client.Socket.Close()
//			break
//		}
//		if length > 0 {
//			fmt.Println("RECEIVEDserver: " + string(message))
//
//
//			thing := model.DuoCMSthing{}
//			if err := json.Unmarshal(message, &thing); err != nil {
//			}
//
//
//			//func(){cms.Theme.H6("Nononono").Layout(gtx)}
//			send := &model.DuoCMSthing{
//				Id:       "999",
//				Type:     "sysTest",
//				Name:     "ParallelCoin",
//				Enabled:  true,
//				Slug:     "duo",
//				Version:  "0.0",
//				CompType: "system",
//				SubType:  "",
//			}
//			//var network bytes.Buffer        // Stand-in for a network connection
//			//enc := gob.NewEncoder(&network) // Will write to network.
//			//
//			//err := enc.Encode(send)
//			//if err != nil {
//			//	log.Fatal("encode error:", err)
//			//}
//			b, err := json.MarshalIndent(send, "", "\t")
//			if err != nil {
//			}
//
//			manager.broadcast <- b
//			//manager.broadcast <- []byte(strings.TrimRight(string(message)+"KaoCadatesta", "\n"))
//		}
//	}
//}
func thingInterface(t *model.WingCalElement) interface{} {
	return t
}

func passInterface(v interface{}) {
	b, ok := v.(*[]byte)
	fmt.Println(ok)
	fmt.Println(b)
}

//func (client *Client) Receive(t *model.WingCalGrupaElemenata) {
//	for {
//		message := make([]byte, 4096)
//		length, err := client.Socket.Read(message)
//		if err != nil {
//			client.Socket.Close()
//			break
//		}
//		//var network bytes.Buffer
//		//dec := gob.NewDecoder(&network)
//		//err = dec.Decode(&t)
//		//if err != nil {
//		//	log.Fatal("decode error:", err)
//		//}
//		t = &model.WingCalGrupaElemenata{}
//		if err := json.Unmarshal([]byte(message), &t); err != nil {
//			fmt.Println("Error", err)
//		}
//		//t = interface{}(message).(*model.DuoCMSthing)
//		if length > 0 {
//			fmt.Println("RECEIVED111: " + string(message))
//		}
//	}
//}

//func (manager *ClientManager) Send(client *Client) {
//	defer client.Socket.Close()
//	for {
//		select {
//		case message, ok := <-client.data:
//			if !ok {
//				return
//			}
//			client.Socket.Write(message)
//		}
//	}
//}

//func StartServerModule(d *model.DuoCMSapi) {
//	fmt.Println("Starting server...")
//	listener, error := net.Listen("tcp", ":19999")
//	if error != nil {
//		fmt.Println(error)
//	}
//	manager := ClientManager{
//		clients:    make(map[*Client]bool),
//		broadcast:  make(chan []byte),
//		register:   make(chan *Client),
//		unregister: make(chan *Client),
//	}
//	go manager.start()
//	for {
//		connection, _ := listener.Accept()
//		if error != nil {
//			fmt.Println(error)
//		}
//		client := &Client{Socket: connection, data: make(chan []byte)}
//		manager.register <- client
//		go manager.Receive(client)
//		go manager.Send(client)
//	}
//}

//func StartClientModule(c net.Conn, t *model.WingCalGrupaElemenata) {
//	fmt.Println("Starting client...")
//	client := &Client{Socket: c}
//	go client.Receive(t)
//	for {
//		reader := bufio.NewReader(os.Stdin)
//		message, _ := reader.ReadString('\n')
//		c.Write([]byte(strings.TrimRight(message, "\n")))
//	}
//}

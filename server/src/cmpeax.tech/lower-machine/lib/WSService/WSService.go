package WSService

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"cmpeax.tech/lower-machine/lib/SService"

	"cmpeax.tech/lower-machine/lib/routerDI"

	"github.com/gorilla/websocket"
)

type WSService struct {
	clients   map[*websocket.Conn]bool
	broadcast chan routerDI.Message
	upgrader  *websocket.Upgrader
	addr      string

	ssConn *SService.SService //socket设备指针
	db     *sql.DB
}

func NewWSService(addr string, dbobj *sql.DB, ssConn *SService.SService) *WSService {
	return &WSService{
		addr:      addr,
		db:        dbobj,
		ssConn:    ssConn,
		clients:   make(map[*websocket.Conn]bool), // connected clients
		broadcast: make(chan routerDI.Message),    // broadcast channel
		// Configure the upgrader
		upgrader: &websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}
}

func (wss *WSService) StartWService() {

	http.HandleFunc("/", wss.handleConnections)
	fmt.Println("websocket Server Start!")
	go wss.handleMessages()

	err := http.ListenAndServe(wss.addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

func (wss *WSService) handleConnections(w http.ResponseWriter, r *http.Request) {
	fmt.Println("one Connected!")
	//get Connection Point
	ws, err := wss.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	defer ws.Close()

	// Register our new client
	wss.clients[ws] = true

	for {
		var msg routerDI.Message
		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(wss.clients, ws)
			break
		}

		// Send the newly received message to the broadcast channel
		wss.broadcast <- msg
	}
}

func (wss *WSService) handleMessages() {

	for {
		// Grab the next message from the broadcast channel

		msg := <-wss.broadcast
		fmt.Println("outputstring")
		fmt.Println(msg.Code)
		wss.ssConn.RequestOfWS <- msg

		// 转发
		// for client := range wss.clients {
		// 	err := client.WriteJSON(msg)
		// 	if err != nil {
		// 		client.Close()
		// 		delete(wss.clients, client)
		// 	}
		// }
	}
}

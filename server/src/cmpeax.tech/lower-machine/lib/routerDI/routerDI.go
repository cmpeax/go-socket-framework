package routerDI

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net"
)

type CallbackJSONFunc func(jsonData map[string]interface{}, conn net.Conn, db *sql.DB)
type WSCallbackJSONFunc func(jsonData Message)

type MapOfCallbackJSONFunc map[string]CallbackJSONFunc
type MapOfWSCallbackJSONFunc map[string]WSCallbackJSONFunc

type Message struct {
	Code     string `json:"code"`
	DeviceID string `json:"deviceID"`
	Ip       string `json:"ip"`
	Data     string `json:"data"`
}

// Define our message object
type Router struct {
	conn      *net.Conn
	db        *sql.DB
	matchList MapOfCallbackJSONFunc
}

type WSRouter struct {
	db        *sql.DB
	matchList MapOfWSCallbackJSONFunc
}

//初始化 socket 路由
func InitRouter(dconn *net.Conn, matchList MapOfCallbackJSONFunc, dbobj *sql.DB) *Router {
	return &Router{
		conn:      dconn,
		matchList: matchList,
		db:        dbobj,
	}
}

//初始化 websocket 路由
func InitWSRouter(matchList MapOfWSCallbackJSONFunc, dbobj *sql.DB) *WSRouter {
	return &WSRouter{
		matchList: matchList,
		db:        dbobj,
	}
}

// startMatchJson
func (r *Router) StartMatchJson(matchCodeStr string) {
	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(matchCodeStr), &dat); err == nil {

		for key, callback := range r.matchList {
			if dat["code"] == key {
				callback(dat, *r.conn, r.db)
			}
		}
	} else {
		fmt.Println(err)
	}
}

// startWSMatchJson'

func (r *WSRouter) StartWSMatchJson(msg Message) {

	for key, callback := range r.matchList {
		if msg.Code == key {
			fmt.Printf("match: %s\n", key)
			callback(msg)
		}

	}
}

package routerDI

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net"
)

type CallbackJSONFunc func(jsonData map[string]interface{}, conn net.Conn, db *sql.DB)

type Router struct {
	conn      net.Conn
	db        *sql.DB
	matchList map[string]CallbackJSONFunc
}

func InitRouter(dconn net.Conn, matchList map[string]CallbackJSONFunc, dbobj *sql.DB) *Router {
	return &Router{
		conn:      dconn,
		matchList: matchList,
		db:        dbobj,
	}
}

// func (r *Router) StartMatch(matchCodeStr string) {
// 	for key, callback := range r.matchList {
// 		if matchCodeStr == key {
// 			fmt.Printf("match: %s", key)
// 			callback(r.conn, r.db)
// 		}
// 	}
// }

// startMatchJson
func (r *Router) StartMatchJson(matchCodeStr string) {
	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(matchCodeStr), &dat); err == nil {

		for key, callback := range r.matchList {
			fmt.Println(dat["code"])
			if dat["code"] == key {
				fmt.Printf("match: %s", key)
				callback(dat, r.conn, r.db)
			}
		}
	} else {
		fmt.Println(err)
	}

}
